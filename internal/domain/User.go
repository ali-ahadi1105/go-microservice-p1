package domain

import "time"

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Code      int       `json:"code"`
	Expiry    time.Time `json:"expiry"`
	Verified  bool      `json:"verified"`
	UserType  string    `json:"usertype"`
}
