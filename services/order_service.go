package services

import (
	"backend-go/models"
	"backend-go/repositories"
	"errors"
)

func PlaceOrder(userID uint) (*models.Order, error) {
	cartItems, err := repositories.GetCartItemsWithProduct(userID)
	if err != nil {
		return nil, errors.New("failed to retrieve cart items")
	}
	if len(cartItems) == 0 {
		return nil, errors.New("cart is empty")
	}

	var total float64
	var orderItems []models.OrderItem
	for _, item := range cartItems {
		total += float64(item.Quantity) * item.Product.Price
		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Product.Price,
		})
	}

	orderID, err := repositories.CreateOrder(userID, total, orderItems)
	if err != nil {
		return nil, errors.New("failed to place order")
	}

	if err := repositories.ClearCart(userID); err != nil {
		return nil, err
	}

	fullOrder, err := repositories.GetFullOrder(orderID)
	if err != nil {
		return nil, errors.New("failed to load order")
	}

	return fullOrder, nil
}

func ViewOrders(userID uint, limit, offset int) ([]models.Order, error) {
	return repositories.GetOrdersByUser(userID, limit, offset)
}

func CancelOrder(userID uint, orderID int) (*models.Order, error) {
	order, err := repositories.GetOrderByID(orderID)
	if err != nil || order.UserID != userID {
		return nil, errors.New("order not found or unauthorized")
	}
	if err := repositories.CancelOrder(orderID); err != nil {
		return nil, err
	}
	return repositories.GetFullOrder(orderID)
}
