package server

import (
	"daily-routine-backend/internal/session"
	"daily-routine-backend/pkg/response"
	"net/http"
)

func (s *Server) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := session.Get(r, s.db)
		if err != nil {
			response.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})

}
