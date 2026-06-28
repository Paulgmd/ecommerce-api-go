package services

import "ecommerce-api-go/models"

// Interface que define las operaciones del servicio de productos
type ProductService interface {
	GetProducts() []models.Product
	GetProductByID(id int) (models.Product, bool)
	AddProduct(product models.Product)
	DeleteProduct(id int) bool
	UpdateProduct(id int, product models.Product) bool
}
