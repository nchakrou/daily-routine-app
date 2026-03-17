package server

import (
	auth "daily-routine-backend/internal/handlers/auth"
	"net/http"
)

func (s *Server) Router() {
	authHandler := auth.New(s.db)
	s.mux.HandleFunc("/api/auth/register", authHandler.Register)
	s.mux.HandleFunc("/api/auth/login", authHandler.Login)
	s.mux.HandleFunc("/api/auth/refresh", authHandler.Logout)

	s.mux.Handle("/api/auth/userinfo", http.HandlerFunc(authHandler.UserInfo))
}
