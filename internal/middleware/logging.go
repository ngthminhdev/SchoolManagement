package middleware

import (
	"GolangBackend/helper"
	"net/http"
	"time"
)

func HttpLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		helper.LogInfo(
			"%s %s %dms\n",
			r.Method,
			r.RequestURI,
			time.Since(start).Milliseconds(),
		)
	})
}
