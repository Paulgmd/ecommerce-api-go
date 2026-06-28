package database

import "ecommerce-api-go/models"

var Products = []models.Product{
	{
		ID:          1,
		Name:        "Laptop",
		Description: "Laptop Lenovo",
		Price:       1200.50,
		Stock:       10,
	},
	{
		ID:          2,
		Name:        "Mouse",
		Description: "Mouse Logitech",
		Price:       25.99,
		Stock:       50,
	},
}
