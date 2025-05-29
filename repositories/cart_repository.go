package repositories

import (
	"backend-go/database"
	"backend-go/models"
	"database/sql"
	"errors"
)
func GetProductByID(id uint) (*models.Product, error) {
	row := database.GetDB().QueryRow(
		"SELECT id, name, price, description, image_url, quantity FROM products WHERE id = $1",
		id,
	)
	var p models.Product
	err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.ImageURL, &p.Quantity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func GetCartItem(userID uint, productID uint) (*models.CartItem, error) {
	row := database.GetDB().QueryRow(
		"SELECT id, user_id, product_id, quantity FROM cart_items WHERE user_id = $1 AND product_id = $2",
		userID, productID,
	)
	var item models.CartItem
	err := row.Scan(&item.ID, &item.UserID, &item.ProductID, &item.Quantity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &item, nil
}

func InsertCartItem(item models.CartItem) error {
	_, err := database.GetDB().Exec(
		"INSERT INTO cart_items (user_id, product_id, quantity) VALUES ($1, $2, $3)",
		item.UserID, item.ProductID, item.Quantity,
	)
	return err
}

func UpdateCartItemQuantity(id uint, quantity int) error {
	_, err := database.GetDB().Exec(
		"UPDATE cart_items SET quantity = $1 WHERE id = $2",
		quantity, id,
	)
	return err
}

func GetUserCart(userID uint, limit, offset int) ([]models.CartItem, error) {
	rows, err := database.GetDB().Query(`
		SELECT c.id, c.user_id, c.product_id, c.quantity,
		       p.id, p.name, p.price, p.description, p.image_url, p.quantity
		FROM cart_items c
		JOIN products p ON c.product_id = p.id
		WHERE c.user_id = $1
		ORDER BY c.id
		LIMIT $2 OFFSET $3`, userID, limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cart []models.CartItem
	for rows.Next() {
		var item models.CartItem
		var p models.Product
		err = rows.Scan(&item.ID, &item.UserID, &item.ProductID, &item.Quantity,
			&p.ID, &p.Name, &p.Price, &p.Description, &p.ImageURL, &p.Quantity)
		if err != nil {
			return nil, err
		}
		item.Product = p
		cart = append(cart, item)
	}

	return cart, nil
}


func DeleteCartItem(id uint) error {
	_, err := database.GetDB().Exec(
		"DELETE FROM cart_items WHERE id = $1",
		id,
	)
	return err
}
