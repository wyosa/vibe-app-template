# inka

Monorepo starter template: Go backend + Vue SPA, spec-driven (OpenAPI), everything in Docker.

## Stack

- **Backend** — Go 1.26, Clean Architecture, PostgreSQL (sqlx + squirrel), golang-migrate migrations
- **Frontend** — Vue 3 (Vapor), Vite, Nuxt UI 4 + Tailwind 4, Pinia, Vue Router, Vitest
- **API contract** — OpenAPI (Redocly) → server codegen (oapi-codegen) and client codegen (HeyAPI)
- **Environment** — Docker Compose (dev — hotreload, local — built images), Taskfile

## Requirements

- Docker + docker compose
- [Task](https://taskfile.dev) (`brew install go-task`)
- For local checks and codegen: Go 1.26+, bun, golangci-lint, golang-migrate

## Quick start

```bash
task init     # env files from templates, dependencies, dev stack
```

- UI: http://localhost:5173
- API: http://localhost:8080/api/healthz

## Main commands

| Command           | What it does                                         |
| ----------------- | ---------------------------------------------------- |
| `task dev`        | start the dev stack (hotreload) and verify it        |
| `task down`       | stop the dev stack                                   |
| `task smoke`      | check that the stack is alive                        |
| `task check`      | full self-check: fmt + lint + openapi + test + build |
| `task gen`        | regenerate code from OpenAPI (server + client)       |
| `task migrate`    | apply migrations                                     |
| `task seed`       | apply seeds                                          |
| `task --list-all` | all tasks                                            |

## Structure and process

See [ARCHITECTURE.md](./ARCHITECTURE.md) — the monorepo map and the process description
(spec-driven → test-driven → implementation). Rules for AI agents — [AGENTS.md](./AGENTS.md)
and the nested `apps/*/AGENTS.md`. Feature spec template — [docs/specs/feature.md](./docs/specs/feature.md).

## Environment variables

- **Compose stacks**: `deploy/compose/{dev,local}/.env` — created from `.env.example` by
  `task init`, not committed to git.
- **Backend**: `DATABASE_URL` (required), `HTTP_PORT` (default 8080), `APP_ENV` —
  set in compose files.
- **Frontend**: the API goes to same-origin `/api`; in dev vite proxies it to the backend
  (`API_PROXY_TARGET`, default `http://localhost:8080`).
