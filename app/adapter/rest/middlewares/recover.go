package middlewares

import (
	"auth-service/logger"
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
)

func Recover(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				slog.Error(
					fmt.Sprintf("Recovered from panic: %v", err),
					logger.Extra(map[string]any{
						"stack": string(debug.Stack()),
					}),
				)
				http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)

}
