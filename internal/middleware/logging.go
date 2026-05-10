package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start)

		w.Header().Set("X-Response-Time", duration.String())

		// 2. Usando fmt.Printf para garantir que saia no mesmo canal do "Servidor rodando..."
		fmt.Printf(
			"[LOG] [%s] %s %s | %v\n",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
			time.Since(start),
		)
	})
}
