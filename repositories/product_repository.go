package repositories

import (
	"backend-go/database"
	"backend-go/models"
	"backend-go/redis"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

func GetProductByID(id uint) (*models.Product, error) {
	cacheKey := fmt.Sprintf("product:%d", id)

	// 1. Try to get from Redis
	cached, err := redis.Get(cacheKey)
	if err == nil {
		var p models.Product
		if jsonErr := json.Unmarshal([]byte(cached), &p); jsonErr == nil {
			return &p, nil
		}
	}

	// 2. If cache miss, query DB
	row := database.GetDB().QueryRow(`
		SELECT id, name, price, description, image_url, quantity
		FROM products
		WHERE id = $1`, id)

	var p models.Product
	if err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.ImageURL, &p.Quantity); err != nil {
		return nil, errors.New("product not found")
	}

	// 3. Cache the result
	data, _ := json.Marshal(p)
	redis.Set(cacheKey, string(data), 10*time.Minute)

	return &p, nil
}

func GetAllProducts(limit, offset int) ([]models.Product, error) {
	rows, err := database.GetDB().Query(`
		SELECT id, name, price, description, image_url, quantity
		FROM products
		ORDER BY id
		LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.ImageURL, &p.Quantity); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

// InsertProduct inserts a new product and returns it with its new ID
func InsertProduct(product models.Product) (models.Product, error) {
	err := database.GetDB().QueryRow(`
		INSERT INTO products (name, price, description, image_url, quantity)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`,
		product.Name, product.Price, product.Description, product.ImageURL, product.Quantity,
	).Scan(&product.ID)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func UpdateProduct(p models.Product) error {
	_, err := database.GetDB().Exec(`
		UPDATE products 
		SET name = $1, price = $2, description = $3, image_url = $4, quantity = $5
		WHERE id = $6`,
		p.Name, p.Price, p.Description, p.ImageURL, p.Quantity, p.ID)

	if err == nil {
		redis.Del(fmt.Sprintf("product:%d", p.ID))
	}
	return err
}

func DeleteProduct(id uint) error {
	_, err := database.GetDB().Exec("DELETE FROM products WHERE id = $1", id)

	if err == nil {
		redis.Del(fmt.Sprintf("product:%d", id))
	}
	return err
}
