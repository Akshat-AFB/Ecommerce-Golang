package handlers

import(
	"net/http"
	"encoding/json"
	"time"
	// "fmt"
	"strings"
	// "log"
	// "os"

	"golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v5"

	"backend-go/models"
	"backend-go/database"
)

var jwtKey = []byte("231d11c697b4a11fed49886a62cf5cc8d50572543beb9ed16a9bd82cbf59a986")

type Credentials struct {
	Login    string `json:"login"` // "login" can be email or username
	Password string `json:"password"`
}
type Claims struct {
	UserID uint
    Role   string
    jwt.RegisteredClaims
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	
	// Decode JSON body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Password) == "" {
		http.Error(w, "Email, Username, and Password are required fields", http.StatusBadRequest)
		return
	}

	// Check if email already exists
	var existingUser models.User
	if err := database.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		http.Error(w, "Email already registered", http.StatusConflict)
		return
	}

	// Check if username already exists
	if err := database.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		http.Error(w, "Username already taken", http.StatusConflict)
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Set default role
	if user.Role == "" {
		user.Role = "user"
	}

	// Create user
	result := database.DB.Create(&user)
	if result.Error != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}


func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// fmt.Print(creds)


	var user models.User
	result := database.DB.Where("email = ? OR username = ?", creds.Login, creds.Login).First(&user)

	if result.Error != nil {
		http.Error(w, "User not found. Please register first.", http.StatusUnauthorized)
		return
	}

	// fmt.Print(result)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	//Generate JWT token
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: uint(user.ID),
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer: "backend-go",
			Audience: []string{"users"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Error creating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}