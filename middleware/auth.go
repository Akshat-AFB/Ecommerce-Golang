package middleware

import (
	"backend-go/types"
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	ContextUserID contextKey = "userID"
	ContextRole   contextKey = "role"
)

var jwtKey = []byte("231d11c697b4a11fed49886a62cf5cc8d50572543beb9ed16a9bd82cbf59a986")
func AdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		role := GetUserRoleFromContext(r.Context())
		if role != "admin" {
			http.Error(w, "Forbidden: Admins only", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" || !strings.HasPrefix(tokenHeader, "Bearer ") {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}
		tokenStr := strings.TrimPrefix(tokenHeader, "Bearer ")

		claims := &types.Claims{}
        token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }
		log.Printf("Authenticated user ID: %d, Role: %s", claims.UserID, claims.Role)
		// Add user info to request context
		ctx := context.WithValue(r.Context(), ContextUserID, claims.UserID)
		ctx = context.WithValue(ctx, ContextRole, claims.Role)

		// Pass modified request with context to next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
// Helpers
func GetUserIDFromContext(ctx context.Context) uint {
	if userID, ok := ctx.Value(ContextUserID).(uint); ok {
		return userID
	}
	return 0
}

func GetUserRoleFromContext(ctx context.Context) string {
	if role, ok := ctx.Value(ContextRole).(string); ok {
		return role
	}
	return ""
}