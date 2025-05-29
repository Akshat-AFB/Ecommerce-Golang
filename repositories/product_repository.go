package repositories

import (
	"backend-go/database"
	"backend-go/models"
)

// GetAllProducts retrieves all products from the database
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

// UpdateProduct updates an existing product
func UpdateProduct(p models.Product) error {
	_, err := database.GetDB().Exec(`
		UPDATE products 
		SET name = $1, price = $2, description = $3, image_url = $4, quantity = $5
		WHERE id = $6`,
		p.Name, p.Price, p.Description, p.ImageURL, p.Quantity, p.ID)
	return err
}

// DeleteProduct deletes a product by its ID
func DeleteProduct(id uint) error {
	_, err := database.GetDB().Exec("DELETE FROM products WHERE id = $1", id)
	return err
}

// GetProductByID retrieves a single product by its ID
// func GetProductByID(id uint) (*models.Product, error) {
// 	row := database.GetDB().QueryRow(`
// 		SELECT id, name, price, description, image_url, quantity 
// 		FROM products WHERE id = $1`, id)

// 	var p models.Product
// 	if err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.ImageURL, &p.Quantity); err != nil {
// 		return nil, err
// 	}
// 	return &p, nil
// }
