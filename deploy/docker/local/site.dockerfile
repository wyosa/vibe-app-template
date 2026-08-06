# Local-образ frontend: сборка статики и раздача через nginx.
# Контекст сборки — корень репозитория.
#
# База — node (не oven/bun): bun-обёртка над node ломает vue-tsc/volar под Linux,
# поэтому type-check должен исполняться реальным node. bun остаётся пакетным менеджером.

FROM node:25-alpine AS build

RUN npm install -g bun@1.3.14

WORKDIR /app

COPY apps/site/package.json apps/site/bun.lock ./
RUN bun install

COPY apps/site/ ./

RUN bun run build

FROM nginx:1.27-alpine

COPY deploy/docker/local/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80
