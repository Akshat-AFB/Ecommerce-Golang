package controllers

import (
	"backend-go/services"
	"backend-go/types"
	// "context"
	"encoding/json"
	// "log"
	"net/http"
	"strconv"
	"strings"
	"backend-go/middleware"
)
type contextKey string

const (
	ContextUserID contextKey = "userID"
	ContextRole   contextKey = "role"
)
// func getUserIDFromContext(ctx context.Context) uint {
// 	log.Printf("getUserIDFromContext called with context: %v", ctx)
// 	userID, ok := ctx.Value(ContextUserID).(uint)
// 	log.Print("userID from getUserIDFromContext: ", userID)
// 	log.Print("ok from getUserIDFromContext: ", ok)
// 	if !ok {
// 		return 0
// 	}
// 	return userID
// }

func AddToCart(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r.Context())
	if userID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var input types.ChangeQuantityPayload
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := services.AddToCart(userID, input.ProductID, input.Quantity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Item added to cart"})
}

func ViewCart(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r.Context())
	if userID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Parse pagination params
	pageParam := r.URL.Query().Get("page")
	limitParam := r.URL.Query().Get("limit")

	page := 1
	limit := 10

	if p, err := strconv.Atoi(pageParam); err == nil && p > 0 {
		page = p
	}
	if l, err := strconv.Atoi(limitParam); err == nil && l > 0 {
		limit = l
	}
	offset := (page - 1) * limit

	// Fetch paginated cart
	cart, err := services.GetCart(userID, limit, offset)
	if err != nil || len(cart) == 0 {
		json.NewEncoder(w).Encode(map[string]string{"message": "Your cart is empty"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}


func RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r.Context())
	if userID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/cart/remove/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = services.RemoveCartItem(userID, uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func ChangeQuantity(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r.Context())
	if userID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	productIDStr := strings.TrimPrefix(r.URL.Path, "/api/v1/cart/change/")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var payload types.ChangeQuantityPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	if payload.Quantity <= 0 {
		http.Error(w, "Quantity must be > 0", http.StatusBadRequest)
		return
	}

	err = services.ChangeQuantity(userID, uint(productID), payload.Quantity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Quantity updated"})
}
