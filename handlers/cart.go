package handlers

import (
    "encoding/json"
    "net/http"
    "backend-go/models"
    "backend-go/database"
	"backend-go/middleware"
	"backend-go/types"
    "strings"
    "github.com/golang-jwt/jwt/v5"
	"strconv"
	
)

type AddToCartInput struct {
    ProductID uint `json:"product_id"`
    Quantity  int  `json:"quantity"`
}

func AddToCart(w http.ResponseWriter, r *http.Request) {
    var input AddToCartInput
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Get user from token
    tokenStr := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
    claims := &Claims{}
    _, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte("231d11c697b4a11fed49886a62cf5cc8d50572543beb9ed16a9bd82cbf59a986"), nil
    })
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    db := database.GetDB()
	// Check product availability
	var product models.Product
	if err := db.First(&product, input.ProductID).Error; err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if product.Quantity <= 0 {
		http.Error(w, "Product is out of stock", http.StatusBadRequest)
		return
	}
    var existing models.CartItem
    result := db.Where("user_id = ? AND product_id = ?", claims.UserID, input.ProductID).First(&existing)

    if result.RowsAffected > 0 {
        // Update quantity
        existing.Quantity += input.Quantity
        db.Save(&existing)
    } else {
        // Add new cart item
        item := models.CartItem{
            UserID:    claims.UserID,
            ProductID: input.ProductID,
            Quantity:  input.Quantity,
        }
        db.Create(&item)
    }

    json.NewEncoder(w).Encode(map[string]string{"message": "Item added to cart"})
}

func ViewCart(w http.ResponseWriter, r *http.Request) {
    tokenStr := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
    claims := &Claims{}
    _, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte("231d11c697b4a11fed49886a62cf5cc8d50572543beb9ed16a9bd82cbf59a986"), nil
    })
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    db := database.GetDB()

    var cartItems []models.CartItem
    db.Preload("Product").Where("user_id = ?", claims.UserID).Find(&cartItems)

    if len(cartItems) == 0 {
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"message": "Your cart is empty"})
        return
    }

    json.NewEncoder(w).Encode(cartItems)
}

func RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/cart/remove/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid cart item ID", http.StatusBadRequest)
		return
	}

	userID := middleware.GetUserIDFromContext(r.Context())

	var item models.CartItem
	if err := database.DB.First(&item, id).Error; err != nil {
		http.Error(w, "Cart item not found", http.StatusNotFound)
		return
	}

	if item.UserID != userID {
		http.Error(w, "Unauthorized to delete this item", http.StatusForbidden)
		return
	}

	database.DB.Delete(&item)
	w.WriteHeader(http.StatusNoContent)
}

func ChangeCartQuantity(w http.ResponseWriter, r *http.Request) {
	// Extract product ID from URL path
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 {
		http.Error(w, "Product ID missing in URL", http.StatusBadRequest)
		return
	}
	productIDStr := pathParts[3]
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Parse new quantity from body
	var payload types.ChangeQuantityPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	if payload.Quantity <= 0 {
		http.Error(w, "Quantity must be greater than zero", http.StatusBadRequest)
		return
	}

	// Get user ID from context
	userID := middleware.GetUserIDFromContext(r.Context())
	if userID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	// Find the cart item
	var cartItem models.CartItem
	err = database.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&cartItem).Error
	if err != nil {
		http.Error(w, "Cart item not found", http.StatusNotFound)
		return
	}

	// Update the quantity
	cartItem.Quantity = payload.Quantity
	if err := database.DB.Save(&cartItem).Error; err != nil {
		http.Error(w, "Failed to update cart item", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Cart quantity updated",
		"item":    cartItem,
	})
}