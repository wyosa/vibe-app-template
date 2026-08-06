# <Feature name>

<!-- Template for a new feature spec. Copy to docs/specs/<name>.md and fill it in. -->

## Goal

What and why we're doing it, 1–2 sentences. Which problem it solves.

## Contract (OpenAPI)

Changes in `api/openapi/`:

- Endpoints: `METHOD /api/...` — what they do, response codes
- Models: fields, requiredness, constraints
- Errors: 400/404/… — in which cases

## Database

- Migration (`task migrate:create -- <name>`): tables, columns, indexes
- Seeds (`apps/api/seeds/seed.sql`): which test data

## Backend

Layers (see conventions in `apps/api/AGENTS.md`):

- `entity/` — domain model and errors
- `repository/` — queries (sqlx + squirrel), `sql.ErrNoRows` → domain `ErrNotFound`
- `usecase/` — business rules, repository dependency via an interface
- `delivery/http/handlers/` — mapping of gen models and errors to HTTP
- Tests: usecase — unit with an in-memory repository; handler — end-to-end httptest without a DB

## Frontend

- Page + route, store (Pinia), Nuxt UI components (look for ready-made ones via nuxt-ui MCP first)
- API calls — only via the generated client `src/types/api`
- Tests: store/component with a mocked client

## Definition of done

- [ ] Spec updated, `task api:contract` is green
- [ ] Codegen regenerated and committed (`task gen`)
- [ ] Tests written before the implementation and green (`task test`)
- [ ] `task check` is green
- [ ] Verified in the browser (playwright): the page opens, forms work, console is clean
- [ ] Commit in Conventional Commits format
