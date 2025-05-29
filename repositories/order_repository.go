package repositories

import (
	"backend-go/database"
	"backend-go/models"
	"time"
)

func GetCartItemsWithProduct(userID uint) ([]models.CartItem, error) {
	rows, err := database.GetDB().Query(`
		SELECT c.product_id, c.quantity, p.id, p.name, p.price, p.description, p.image_url, p.quantity
		FROM cart_items c
		JOIN products p ON c.product_id = p.id
		WHERE c.user_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.CartItem
	for rows.Next() {
		var item models.CartItem
		var product models.Product
		if err := rows.Scan(
			&item.ProductID, &item.Quantity,
			&product.ID, &product.Name, &product.Price, &product.Description, &product.ImageURL, &product.Quantity,
		); err != nil {
			return nil, err
		}
		item.Product = product
		items = append(items, item)
	}
	return items, nil
}

func CreateOrder(userID uint, total float64, items []models.OrderItem) (int, error) {
	tx, err := database.GetDB().Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	now := time.Now()
	var orderID int
	err = tx.QueryRow(`
		INSERT INTO orders (user_id, total, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`,
		userID, total, "placed", now, now).Scan(&orderID)
	if err != nil {
		return 0, err
	}

	for _, item := range items {
		_, err := tx.Exec(`
			INSERT INTO order_items (order_id, product_id, quantity, price) 
			VALUES ($1, $2, $3, $4)`,
			orderID, item.ProductID, item.Quantity, item.Price)
		if err != nil {
			return 0, err
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return orderID, nil
}

func ClearCart(userID uint) error {
	_, err := database.GetDB().Exec("DELETE FROM cart_items WHERE user_id = $1", userID)
	return err
}

func GetFullOrder(orderID int) (*models.Order, error) {
	var order models.Order
	var createdAt, updatedAt time.Time

	err := database.GetDB().QueryRow(`
		SELECT id, user_id, total, status, created_at, updated_at
		FROM orders
		WHERE id = $1`, orderID).
		Scan(&order.ID, &order.UserID, &order.Total, &order.Status, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}
	order.CreatedAt = &createdAt
	order.UpdatedAt = &updatedAt

	// Load order items with product details
	rows, err := database.GetDB().Query(`
		SELECT oi.product_id, oi.quantity, oi.price, 
		       p.id, p.name, p.price, p.description, p.image_url, p.quantity
		FROM order_items oi
		JOIN products p ON oi.product_id = p.id
		WHERE oi.order_id = $1`, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.OrderItem
		var product models.Product
		if err := rows.Scan(
			&item.ProductID, &item.Quantity, &item.Price,
			&product.ID, &product.Name, &product.Price, &product.Description, &product.ImageURL, &product.Quantity,
		); err != nil {
			return nil, err
		}
		item.Product = product
		order.OrderItems = append(order.OrderItems, item)
	}

	// Load user details
	var user models.User
	err = database.GetDB().QueryRow(`
		SELECT id, username, email, role
		FROM users
		WHERE id = $1`, order.UserID).
		Scan(&user.ID, &user.Username, &user.Email, &user.Role)
	if err != nil {
		return nil, err
	}
	order.User = user

	return &order, nil
}

func GetOrdersByUser(userID uint, limit, offset int) ([]models.Order, error) {
	rows, err := database.GetDB().Query(`
		SELECT id, total, status, created_at, updated_at
		FROM orders
		WHERE user_id = $1
		ORDER BY id DESC
		LIMIT $2 OFFSET $3`, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var o models.Order
		var createdAt, updatedAt time.Time
		o.UserID = userID
		if err := rows.Scan(&o.ID, &o.Total, &o.Status, &createdAt, &updatedAt); err != nil {
			return nil, err
		}
		o.CreatedAt = &createdAt
		o.UpdatedAt = &updatedAt

		// Fetch OrderItems with product details
		items, err := fetchOrderItemsWithProducts(uint(o.ID))
		if err != nil {
			return nil, err
		}
		o.OrderItems = items

		// Fetch user details
		var user models.User
		err = database.GetDB().QueryRow(`
			SELECT id, username, email, role
			FROM users
			WHERE id = $1`, o.UserID).
			Scan(&user.ID, &user.Username, &user.Email, &user.Role)
		if err != nil {
			return nil, err
		}
		o.User = user

		orders = append(orders, o)
	}
	return orders, nil
}


func fetchOrderItemsWithProducts(orderID uint) ([]models.OrderItem, error) {
	rows, err := database.GetDB().Query(`
		SELECT oi.product_id, oi.quantity, oi.price, 
		       p.id, p.name, p.price, p.description, p.image_url, p.quantity
		FROM order_items oi
		JOIN products p ON oi.product_id = p.id
		WHERE oi.order_id = $1`, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.OrderItem
	for rows.Next() {
		var item models.OrderItem
		var product models.Product
		if err := rows.Scan(
			&item.ProductID, &item.Quantity, &item.Price,
			&product.ID, &product.Name, &product.Price, &product.Description, &product.ImageURL, &product.Quantity,
		); err != nil {
			return nil, err
		}
		item.Product = product
		items = append(items, item)
	}
	return items, nil
}

func GetOrderByID(orderID int) (*models.Order, error) {
	return GetFullOrder(orderID)
}

func CancelOrder(orderID int) error {
	_, err := database.GetDB().Exec("UPDATE orders SET status = 'cancelled' WHERE id = $1", orderID)
	return err
}
