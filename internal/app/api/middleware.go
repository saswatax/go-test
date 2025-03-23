package api

import (
	"context"
	"net/http"
)

func testMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "message", "hello")

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
