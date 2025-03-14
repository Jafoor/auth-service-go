package middlewares

import (
	"log/slog"
	"net/http"
	"time"

	"auth-service/logger"
)

type recorder struct {
	http.ResponseWriter
	statusCode int
	length     int
}

func (r *recorder) Write(data []byte) (int, error) {
	sz, err := r.ResponseWriter.Write(data)
	r.length += sz
	return sz, err
}

func (r *recorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func (r *recorder) Flush() {
	if flusher, ok := r.ResponseWriter.(http.Flusher); ok {
		flusher.Flush()
	}
}

func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		rec := &recorder{
			ResponseWriter: w,
		}
		start := time.Now()

		handler.ServeHTTP(rec, r)

		// ignore healthcheck route
		if path == "/hello" {
			return
		}

		slog.Info(
			"",
			logger.Path(path),
			logger.Query(r.URL.Query()),
			logger.Method(r.Method),
			logger.Status(rec.statusCode),
			logger.UserAgent(r.UserAgent()),
			logger.Ip(r.RemoteAddr),
			logger.Latency(time.Since(start)),
			logger.Length(rec.length),
		)
	})
}
