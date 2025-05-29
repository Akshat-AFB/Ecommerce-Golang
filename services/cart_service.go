package services

import (
	"backend-go/models"
	"backend-go/repositories"
	"errors"
	"log"
)

func AddToCart(userID uint, productID uint, quantity int) error {
	product, err := repositories.GetProductByID(productID)
	if err != nil {
		return errors.New("failed to fetch product")
	}
	if product == nil {
		return errors.New("product not found")
	}
	if product.Quantity <= 0 {
		return errors.New("product is out of stock")
	}

	existing, err := repositories.GetCartItem(userID, productID)
	if err != nil {
		return err
	}
	if existing != nil {
		return repositories.UpdateCartItemQuantity(existing.ID, existing.Quantity+quantity)
	}

	newItem := models.CartItem{UserID: userID, ProductID: productID, Quantity: quantity}
	return repositories.InsertCartItem(newItem)
}

func GetCart(userID uint, limit, offset int) ([]models.CartItem, error) {
	return repositories.GetUserCart(userID, limit, offset)
}


func ChangeQuantity(userID, productID uint, quantity int) error {
	item, err := repositories.GetCartItem(userID, productID)
	if err != nil || item == nil {
		return errors.New("cart item not found")
	}
	return repositories.UpdateCartItemQuantity(item.ID, quantity)
}

func RemoveCartItem(userID, cartItemID uint) error {
	// Fetch item by ID to validate ownership
	log.Printf("Removing cart item %d for user %d", cartItemID, userID)
	item, err := repositories.GetCartItem(userID, cartItemID) // optional improvement: add GetCartItemByID()
	if err != nil || item == nil || item.UserID != userID {
		return errors.New("unauthorized or cart item not found")
	}
	return repositories.DeleteCartItem(cartItemID)
}
