package types

import (
	"github.com/golang-jwt/jwt/v5"
)
type Credentials struct {
	Login    string `json:"login"` // "login" can be email or username
	Password string `json:"password"`
}
type Claims struct {
	UserID uint
    Role   string
    jwt.RegisteredClaims
}
type ChangeQuantityPayload struct {
	Quantity int `json:"quantity"`
}