# AGENTS.md — frontend (apps/site)

## Конвенции

- **Сначала nuxt-ui MCP** (`mcp__nuxt-ui__*`): ищи готовый компонент Nuxt UI, проверяй его
  props/слоты. Самописную вёрстку не делай, если есть готовый компонент.
- API — только через сгенерированный клиент `src/types/api` (HeyAPI). Руками не редактировать;
  обновился контракт → `task site:gen` (или `task gen` из корня). Каталог коммитится в git
  и не линтится (см. `eslint.config.ts`).
- Клиент ходит на same-origin `/api`: в dev — vite proxy (`API_PROXY_TARGET`), в local — nginx.
  Абсолютные URL и CORS не нужны, fetch в обход клиента не использовать.
- Страницы — `pages/` + регистрация в `router/`, стейт — `stores/` (Pinia, setup-стиль),
  переиспользуемое — `components/` и `composables/`. Заголовок вкладки — `usePageTitle`.
- Тесты — `src/__tests__/`, Vitest + @vue/test-utils; вызовы API мокаются
  `vi.mock('@/types/api')`.
- Проверки перед коммитом: `bun run lint`, `bun run test:unit run`, `bun run type-check`
  (или `task check` из корня).
