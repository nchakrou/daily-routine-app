package server

import (
	"database/sql"
	"net/http"
)

type Server struct {
	db  *sql.DB
	mux *http.ServeMux
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
func New(db *sql.DB) *Server {
	serv := &Server{
		db:  db,
		mux: http.NewServeMux(),
	}
	serv.Router()
	return serv
}
