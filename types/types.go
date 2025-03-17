package types

import (
	"errors"
	"regexp"
	"strings"
	"time"
	"unicode"
)

type SignUpUserPayload struct {
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type SignUpUser struct {
	Id        int       `json:"_id" db:"id"`
	FirstName string    `json:"firstName" db:"first_name"`
	LastName  string    `json:"lastName" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

func (u SignUpUserPayload) Validate() error {
	// Trim whitespace from fields
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
	u.Email = strings.TrimSpace(u.Email)
	u.Password = strings.TrimSpace(u.Password)
	u.ConfirmPassword = strings.TrimSpace(u.ConfirmPassword)

	// Check required fields
	if u.FirstName == "" || u.LastName == "" || u.Email == "" || u.Password == "" || u.ConfirmPassword == "" {
		return errors.New("first name, last name, email, password, and confirm password are required")
	}

	// Validate email format
	if !isValidEmail(u.Email) {
		return errors.New("invalid email format")
	}

	// Validate password strength
	if !isValidPassword(u.Password) {
		return errors.New("password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one number, and one special character")
	}

	// Check if password and confirm password match
	if u.Password != u.ConfirmPassword {
		return errors.New("password and confirm password do not match")
	}

	// Validate name length
	if len(u.FirstName) < 2 || len(u.FirstName) > 50 {
		return errors.New("first name must be between 2 and 50 characters")
	}
	if len(u.LastName) < 2 || len(u.LastName) > 50 {
		return errors.New("last name must be between 2 and 50 characters")
	}

	return nil
}

func (u SignUpUser) Validate() error {
	// Validate email format
	if !isValidEmail(u.Email) {
		return errors.New("invalid email format")
	}

	if u.CreatedAt.IsZero() {
		return errors.New("created_at is required")
	}

	if u.UpdatedAt.IsZero() {
		return errors.New("updated_at is required")
	}

	if u.UpdatedAt.Before(u.CreatedAt) {
		return errors.New("updated_at cannot be before created_at")
	}

	return nil
}

type SignInUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u SignInUserPayload) Validate() error {
	// Trim whitespace from fields
	u.Email = strings.TrimSpace(u.Email)
	u.Password = strings.TrimSpace(u.Password)

	// Check required fields
	if u.Email == "" || u.Password == "" {
		return errors.New("email and password are required")
	}

	// Validate email format
	if !isValidEmail(u.Email) {
		return errors.New("invalid email format")
	}

	return nil
}

func isValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}

func isValidPassword(password string) bool {
	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	if len(password) < 8 {
		return false
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSpecial
}
