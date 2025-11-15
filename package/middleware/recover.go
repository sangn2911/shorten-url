package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"shorten-url/package/httputils"
	"shorten-url/package/logger"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if value := recover(); value != nil {
				logger.Errorf(
					fmt.Errorf(`panic "%s"`, value),
					"stack trace:\n%s",
					debug.Stack(),
				)
				httputils.JSON500(w, r, "Something is wrong")
			}
		}()
		next.ServeHTTP(w, r)
	})
}
