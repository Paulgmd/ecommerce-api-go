package services

import (
	"ecommerce-api-go/database"
	"ecommerce-api-go/models"
	"strings"
)

// Devuelve todos los productos registrados.
func GetProducts() []models.Product {
	return database.Products
}

// Busca un producto por su ID.
func GetProductByID(id int) (models.Product, bool) {
	for _, product := range database.Products {
		if product.ID == id {
			return product, true
		}
	}

	return models.Product{}, false
}

// Agrega un nuevo producto.
func AddProduct(product models.Product) {
	database.Products = append(database.Products, product)
}

// Elimina un producto por su ID.
func DeleteProductByID(id int) bool {
	for index, product := range database.Products {
		if product.ID == id {
			database.Products = append(database.Products[:index], database.Products[index+1:]...)
			return true
		}
	}

	return false
}

// Actualiza un producto existente.
func UpdateProductByID(id int, updatedProduct models.Product) bool {
	for index, product := range database.Products {
		if product.ID == id {
			database.Products[index] = updatedProduct
			return true
		}
	}

	return false
}

// Cuenta el total de productos registrados.
func CountProducts() int {
	return len(database.Products)
}

// Obtiene los productos que tienen stock disponible.
func GetProductsWithStock() []models.Product {
	var productsWithStock []models.Product

	for _, product := range database.Products {
		if product.Stock > 0 {
			productsWithStock = append(productsWithStock, product)
		}
	}

	return productsWithStock
}

// Obtiene el producto con el precio más alto.
func GetMostExpensiveProduct() models.Product {
	mostExpensive := database.Products[0]

	for _, product := range database.Products {
		if product.Price > mostExpensive.Price {
			mostExpensive = product
		}
	}

	return mostExpensive
}

// Busca productos por nombre.
func SearchProductByName(name string) []models.Product {
	var result []models.Product

	for _, product := range database.Products {
		if strings.Contains(strings.ToLower(product.Name), strings.ToLower(name)) {
			result = append(result, product)
		}
	}

	return result
}

// Devuelve los productos con stock menor o igual al límite indicado.
func GetLowStockProducts(limit int) []models.Product {
	var results []models.Product

	for _, product := range database.Products {
		if product.Stock <= limit {
			results = append(results, product)
		}
	}

	return results
}
