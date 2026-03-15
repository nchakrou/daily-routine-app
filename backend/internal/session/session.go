package session

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

const cookieName = "token"

func Create(w http.ResponseWriter, db *sql.DB, userId int) error {
	token, err := uuid.NewV4()
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO sessions (user_id, token, expires_at) VALUES (?, ?, ?)", userId, token.String(), time.Now().Add(24*time.Hour))
	if err != nil {
		return err
	}
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    token.String(),
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	})
	return nil
}
func Delete(w http.ResponseWriter, db *sql.DB, userId int) error {
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		HttpOnly: true,
	})
	_, err := db.Exec("DELETE FROM sessions WHERE user_id = ?", userId)
	if err != nil {
		return err
	}
	return nil
}
func Get(r *http.Request, db *sql.DB) (int, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return 0, errors.New("unauthorized")
	}
	var userId int
	var expiresAt time.Time
	if err := db.QueryRow("SELECT user_id, expires_at FROM sessions WHERE token = ?", cookie.Value).Scan(&userId, &expiresAt); err != nil {

		if err == sql.ErrNoRows {
			return 0, errors.New("unauthorized")
		}
		return 0, err
	}
	if expiresAt.Before(time.Now()) {
		return 0, errors.New("session expired")
	}
	return userId, nil
}
