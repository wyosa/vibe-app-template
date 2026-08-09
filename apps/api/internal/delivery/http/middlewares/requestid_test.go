package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRequestID_GeneratesAndPropagates(t *testing.T) {
	t.Parallel()

	var gotFromCtx string
	next := http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		gotFromCtx = RequestIDFromContext(r.Context())
	})

	req := httptest.NewRequestWithContext(t.Context(), http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	RequestID(next).ServeHTTP(rec, req)

	id := rec.Header().Get(RequestIDHeader)
	require.NotEmpty(t, id)
	require.Equal(t, id, gotFromCtx)
}

func TestRequestID_KeepsIncomingHeader(t *testing.T) {
	t.Parallel()

	next := http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		require.Equal(t, "incoming-id", RequestIDFromContext(r.Context()))
	})

	req := httptest.NewRequestWithContext(t.Context(), http.MethodGet, "/", nil)
	req.Header.Set(RequestIDHeader, "incoming-id")
	rec := httptest.NewRecorder()

	RequestID(next).ServeHTTP(rec, req)

	require.Equal(t, "incoming-id", rec.Header().Get(RequestIDHeader))
}
