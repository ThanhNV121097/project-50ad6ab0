# Architecture Overview — hello-word

## Scope

`hello-word` is fullstack: Next.js frontend, Go backend, PostgreSQL database. Product shows one stored message, centered. No navigation, auth, editing, or extra pages.

## Stack

| Part | Choice | Reason | Rejected alternative |
|---|---|---|---|
| Frontend | Next.js 15 App Router, TypeScript, Tailwind v3 | Matches project default and Docker runtime | Static HTML rejected because frontend must fetch backend data |
| Backend | Go 1.22 HTTP server | Small API, fast build, existing container convention | Node backend rejected to avoid second runtime pattern |
| Database | PostgreSQL 16 | Required source of truth for message row | Frontend hardcode rejected by SRS |
| Styling | CSS tokens + Tailwind available | Plain centered screen, token gate stays enforceable | Component-only hardcoded values rejected because review churn |
| Runtime | `docker compose --profile local up --build` | Boots DB, backend, frontend together | Separate service commands rejected for local drift |

## Repository layout

```text
code/
  backend/
    cmd/api/main.go
    internal/migrations/
    migrations/
    .env.example
  frontend/
    app/layout.tsx
    app/page.tsx
    app/globals.css
    components/
    lib/mock/
    .env.example
docs/
  architecture/overview.md
  hello-word/SRS.md
```

## Backend contract

- Entrypoint: `code/backend/cmd/api/main.go`.
- Reads `DATABASE_URL` and `PORT`; `APP_PORT` fallback allowed before default `8080`.
- Applies pending SQL migrations before serving HTTP.
- Tracks applied migrations in `schema_migrations`.
- `/healthz` returns 200 only after migrations succeed and `SELECT 1` works.
- Future message API belongs in backend story; scaffold exposes health only.

## Frontend contract

- App Router under `code/frontend/app`.
- `app/page.tsx` stays composition root: imports grouped at top, children listed in body.
- Story components use `export default function ComponentName()`.
- Client components must begin with literal first line `"use client"` if using browser APIs, React state/effects/refs, or event handlers.
- `globals.css` owns shared tokens; story CSS modules must use `var(--token)` with no fallback.

## Persistence

Schema details live in future ERD. Scaffold only proves migrations run. Product seed and message table design come in story technical design.

## Env vars

### Root compose

| Key | Purpose |
|---|---|
| `POSTGRES_USER` | Local PostgreSQL user |
| `POSTGRES_PASSWORD` | Local PostgreSQL password |
| `POSTGRES_DB` | Local PostgreSQL database |
| `BACKEND_PORT` | Host port for backend |
| `FRONTEND_PORT` | Host port for frontend |
| `NEXT_PUBLIC_API_URL` | Browser-visible backend base URL |

### Backend

| Key | Purpose |
|---|---|
| `DATABASE_URL` | PostgreSQL connection string |
| `PORT` | HTTP port |
| `APP_PORT` | Legacy fallback HTTP port |

### Frontend

| Key | Purpose |
|---|---|
| `NEXT_PUBLIC_API_URL` | Browser-visible backend base URL |

## Run

```bash
cp .env.example .env
docker compose --profile local up --build
```

Open `http://localhost:3000`. Backend health: `http://localhost:8080/healthz`.

## Checks

```bash
cd code/backend && go build ./... && go vet ./... && go test ./...
cd code/frontend && npm ci && npm run lint && npm run build && npm test --if-present
docker compose config -q
```

CI workflow intended path: `.github/workflows/ci.yml`. Current repository permission blocks agent writes under `.github`; container workflows remain precommitted.

## Naming conventions

| Item | Convention |
|---|---|
| Go packages | lowercase, no hyphen |
| Go exported names | PascalCase only when needed outside package |
| React components | PascalCase filename and `export default function` |
| CSS tokens | `--category-name` in kebab case |
| Env vars | UPPER_SNAKE_CASE |
| Migrations | `YYYYMMDDHHMMSS_name.up.sql` and `.down.sql` |

## Security and reliability

- No secrets committed; examples list keys only.
- DB access uses PostgreSQL driver with context timeouts.
- Health checks include database connectivity.
- Migrations are idempotent through `schema_migrations`.
- No user input paths in scaffold.

## Decisions

| Decision | Tradeoff |
|---|---|
| Self-migrate on backend boot | Simpler runtime; startup can fail if SQL bad, which is desired |
| Keep page empty shell now | Avoids hardcoded product feature in scaffold; story owns rendering |
| Use Dockerfiles already present | Less drift with platform; not personalized |
| Keep database profile as `local` | Deploy uses external DB; local command needs profile |

## Unknowns

- Message table and seed rules finalized in ERD/service design.
- CI custom workflow cannot be added until repository permission allows `.github/workflows/ci.yml` write.
