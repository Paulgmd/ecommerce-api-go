package services

import (
	"ecommerce-api-go/database"
	"ecommerce-api-go/models"
)

func GetProducts() []models.Product {
	return database.Products
}

func GetProductByID(id int) (models.Product, bool) {
	for _, product := range database.Products {
		if product.ID == id {
			return product, true
		}
	}

	return models.Product{}, false
}

func AddProduct(product models.Product) {
	database.Products = append(database.Products, product)
}
func DeleteProductByID(id int) bool {
	for index, product := range database.Products {
		if product.ID == id {
			database.Products = append(database.Products[:index], database.Products[index+1:]...)
			return true
		}
	}

	return false
}
func UpdateProductByID(id int, updatedProduct models.Product) bool {
	for index, product := range database.Products {
		if product.ID == id {
			database.Products[index] = updatedProduct
			return true
		}
	}

	return false
}
func CountProducts() int {
	return len(database.Products)
}

func GetProductsWithStock() []models.Product {
	var productsWithStock []models.Product

	for _, product := range database.Products {
		if product.Stock > 0 {
			productsWithStock = append(productsWithStock, product)
		}
	}

	return productsWithStock
}

func GetMostExpensiveProduct() models.Product {
	mostExpensive := database.Products[0]

	for _, product := range database.Products {
		if product.Price > mostExpensive.Price {
			mostExpensive = product
		}
	}

	return mostExpensive
}
