package handlers

import (
	"backend-go/database"
	"backend-go/models"
	"encoding/json"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	database.GetDB().Find(&products)
	// products := []models.Product{
	// 	{ID: 1, Name: "iPhone 13", Price: 999.99, Description: "Latest Apple iPhone", ImageURL: "https://example.com/iphone13.jpg"},
	// 	{ID: 2, Name: "Samsung S22", Price: 899.99, Description: "Latest Samsung Galaxy", ImageURL: "https://example.com/samsung-s22.jpg"},
	// 	{ID: 3, Name: "Google Pixel 6", Price: 799.99, Description: "Latest Google Pixel", ImageURL: "https://example.com/pixel6.jpg"},
	// }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	// Decode JSON body
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Manual strict validation
	if product.Name == "" || product.Price <= 0 || product.Description == "" || product.ImageURL == "" {
		http.Error(w, "Missing or invalid fields: Name, Price (> 0), Description, and ImageURL are required", http.StatusBadRequest)
		return
	}

	// Optionally: Quantity must be >= 0
	if product.Quantity < 0 {
		http.Error(w, "Quantity cannot be negative", http.StatusBadRequest)
		return
	}

	// Create product
	result := database.GetDB().Create(&product)
	if result.Error != nil {
		http.Error(w, "Database error: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Send success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
