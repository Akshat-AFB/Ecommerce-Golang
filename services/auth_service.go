package services

import (
    "backend-go/models"
    "backend-go/repositories"
    "time"
	"backend-go/types"
    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("231d11c697b4a11fed49886a62cf5cc8d50572543beb9ed16a9bd82cbf59a986")

func RegisterUser(user models.User) (models.User, error) {
    if exists, _ := repositories.IsEmailTaken(user.Email); exists {
        return models.User{}, models.ErrEmailExists
    }
    if exists, _ := repositories.IsUsernameTaken(user.Username); exists {
        return models.User{}, models.ErrUsernameExists
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return models.User{}, err
    }

    user.Password = string(hashedPassword)
    if user.Role == "" {
        user.Role = "user"
    }

    return repositories.CreateUser(user)
}

func LoginUser(creds types.Credentials) (string, error) {
    user, err := repositories.FindUserByLogin(creds.Login)
    if err != nil {
        return "", models.ErrUserNotFound
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
        return "", models.ErrInvalidPassword
    }

    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &types.Claims{
        UserID: uint(user.ID),
        Role:   user.Role,
        RegisteredClaims: jwt.RegisteredClaims{
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            ExpiresAt: jwt.NewNumericDate(expirationTime),
            Issuer:    "backend-go",
            Audience:  []string{"users"},
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    go EmitLoginEvent(uint(user.ID))
    return token.SignedString(jwtKey)
}
