package services

import (
	"backend-go/models"
	"backend-go/repositories"
)

func GetProductsService(limit, offset int) ([]models.Product, error) {
	return repositories.GetAllProducts(limit, offset)
}
func CreateProductService(product models.Product) (models.Product, error) {
	return repositories.InsertProduct(product)
}
