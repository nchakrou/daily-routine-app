package auth

import (
	"daily-routine-backend/internal/session"
	"daily-routine-backend/pkg/response"
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
	if r.Method != http.MethodPost {
		response.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	h.handleRegister(w, r)

}
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	h.handleLogin(w, r)
}
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	userId, err := session.Get(r, h.db)
	if err != nil {
		response.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	err = session.Delete(w, h.db, userId)
	if err != nil {
		response.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	response.Success(w, "logout successful", http.StatusOK)
}
