package controllers

import (
    "encoding/json"
    "net/http"
    "strings"

    "backend-go/models"
    "backend-go/services"
	"backend-go/types"
)

func Register(w http.ResponseWriter, r *http.Request) {
    var user models.User

    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Password) == "" {
        http.Error(w, "Email, Username, and Password are required fields", http.StatusBadRequest)
        return
    }

    createdUser, err := services.RegisterUser(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(createdUser)
}

func Login(w http.ResponseWriter, r *http.Request) {
    var creds types.Credentials

    if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    token, err := services.LoginUser(creds)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"token": token})
}
