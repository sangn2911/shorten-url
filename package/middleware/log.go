package middleware

import (
	"net/http"
	"shorten-url/package/logger"
	"strings"

	"github.com/google/uuid"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/swagger/") {
			next.ServeHTTP(w, r)
			return
		}
		pid := uuid.NewString()
		ctx := logger.SetPID(r.Context(), pid)
		logger.WithContext(ctx).Infof("[%s]%s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
