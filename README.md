# inka

Стартовый шаблон монорепо: Go backend + Vue SPA, spec-driven (OpenAPI), всё в Docker.

## Стек

- **Backend** — Go 1.26, Clean Architecture, PostgreSQL (sqlx + squirrel), миграции golang-migrate
- **Frontend** — Vue 3 (Vapor), Vite, Nuxt UI 4 + Tailwind 4, Pinia, Vue Router, Vitest
- **API-контракт** — OpenAPI (Redocly) → кодоген сервера (oapi-codegen) и клиента (HeyAPI)
- **Окружение** — Docker Compose (dev — hotreload, local — собранные образы), Taskfile

## Требования

- Docker + docker compose
- [Task](https://taskfile.dev) (`brew install go-task`)
- Для локальных проверок и кодогенерации: Go 1.26+, bun, golangci-lint, golang-migrate

## Быстрый старт

```bash
task init     # env-файлы из шаблонов, зависимости, dev-стек
```

- Интерфейс: http://localhost:5173
- API: http://localhost:8080/api/healthz

## Основные команды

| Команда           | Что делает                                               |
| ----------------- | -------------------------------------------------------- |
| `task dev`        | поднять dev-стек (hotreload) и проверить                 |
| `task down`       | остановить dev-стек                                      |
| `task smoke`      | проверить, что стек жив                                  |
| `task check`      | полная самопроверка: fmt + lint + openapi + test + build |
| `task gen`        | перегенерировать код из OpenAPI (сервер + клиент)        |
| `task migrate`    | применить миграции                                       |
| `task seed`       | применить сиды                                           |
| `task --list-all` | все задачи                                               |

## Структура и процесс

См. [ARCHITECTURE.md](./ARCHITECTURE.md) — карта монорепо и описание процесса
(spec-driven → test-driven → реализация). Правила для AI-агентов — [AGENTS.md](./AGENTS.md)
и вложенные `apps/*/AGENTS.md`. Шаблон спеки новой фичи — [docs/specs/feature.md](./docs/specs/feature.md).

## Переменные окружения

- **Compose-стеки**: `deploy/compose/{dev,local}/.env` — создаются из `.env.example` командой
  `task init`, в git не коммитятся.
- **Backend**: `DATABASE_URL` (обязательная), `HTTP_PORT` (по умолчанию 8080), `APP_ENV` —
  задаются в compose-файлах.
- **Frontend**: API ходит на same-origin `/api`; в dev vite проксирует его на backend
  (`API_PROXY_TARGET`, по умолчанию `http://localhost:8080`).
