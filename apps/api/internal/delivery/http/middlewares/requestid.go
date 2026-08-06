package middlewares

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

// RequestIDHeader — заголовок, через который request id приходит в запросе
// и возвращается в ответе.
const RequestIDHeader = "X-Request-Id"

type requestIDKey struct{}

// RequestID проставляет каждому запросу идентификатор: берёт из входящего
// заголовка X-Request-Id или генерирует новый. Кладёт его в контекст запроса
// (см. RequestIDFromContext) и в одноимённый ответный заголовок.
// Должен идти до Logging, чтобы request id попадал в логи.
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get(RequestIDHeader)
		if id == "" {
			id = uuid.NewString()
		}

		w.Header().Set(RequestIDHeader, id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), requestIDKey{}, id)))
	})
}

// RequestIDFromContext возвращает request id, проставленный middleware RequestID,
// или пустую строку, если его нет.
func RequestIDFromContext(ctx context.Context) string {
	id, _ := ctx.Value(requestIDKey{}).(string)
	return id
}
