package dto

import (
	"errors"
	"regexp"
)

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterUser struct {
	UserLogin
	Phone string `json:"phone" validate:"required,phone"`
}

// ValidateEmail validates email format using regex
func (u *UserLogin) ValidateEmail() error {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(emailRegex, u.Email)
	if !matched {
		return errors.New("invalid email format")
	}
	return nil
}

// Validate validates the UserLogin struct
func (u *UserLogin) Validate() error {
	if u.Email == "" {
		return errors.New("email is required")
	}

	if err := u.ValidateEmail(); err != nil {
		return err
	}

	if u.Password == "" {
		return errors.New("password is required")
	}

	if len(u.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	return nil
}

// Validate validates the RegisterUser struct
func (r *RegisterUser) Validate() error {
	if err := r.UserLogin.Validate(); err != nil {
		return err
	}

	if r.Phone == "" {
		return errors.New("phone is required")
	}

	// Simple phone validation (you can make it more specific)
	phoneRegex := `^[\+]?[1-9][\d]{0,15}$`
	matched, _ := regexp.MatchString(phoneRegex, r.Phone)
	if !matched {
		return errors.New("invalid phone format")
	}

	return nil
}
