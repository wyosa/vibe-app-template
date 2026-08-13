# Dev-образ frontend: hotreload через vite (bun run dev).
# Контекст сборки — корень репозитория.
#
# База — node (не oven/bun): bun-обёртка над node ломает vue-tsc/volar под Linux,
# поэтому скрипты должны исполняться реальным node. bun остаётся пакетным менеджером.
FROM node:26-alpine

RUN npm install -g bun@1.3.14

WORKDIR /app

# Зависимости отдельным слоем для кэша; bun.lock фиксирует версии.
COPY apps/site/package.json apps/site/bun.lock ./
RUN bun install

COPY apps/site/ ./

EXPOSE 5173

# --host 0.0.0.0 — чтобы vite был доступен извне контейнера.
CMD ["bun", "run", "dev", "--", "--host", "0.0.0.0"]
