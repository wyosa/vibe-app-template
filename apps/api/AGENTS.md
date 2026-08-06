# AGENTS.md — backend (apps/api)

## Conventions

- Clean Architecture: dependencies only downward `delivery → usecase → repository → entity`.
- A new endpoint starts with the spec (`api/openapi/`), then `task api:gen`, then the implementation.
  Never edit the `internal/delivery/http/gen/` directory by hand.
- One domain slice per domain: `entity/<domain>.go` (model), `repository/<domain>.go`
  (queries, sqlx + squirrel), `usecase/<domain>.go` (business rules; the repository dependency
  is an interface declared in the usecase), `delivery/http/handlers/<domain>.go`
  (mapping of gen models and errors). Wiring — in `internal/app/app.go`, the repository —
  in `internal/app/container.go`.
- Errors: domain sentinel errors in `entity/`; the repository maps
  `sql.ErrNoRows` → domain `ErrNotFound`; handlers map domain errors to HTTP
  via `errors.Is`.
- Migrations: `task migrate:create -- <name>`, both up and down files. Applied automatically
  at api startup (see `migrations/migrator.go`), manually — `task migrate`.
  Seeds — `apps/api/seeds/seed.sql` (`task seed`), idempotent.
- Tests: usecase — unit with an in-memory repository implementation; handlers — end-to-end
  httptest without a DB (strict handler + mux from gen). Assertions — testify `require`.
- Pre-commit checks: `task api:lint`, `task api:test`, `task api:fmt` (or `task check` from the root).
