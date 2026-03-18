package auth

import (
	"daily-routine-backend/pkg/response"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

type RegisterRequest struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	DateOfBirth string `json:"dateOfBirth"`
	Username    string `json:"username,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
	About       string `json:"about,omitempty"`
}

func (h *AuthHandler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, "bad request", http.StatusBadRequest)
		log.Println("bad request", err)
		return
	}
	log.Println("request", req)
	err := validateRegisterReq(&req)
	if err != nil {
		response.MultiError(w, err, http.StatusBadRequest)
		log.Println("validation error", err)
		return
	}
	if req.Username == "" {
		req.Username = req.FirstName + " " + req.LastName
	}
	if req.Avatar == "" {
		req.Avatar = "default.png"
	}
	exists := false
	log.Println("email", req.Email)
	if err := h.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", req.Email).Scan(&exists); err != nil {
		response.Error(w, "internal server error", http.StatusInternalServerError)
		log.Printf("register: check email exists: %v", err)
		return
	}
	if exists {
		response.Error(w, "email already exists", http.StatusBadRequest)
		log.Println("email already exists")
		return
	}
	hashedPassword, errr := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if errr != nil {
		response.Error(w, "internal server error", http.StatusInternalServerError)
		log.Println(" bcrypt error", errr)
		return
	}
	if _, err := h.db.Exec("INSERT INTO users (firstname, lastname, email, password, date_of_birth, username, avatar, about_me) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", req.FirstName, req.LastName, req.Email, hashedPassword, req.DateOfBirth, req.Username, req.Avatar, req.About); err != nil {
		response.Error(w, "failed to register user", http.StatusInternalServerError)
		log.Println(" db error", err)
		return
	}
	response.Success(w, "user registered successfully", http.StatusCreated)

}
func validateRegisterReq(req *RegisterRequest) map[string]string {
	errs := make(map[string]string)
	if req.FirstName == "" {
		errs["firstName"] = "first name is required"
	}
	if req.LastName == "" {
		errs["lastName"] = "last name is required"
	}
	if req.Email == "" {
		errs["email"] = "email is required"
	} else {
		if !emailRegex.MatchString(req.Email) {
			errs["email"] = "invalid email format"
		}
	}
	if req.Password == "" {
		errs["password"] = "password is required"
	} else {
		if len(req.Password) < 8 {
			errs["password"] = "password must be at least 8 characters long"
		}
	}
	if req.DateOfBirth == "" {
		errs["dateOfBirth"] = "date of birth is required"

	} else {
		if _, err := time.Parse("2006-01-02", req.DateOfBirth); err != nil {
			errs["dateOfBirth"] = "invalid format — use YYYY-MM-DD"
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}
