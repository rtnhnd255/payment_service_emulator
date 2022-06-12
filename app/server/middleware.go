package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func (s *Server) logMiddleware() mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			h.ServeHTTP(w, r)
			log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(startTime))
		})
	}
}
