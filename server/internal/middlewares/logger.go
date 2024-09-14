package middlewares

import (
	"log"
	"net/http"

	"github.com/felixge/httpsnoop"
)

func WithLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(next, w, r)
		log.Printf("http[%d]-- %s -- %s", m.Code, r.Method, r.URL.Path)
	})
}
