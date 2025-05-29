package models

import(
	"errors"
)
type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string	`gorm:"not null;default:'user'"`
}

// Custom errors
var (
    ErrEmailExists     = errors.New("email already registered")
    ErrUsernameExists  = errors.New("username already taken")
    ErrUserNotFound    = errors.New("user not found")
    ErrInvalidPassword = errors.New("invalid password")
)