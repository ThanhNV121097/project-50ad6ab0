package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/ThanhNV121097/project-50ad6ab0/backend/migrations"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const canonicalMessageID = "00000000-0000-0000-0000-000000000001"

func main() {
	ctx := context.Background()
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is required")
	}

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("connect database: %v", err)
	}
	defer pool.Close()

	if err := runMigrations(ctx, pool); err != nil {
		log.Fatalf("run migrations: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		requestID := requestIDFrom(r)
		if r.Method != http.MethodGet {
			writeError(w, http.StatusMethodNotAllowed, "INTERNAL", "method not allowed", requestID)
			return
		}
		healthCtx, cancel := context.WithTimeout(r.Context(), time.Second)
		defer cancel()
		if err := pool.Ping(healthCtx); err != nil {
			writeError(w, http.StatusServiceUnavailable, "UNAVAILABLE", "database unavailable", requestID)
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"}, requestID)
	})
	mux.HandleFunc("/v1/message", func(w http.ResponseWriter, r *http.Request) {
		requestID := requestIDFrom(r)
		if r.Method != http.MethodGet {
			writeError(w, http.StatusMethodNotAllowed, "INTERNAL", "method not allowed", requestID)
			return
		}
		if err := rejectBody(r); err != nil {
			writeError(w, http.StatusBadRequest, "VALIDATION_FAILED", "request body not allowed", requestID)
			return
		}
		messageCtx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		message, err := loadMessage(messageCtx, pool)
		if err != nil {
			writeAppError(w, err, requestID)
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"message": message}, requestID)
	})

	addr := ":" + port()
	server := &http.Server{Addr: addr, Handler: mux, ReadHeaderTimeout: 5 * time.Second}
	log.Printf("listening on %s", addr)
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

func requestIDFrom(r *http.Request) string {
	if value := strings.TrimSpace(r.Header.Get("X-Request-Id")); value != "" {
		return value
	}
	var b [16]byte
	if _, err := rand.Read(b[:]); err == nil {
		return hex.EncodeToString(b[:])
	}
	return fmt.Sprintf("req-%d", time.Now().UnixNano())
}

func rejectBody(r *http.Request) error {
	if r.Body == nil {
		return nil
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if len(body) == 0 {
		return nil
	}
	return fmt.Errorf("request body not allowed")
}

func loadMessage(ctx context.Context, pool *pgxpool.Pool) (string, error) {
	var content string
	err := pool.QueryRow(ctx, `SELECT content FROM messages WHERE id = $1`, canonicalMessageID).Scan(&content)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", appError{status: http.StatusNotFound, code: "NOT_FOUND", message: "message missing"}
		}
		if isUnavailable(err) {
			return "", appError{status: http.StatusServiceUnavailable, code: "UNAVAILABLE", message: "database unavailable"}
		}
		return "", appError{status: http.StatusInternalServerError, code: "INTERNAL", message: "internal server error"}
	}
	if strings.TrimSpace(content) == "" || strings.ContainsAny(content, "\r\n") {
		return "", appError{status: http.StatusUnprocessableEntity, code: "VALIDATION_FAILED", message: "stored message invalid"}
	}
	return content, nil
}

type appError struct {
	status  int
	code    string
	message string
}

func (e appError) Error() string { return e.message }

func writeAppError(w http.ResponseWriter, err error, requestID string) {
	var appErr appError
	if errors.As(err, &appErr) {
		writeError(w, appErr.status, appErr.code, appErr.message, requestID)
		return
	}
	writeError(w, http.StatusInternalServerError, "INTERNAL", "internal server error", requestID)
}

func writeError(w http.ResponseWriter, status int, code, message, requestID string) {
	writeJSON(w, status, map[string]any{"error": map[string]any{"code": code, "message": message, "details": []any{}, "request_id": requestID}}, requestID)
}

func writeJSON(w http.ResponseWriter, status int, payload any, requestID string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Request-Id", requestID)
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func isUnavailable(err error) bool { return strings.Contains(strings.ToLower(err.Error()), "timeout") || strings.Contains(strings.ToLower(err.Error()), "closed") || strings.Contains(strings.ToLower(err.Error()), "unavailable") }

func port() string {
	if value := os.Getenv("PORT"); value != "" {
		return value
	}
	if value := os.Getenv("APP_PORT"); value != "" {
		return value
	}
	return "8080"
}

func runMigrations(ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, `CREATE TABLE IF NOT EXISTS schema_migrations (version text PRIMARY KEY, applied_at timestamptz NOT NULL DEFAULT now())`)
	if err != nil {
		return fmt.Errorf("ensure schema_migrations: %w", err)
	}

	entries, err := fs.ReadDir(migrations.Files, ".")
	if err != nil {
		return fmt.Errorf("read embedded migrations: %w", err)
	}
	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		name := entry.Name()
		if !entry.IsDir() && strings.HasSuffix(name, ".up.sql") {
			names = append(names, name)
		}
	}
	sort.Strings(names)

	for _, name := range names {
		version := strings.TrimSuffix(name, ".up.sql")
		var exists bool
		if err := pool.QueryRow(ctx, `SELECT EXISTS (SELECT 1 FROM schema_migrations WHERE version = $1)`, version).Scan(&exists); err != nil {
			return fmt.Errorf("check migration %s: %w", name, err)
		}
		if exists {
			continue
		}
		sqlBytes, err := migrations.Files.ReadFile(name)
		if err != nil {
			return fmt.Errorf("read migration %s: %w", name, err)
		}
		tx, err := pool.Begin(ctx)
		if err != nil {
			return fmt.Errorf("begin migration %s: %w", name, err)
		}
		if _, err = tx.Exec(ctx, string(sqlBytes)); err == nil {
			_, err = tx.Exec(ctx, `INSERT INTO schema_migrations (version) VALUES ($1)`, version)
		}
		if err != nil {
			_ = tx.Rollback(ctx)
			return fmt.Errorf("apply migration %s: %w", name, err)
		}
		if err := tx.Commit(ctx); err != nil {
			return fmt.Errorf("commit migration %s: %w", name, err)
		}
	}
	return pool.Ping(ctx)
}
