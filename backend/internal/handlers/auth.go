package auth

import (
	"database/sql"
	"net/http"
)

type AuthHandler struct {
	db *sql.DB
}

func New(db *sql.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {

}
