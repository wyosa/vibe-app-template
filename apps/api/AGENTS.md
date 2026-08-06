# AGENTS.md — backend (apps/api)

## Конвенции

- Clean Architecture: зависимости только вниз `delivery → usecase → repository → entity`.
- Новый эндпоинт начинается со спеки (`api/openapi/`), затем `task api:gen`, затем реализация.
  Каталог `internal/delivery/http/gen/` руками не редактировать.
- Доменный слайс на домен: `entity/<домен>.go` (модель), `repository/<домен>.go`
  (запросы, sqlx + squirrel), `usecase/<домен>.go` (бизнес-правила; зависимость на
  репозиторий — через интерфейс, объявленный в usecase), `delivery/http/handlers/<домен>.go`
  (маппинг gen-моделей и ошибок). Wiring — в `internal/app/app.go`, репозиторий —
  в `internal/app/container.go`.
- Ошибки: доменные sentinel-ошибки в `entity/`; репозиторий маппит
  `sql.ErrNoRows` → доменный `ErrNotFound`; хендлер маппит доменные ошибки на HTTP
  через `errors.Is`.
- Миграции: `task migrate:create -- <имя>`, оба файла up/down. Применяются автоматически
  при старте api (см. `migrations/migrator.go`), вручную — `task migrate`.
  Сиды — `apps/api/seeds/seed.sql` (`task seed`), идемпотентные.
- Тесты: usecase — unit с in-memory реализацией репозитория; хендлеры — сквозной
  httptest без БД (strict handler + mux из gen). Ассерты — testify `require`.
- Проверки перед коммитом: `task api:lint`, `task api:test`, `task api:fmt` (или `task check` из корня).
