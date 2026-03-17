package auth

import (
	"daily-routine-backend/internal/models"
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

func (h *AuthHandler) UserInfo(w http.ResponseWriter, r *http.Request) {
	userId, err := session.Get(r, h.db)
	if err != nil {
		response.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	var res models.User
	err = h.db.QueryRow("SELECT id, email, username, avatar, about FROM users WHERE id = ?", userId).Scan(&res.ID, &res.Email, &res.Username, &res.Avatar, &res.About)
	if err != nil {
		response.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	response.JSON(w, res, http.StatusOK)
}
