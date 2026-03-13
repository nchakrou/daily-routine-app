package server

import auth "daily-routine-backend/internal/handlers"

func (s *Server) Router() {
	authHandler := auth.New(s.db)
	s.mux.HandleFunc("/api/auth/register", authHandler.Register)
}
