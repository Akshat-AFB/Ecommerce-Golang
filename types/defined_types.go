package types

import (
	"github.com/golang-jwt/jwt/v5"
)
type Credentials struct {
	Login    string `json:"login"` 
	Password string `json:"password"`
}
type Claims struct {
	UserID uint
    Role   string
    jwt.RegisteredClaims
}
type ChangeQuantityPayload struct {
	ProductID uint `json:"productID"`
	Quantity int `json:"quantity"`
}
