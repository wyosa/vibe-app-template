# Architecture

Монорепо с Go-бэкендом и Vue-фронтендом.

```
.
├── apps/
│   ├── api/          # Go backend
│   └── site/         # Vue frontend
├── api/
│   └── openapi/      # OpenAPI-спецификация (multi-file, Redocly)
├── deploy/
│   ├── docker/
│   │   ├── dev/      # Dockerfile для разработки (hotreload)
│   │   └── local/    # Dockerfile для локального запуска (сборка)
│   └── compose/
│       ├── dev/      # docker-compose: air + bun run dev
│       └── local/    # docker-compose: бинарник + статика
├── docs/
│   └── specs/        # шаблоны спек фич (spec-driven процесс)
├── scripts/
├── Taskfile.yaml     # корневые задачи
└── redocly.yaml
```

## Процесс разработки

Spec-driven → test-driven → реализация.

1. **Spec** — изменение начинается с OpenAPI-спецификации (`api/openapi/`). Контракт валидируется Redocly. Из спеки генерируется серверный код (`delivery/http/gen/`) и клиентский (`apps/site/src/types/api/`, HeyAPI). Шаблон спеки фичи — `docs/specs/feature.md`.
2. **Tests** — под новый контракт пишутся тесты до реализации. Go — `go test`, фронтенд — Vitest.
3. **Code** — реализуется бизнес-логика (usecase → repository), хендлеры подключаются к сгенерированному интерфейсу. На клиенте — страницы и сторы поверх сгенерированных типов и функций.

## Backend — `apps/api`

Go 1.26, Clean Architecture. Зависимости потока: delivery → usecase → repository → entity.

```
apps/api/
├── cmd/api/              # точка входа
├── internal/
│   ├── app/              # композиция (DI-контейнер, запуск)
│   ├── config/           # конфигурация из env
│   ├── delivery/http/
│   │   ├── gen/          # сгенерированный код из OpenAPI
│   │   ├── handlers/     # HTTP-хендлеры (объединяются в handlers.API)
│   │   ├── middlewares/    # recovery + request-id + request-logging
│   │   └── server/       # HTTP-сервер
│   ├── entity/           # доменные модели
│   ├── repository/       # работа с PostgreSQL (sqlx + squirrel)
│   └── usecase/          # бизнес-логика
├── migrations/           # SQL-миграции (golang-migrate)
├── pkg/                  # переиспользуемые пакеты (logger и т.д.)
└── Taskfile.yaml
```

### Data store

PostgreSQL. Доступ через `sqlx` + `squirrel` (placeholder `$1`). Транзакции — обёртка `sqlxTransaction` в `internal/repository/transaction.go`. Миграции — `golang-migrate`, лежат в `migrations/`.

### API-контракт

OpenAPI-спецификация в `api/openapi/` (multi-file: `paths/`, `components/`, `tags.yaml`). Валидация и сборка — Redocly (`redocly.yaml`). Кодогенерация — в `internal/delivery/http/gen/`.

Все пути — с префиксом `/api`. Фронтенд ходит **same-origin**: в dev vite проксирует `/api`
на backend (`API_PROXY_TARGET`), в local — nginx. CORS не нужен, абсолютные URL не используются.

## Frontend — `apps/site`

Vue 3 (Vapor), Vite 8, TypeScript.

| Слой       | Технология                  |
| ---------- | --------------------------- |
| UI-кит     | Nuxt UI 4                   |
| Стили      | Tailwind CSS 4              |
| Стейт      | Pinia 4                     |
| Роутинг    | Vue Router 5                |
| Линтер     | oxlint + eslint             |
| Форматтер  | oxfmt                       |
| API-клиент | HeyAPI (codegen из OpenAPI) |
| Тесты      | Vitest + @vue/test-utils    |

Сгенерированный клиент (`src/types/api/`) коммитится в git и не линтится; перегенерация —
`task site:gen`, дрейф ловит CI.

```
apps/site/src/
├── components/
├── composables/
├── layouts/
├── pages/
├── router/
├── stores/
├── types/
│   └── api/          # сгенерированный HeyAPI-клиент
├── utils/
└── assets/
```

## Deploy

### `deploy/compose/dev` — разработка

Hotreload: **air** для Go, **bun run dev** для фронтенда. Dockerfile в `deploy/docker/dev/`.

### `deploy/compose/local` — локальный запуск

Собранный Go-бинарник + статика фронтенда. Dockerfile в `deploy/docker/local/`.
