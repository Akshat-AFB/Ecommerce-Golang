package models

import(
	"errors"
)
type User struct {
	ID       int   
	Username string 
	Email    string 
	Password string 
	Role     string	
}

var (
    ErrEmailExists     = errors.New("email already registered")
    ErrUsernameExists  = errors.New("username already taken")
    ErrUserNotFound    = errors.New("user not found")
    ErrInvalidPassword = errors.New("invalid password")
)