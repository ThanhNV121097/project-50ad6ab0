# Architecture Overview — hello-word

## Scope

`hello-word` is a fullstack one-screen product. PostgreSQL stores one public message row, Go serves it, and Next.js renders it. No auth, navigation, admin editing, analytics, or extra pages.

## Stack

| Layer | Choice | Reason | Rejected alternative |
|---|---|---|---|
| Frontend | Next.js 15 App Router, TypeScript, Tailwind v3 | Matches repo containers and gives server-rendered page shell | Plain HTML would not fit committed frontend container or later story layout |
| Backend | Go 1.22+ HTTP API | Small binary, existing container expects Go module under `code/backend/` | Node API would add second JS runtime and break backend Dockerfile convention |
| Database | PostgreSQL 16 | Required by SRS: message stored in one row | Frontend hardcoded text rejected by core requirement |
| CI | GitHub Actions in `.github/workflows/ci.yml` | Fast build/lint/test gate before review | Relying only on container workflow hides lint and unit failures |

## Repository layout

```text
code/
  backend/
    cmd/api/main.go
    internal/migrations/migrations.go
    migrations/*.sql
  frontend/
    app/layout.tsx
    app/page.tsx
    app/globals.css
    components/
    lib/
docs/architecture/overview.md
```

`code/backend/` has exactly one main package: `./cmd/api`. `code/frontend/` uses App Router. Story code mounts into `app/page.tsx` with one import and one element; page stays composition root.

## Runtime data flow

1. Backend starts, reads `DATABASE_URL`, applies migrations from embedded SQL files, verifies database with `SELECT 1`, then listens on `PORT`, `APP_PORT`, or `8080`.
2. `/healthz` returns 200 only after migrations succeed and database ping works.
3. Future message endpoint reads exactly one stored row and returns plain JSON.
4. Frontend reads `NEXT_PUBLIC_API_URL` for browser calls and renders only backend-provided message.

## Environment variables

### Root compose

| Key | Used by | Purpose |
|---|---|---|
| `POSTGRES_USER` | postgres, backend DSN | Local database user |
| `POSTGRES_PASSWORD` | postgres, backend DSN | Local database password |
| `POSTGRES_DB` | postgres, backend DSN | Local database name |
| `BACKEND_PORT` | compose | Host port for backend |
| `FRONTEND_PORT` | compose | Host port for frontend |
| `NEXT_PUBLIC_API_URL` | frontend build | Browser-visible backend URL |

### Backend

| Key | Required | Purpose |
|---|---|---|
| `DATABASE_URL` | yes | PostgreSQL URL injected by runtime |
| `PORT` | no | HTTP port, preferred |
| `APP_PORT` | no | HTTP port fallback |

### Frontend

| Key | Required | Purpose |
|---|---|---|
| `NEXT_PUBLIC_API_URL` | yes | Backend base URL visible to browser |

## Naming conventions

| Thing | Convention |
|---|---|
| Go packages | lowercase short names |
| Go entrypoint | `cmd/api/main.go` |
| Migrations | `YYYYMMDDHHMMSS_description.up.sql` and `.down.sql` |
| React components | `export default function ComponentName()` |
| CSS tokens | `--category-name`, defined in `app/globals.css` |
| API paths | `/healthz` for health, `/api/...` for product data |

## Frontend boundaries

`app/page.tsx` remains a Server Component. Files using hooks, browser APIs, or event handlers must start with literal first line `"use client"`. Shared visual values live in `app/globals.css`; story component CSS must use tokens, no literal colors and no token fallbacks.

## Backend boundaries

Startup owns migrations. Product handlers must not create schema. SQL uses parameterized queries. External errors returned to frontend stay generic. Empty or missing message states are handled by product story code, not scaffold.

## Failure handling

| Failure | Behavior |
|---|---|
| Missing `DATABASE_URL` | Backend exits with clear startup error |
| Migration fails | Backend exits unhealthy |
| Database unavailable after boot | `/healthz` returns 503 |
| Frontend cannot reach backend | Feature UI shows empty/error state when implemented |

## Observability

Backend logs startup, migration result, listen address, and health query failures to stdout/stderr. No personal data exists. Frontend keeps default Next.js build/runtime logs.

## Run locally

```bash
cp .env.example .env
cp code/backend/.env.example code/backend/.env
cp code/frontend/.env.example code/frontend/.env.local
docker compose --profile local up --build
```

Frontend: `http://localhost:3000`. Backend health: `http://localhost:8080/healthz`.

## Local checks

```bash
cd code/backend && go build ./... && go vet ./... && go test ./...
cd code/frontend && npm ci && npm run lint && npm run build && npm test --if-present
docker compose config -q
```

## Decisions and tradeoffs

| Decision | Why | Rejected |
|---|---|---|
| Self-migrate backend on boot | Runtime creates empty DB; no separate migrator exists | Manual migrations would make first run fail |
| Keep Dockerfiles and `docker-compose.yml` convention-compatible | Orchestrator-provided container gates assume these paths | Rewriting containers risks breaking known-good pipeline |
| Put migrations embed under `internal/migrations` | `//go:embed` resolves relative to source file | Embedding from `cmd/api` would look under `cmd/api/migrations` |
| Seed only schema in scaffold | Feature story owns product row and endpoint | Implementing message now would bypass story flow |

## Unknowns

Detailed table shape and API contract belong to ERD and service design tasks. Current scaffold only proves build, migration plumbing, database health, and runtime layout.
