package auth

import (
	"daily-routine-backend/internal/session"
	"daily-routine-backend/pkg/response"
	"database/sql"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	errV := validateLoginReq(&req)
	if len(errV) > 0 {
		response.MultiError(w, errV, http.StatusBadRequest)
		return
	}
	var password string
	var userId int
	err := h.db.QueryRow("SELECT password, id FROM users WHERE email = ?", req.Email).Scan(&password, &userId)
	if err == sql.ErrNoRows {
		response.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}
	if err != nil {
		response.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(req.Password)); err != nil {
		response.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}
	if err := session.Create(w, h.db, userId); err != nil {
		response.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	response.Success(w, "login successful", http.StatusOK)
}
func validateLoginReq(req *LoginRequest) map[string]string {
	errs := make(map[string]string)
	if req.Email == "" {
		errs["email"] = "email is required"
	} else {
		if !emailRegex.MatchString(req.Email) {
			errs["email"] = "invalid email format"
		}
	}
	if req.Password == "" {
		errs["password"] = "password is required"
	}
	return errs
}
