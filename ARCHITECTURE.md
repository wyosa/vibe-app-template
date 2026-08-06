# Architecture

Monorepo with a Go backend and a Vue frontend.

```
.
├── apps/
│   ├── api/          # Go backend
│   └── site/         # Vue frontend
├── api/
│   └── openapi/      # OpenAPI specification (multi-file, Redocly)
├── deploy/
│   ├── docker/
│   │   ├── dev/      # Dockerfile for development (hotreload)
│   │   └── local/    # Dockerfile for local run (build)
│   └── compose/
│       ├── dev/      # docker-compose: air + bun run dev
│       └── local/    # docker-compose: binary + static files
├── docs/
│   └── specs/        # feature spec templates (spec-driven process)
├── scripts/
├── Taskfile.yaml     # root tasks
└── redocly.yaml
```

## Development process

Spec-driven → test-driven → implementation.

1. **Spec** — a change starts with the OpenAPI spec (`api/openapi/`). The contract is validated by Redocly. Server code (`delivery/http/gen/`) and client code (`apps/site/src/types/api/`, HeyAPI) are generated from the spec. Feature spec template — `docs/specs/feature.md`.
2. **Tests** — tests for the new contract are written before the implementation. Go — `go test`, frontend — Vitest.
3. **Code** — business logic is implemented (usecase → repository), handlers are wired to the generated interface. On the client — pages and stores on top of the generated types and functions.

## Backend — `apps/api`

Go 1.26, Clean Architecture. Dependency flow: delivery → usecase → repository → entity.

```
apps/api/
├── cmd/api/              # entry point
├── internal/
│   ├── app/              # composition (DI container, startup)
│   ├── config/           # configuration from env
│   ├── delivery/http/
│   │   ├── gen/          # code generated from OpenAPI
│   │   ├── handlers/     # HTTP handlers (combined into handlers.API)
│   │   ├── middlewares/  # recovery + request-id + request-logging
│   │   └── server/       # HTTP server
│   ├── entity/           # domain models
│   ├── repository/       # PostgreSQL access (sqlx + squirrel)
│   └── usecase/          # business logic
├── migrations/           # SQL migrations (golang-migrate)
├── pkg/                  # reusable packages (logger etc.)
└── Taskfile.yaml
```

### Data store

PostgreSQL. Access via `sqlx` + `squirrel` (`$1` placeholders). Transactions — the `sqlxTransaction` wrapper in `internal/repository/transaction.go`. Migrations — `golang-migrate`, live in `migrations/`.

### API contract

OpenAPI specification in `api/openapi/` (multi-file: `paths/`, `components/`, `tags.yaml`). Validation and bundling — Redocly (`redocly.yaml`). Codegen — into `internal/delivery/http/gen/`.

All paths have the `/api` prefix. The frontend calls **same-origin**: in dev vite proxies `/api`
to the backend (`API_PROXY_TARGET`), in local — nginx. No CORS needed, no absolute URLs.

## Frontend — `apps/site`

Vue 3 (Vapor), Vite 8, TypeScript.

| Layer      | Technology                    |
| ---------- | ----------------------------- |
| UI kit     | Nuxt UI 4                     |
| Styles     | Tailwind CSS 4                |
| State      | Pinia 4                       |
| Routing    | Vue Router 5                  |
| Linter     | oxlint + eslint               |
| Formatter  | oxfmt                         |
| API client | HeyAPI (codegen from OpenAPI) |
| Tests      | Vitest + @vue/test-utils      |

The generated client (`src/types/api/`) is committed to git and is not linted; regeneration —
`task site:gen`, drift is caught by CI.

```
apps/site/src/
├── components/
├── composables/
├── layouts/
├── pages/
├── router/
├── stores/
├── types/
│   └── api/          # generated HeyAPI client
├── utils/
└── assets/
```

## Deploy

### `deploy/compose/dev` — development

Hotreload: **air** for Go, **bun run dev** for the frontend. Dockerfile in `deploy/docker/dev/`.

### `deploy/compose/local` — local run

Built Go binary + frontend static files. Dockerfile in `deploy/docker/local/`.
