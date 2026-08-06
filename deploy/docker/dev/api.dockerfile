# Dev-образ backend: hotreload через air.
# Контекст сборки — корень репозитория.
FROM golang:1.26-alpine

RUN go install github.com/air-verse/air@v1.61.7

WORKDIR /app

# Сначала зависимости — слой кэшируется, пока go.mod/go.sum не изменятся.
COPY apps/api/go.mod apps/api/go.sum ./
RUN go mod download

COPY apps/api/.air.toml ./
COPY apps/api/ ./

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
