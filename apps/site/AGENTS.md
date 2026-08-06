# AGENTS.md — frontend (apps/site)

## Conventions

- **nuxt-ui MCP first** (`mcp__nuxt-ui__*`): look for a ready-made Nuxt UI component, check its
  props/slots. Don't hand-roll markup if a ready-made component exists.
- API — only via the generated client `src/types/api` (HeyAPI). Never edit it by hand;
  contract updated → `task site:gen` (or `task gen` from the root). The directory is committed
  to git and not linted (see `eslint.config.ts`).
- The client calls same-origin `/api`: in dev — vite proxy (`API_PROXY_TARGET`), in local — nginx.
  No absolute URLs or CORS needed; don't fetch around the client.
- Pages — `pages/` + registration in `router/`, state — `stores/` (Pinia, setup style),
  reusable pieces — `components/` and `composables/`. Tab title — `usePageTitle`.
- Tests — `src/__tests__/`, Vitest + @vue/test-utils; API calls are mocked
  with `vi.mock('@/types/api')`.
- Pre-commit checks: `bun run lint`, `bun run test:unit run`, `bun run type-check`
  (or `task check` from the root).
