# AGENTS.md

## Project architecture

At the start of any task, read [`ARCHITECTURE.md`](./ARCHITECTURE.md) — it is the project map: monorepo structure, stack (Go backend + Vue frontend), where the code lives and how it is connected, the development process (spec-driven → test-driven → implementation). Rely on it before defining the blast radius of a change.

## Comprehensive approach to changes

When the user asks for any change, approach the task comprehensively:

1. **Define the blast radius.** A change may affect more than the obvious place. Check whether changes are needed in the backend (`apps/api`, `api/`), in other apps (`apps/*`), in the specs (`api/openapi`) and docs (`docs/`).
2. **Trace the connections.** Before editing, find all the places in the system (UI, API, DB, scripts) the change affects and account for them — so the system stays consistent and nothing breaks.
3. **Verify the result.** After the change, make sure the affected parts work together (run the checks/tests if they exist).

Don't limit yourself to a point fix — carry the task through in every affected part of the system.

## MCP tools

MCP servers are connected in the project — use them instead of guessing:

- **nuxt-ui** (`mcp__nuxt-ui__*`) — for any frontend change, check here first: whether a ready-made Nuxt UI component exists, what it is called and which props/slots it has. Use ready-made components instead of hand-rolled markup.
- **context7** (`mcp__context7__*`) — up-to-date library documentation (Nuxt UI, Vue, Go, etc.). Consult it when unsure about an API or syntax instead of answering from memory.
- **playwright** (`mcp__playwright__*`) — browser control. See the rule below.

## Mandatory browser verification

**After completing any work** (frontend and backend alike — an API change is reflected in the UI too), always test the result via **playwright** before telling the user the task is done:

1. Open the pages affected by the change in the browser.
2. Check that the page loads without errors (including console errors), the changes are actually displayed, and forms and buttons work.
3. If the change touches the backend — verify through the UI that data arrives and saves correctly.
4. If something is broken — fix it and test again. The task is considered done only after a successful browser check.

If the local servers are not running, start them first, then test. If browser verification is impossible — tell the user plainly; don't pass the change off as verified.

## Environment: Docker only

All services run **only via Docker Compose** (Taskfile). The user does not work with Docker directly — that is your responsibility.

- **Forbidden** to run processes locally: no `bun run dev`, `go run`, `vite`, `npm run dev` etc. on the host. Everything — only inside containers via `task`.
- The stack is started with **`task dc:dev:up`** (not `task:dev:up` — there is no such task). Before any interaction with the web UI, check that the stack is up: `task dc:dev:ps`. If not — run `task dc:dev:up` and wait for readiness. Initial setup after cloning — `task init`.
- Addresses: UI (CRM) — `http://localhost:5173`, API — `http://localhost:8080` (paths prefixed with `/api`, the frontend calls same-origin via vite proxy). In playwright always use these addresses.
- After `dc:dev:up` make sure everything is alive: `task smoke`. If something doesn't respond — check `task dc:dev:logs` and fix until it works.
- No manual restarts needed: sources are mounted into the containers, frontend (vite) and backend (air) rebuild themselves on file changes. If the behavior is suspicious — `task dc:dev:restart`.
- If migrations/seed data are needed: `task migrate`, `task seed` (they run inside the dev stack).
- Before committing, run the full self-check: **`task check`** (fmt + lint + openapi + test + build). After changing the spec — `task gen` (codegen is committed together with the spec). Feature spec template — `docs/specs/feature.md`. `apps/api/` and `apps/site/` have their own AGENTS.md with conventions — follow them.

## Conventional Commits

**After every completed change** (once the change is made and verified in the browser), immediately make a git commit in [Conventional Commits](https://www.conventionalcommits.org/) format:

- Format: `<type>(<scope>): <description>`, e.g. `feat(crm): add order filter by status` or `fix(api): fix total amount calculation`.
- Types: `feat` — new feature, `fix` — bug fix, `refactor` — refactoring, `docs` — documentation, `chore` — housekeeping, `style` — styles/markup without logic.
- The format is enforced automatically: the husky commit-msg hook rejects commits outside Conventional Commits.
- One commit = one logical change; don't mix different tasks in one commit.
- Commit only verified code — don't commit broken changes.
- Don't `git push` without an explicit request from the user.

This is needed so that any change can be easily found and rolled back (`git revert` / `git reset`) if something goes wrong.
