# Local-образ backend: статический бинарник в минимальном рантайме.
# Контекст сборки — корень репозитория.

FROM golang:1.26-alpine AS build

WORKDIR /app

COPY apps/api/go.mod apps/api/go.sum ./
RUN go mod download

COPY apps/api/ ./

# CGO выключен — бинарник статический, запускается на голом alpine.
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /out/api ./cmd/api

FROM alpine:3.21

RUN apk add --no-cache ca-certificates tzdata \
  && adduser -D -u 10001 app

COPY --from=build /out/api /usr/local/bin/api

USER app
EXPOSE 8080

ENTRYPOINT ["api"]
