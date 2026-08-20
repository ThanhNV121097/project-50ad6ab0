package migrations

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5"
)

//go:embed *.sql
var sqlFiles embed.FS

func Apply(ctx context.Context, conn *pgx.Conn) error {
	if _, err := conn.Exec(ctx, `create table if not exists schema_migrations (filename text primary key, applied_at timestamptz not null default now())`); err != nil {
		return fmt.Errorf("create schema_migrations: %w", err)
	}

	entries, err := fs.ReadDir(sqlFiles, ".")
	if err != nil {
		return fmt.Errorf("read migrations: %w", err)
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
		if err := applyOne(ctx, conn, name); err != nil {
			return err
		}
	}

	return nil
}

func applyOne(ctx context.Context, conn *pgx.Conn, name string) error {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin migration %s: %w", name, err)
	}
	defer tx.Rollback(ctx)

	var exists bool
	if err := tx.QueryRow(ctx, `select exists (select 1 from schema_migrations where filename=$1)`, name).Scan(&exists); err != nil {
		return fmt.Errorf("check migration %s: %w", name, err)
	}
	if exists {
		return nil
	}

	body, err := sqlFiles.ReadFile(name)
	if err != nil {
		return fmt.Errorf("read migration %s: %w", name, err)
	}
	if _, err := tx.Exec(ctx, string(body)); err != nil {
		return fmt.Errorf("apply migration %s: %w", name, err)
	}
	if _, err := tx.Exec(ctx, `insert into schema_migrations (filename) values ($1)`, name); err != nil {
		return fmt.Errorf("record migration %s: %w", name, err)
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit migration %s: %w", name, err)
	}
	return nil
}
