package httpapi

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"example.com/gopracs2-z4-borisovda/internal/metrics"
)

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		metrics.ActiveRequests.Inc()
		defer metrics.ActiveRequests.Dec()

		lrw := NewLoggingResponseWriter(w)
		next.ServeHTTP(lrw, r)

		path := NormalizePath(r.URL.Path)
		duration := time.Since(start).Seconds()

		metrics.HttpRequestsTotal.WithLabelValues(r.Method, path).Inc()
		metrics.HttpRequestDuration.WithLabelValues(r.Method, path).Observe(duration)

		if lrw.StatusCode() >= 400 {
			metrics.HttpErrorsTotal.WithLabelValues(r.Method, path, strconv.Itoa(lrw.StatusCode())).Inc()
		}
	})
}

func NormalizePath(path string) string {
	if strings.HasPrefix(path, "/students/") {
		return "/students/{id}"
	}
	return path
}