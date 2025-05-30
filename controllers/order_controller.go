package controllers

import (
	"backend-go/services"
	"backend-go/middleware"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r.Context())
	order, err := services.PlaceOrder(userID)
	if err != nil {
		if err.Error() == "cart is empty" {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	order.User.Password = ""
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func ViewOrders(w http.ResponseWriter, r *http.Request) {
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

	orders, err := services.ViewOrders(userID, limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for i := range orders {
		orders[i].User.Password = ""
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func GetOrderByID(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r.Context())
	orderIDStr := strings.TrimPrefix(r.URL.Path, "/api/v1/orders/view/")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}
	order, err := services.GetOrderByID(userID, orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	order.User.Password = ""
	json.NewEncoder(w).Encode(order)
}

func CancelOrder(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r.Context())
	orderIDStr := strings.TrimPrefix(r.URL.Path, "/api/v1/orders/cancel/")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}
	order, err := services.CancelOrder(userID, orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	order.User.Password = ""
	json.NewEncoder(w).Encode(order)
}
