package controllers

import (
	"encoding/json"
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
// UpdateProduct handles updating a product
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if product.ID == 0 || product.Name == "" || product.Price <= 0 || product.Description == "" || product.ImageURL == "" || product.Quantity < 0 {
		http.Error(w, "Missing or invalid fields", http.StatusBadRequest)
		return
	}

	err := services.UpdateProductService(product)
	if err != nil {
		http.Error(w, "Failed to update product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product updated successfully"})
}

// DeleteProduct handles deleting a product
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		ID uint `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil || payload.ID == 0 {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err := services.DeleteProductService(payload.ID)
	if err != nil {
		http.Error(w, "Failed to delete product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted successfully"})
}

