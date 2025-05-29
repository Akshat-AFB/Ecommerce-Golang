package controllers

import (
	"encoding/json"
	// "log"
	"net/http"
	"strconv"
	"backend-go/models"
	"backend-go/services"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	pageParam := r.URL.Query().Get("page")
	limitParam := r.URL.Query().Get("limit")

	page := 1
	limit := 10

	if pageParam != "" {
		if p, err := strconv.Atoi(pageParam); err == nil && p > 0 {
			page = p
		}
	}
	if limitParam != "" {
		if l, err := strconv.Atoi(limitParam); err == nil && l > 0 {
			limit = l
		}
	}

	offset := (page - 1) * limit

	products, err := services.GetProductsService(limit, offset)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if len(products) == 0 {
		json.NewEncoder(w).Encode(map[string]string{"message": "No products found"})
		return
	}

	json.NewEncoder(w).Encode(products)
}


func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if product.Name == "" || product.Price <= 0 || product.Description == "" || product.ImageURL == "" {
		http.Error(w, "Missing or invalid fields", http.StatusBadRequest)
		return
	}
	if product.Quantity < 0 {
		http.Error(w, "Quantity cannot be negative", http.StatusBadRequest)
		return
	}

	createdProduct, err := services.CreateProductService(product)
	if err != nil {
		http.Error(w, "Failed to create product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdProduct)
}
