package routes

import (
	"net/http"

	"ecommerce-api-go/controllers"
)

func RegisterRoutes() {
	http.HandleFunc("/products", controllers.GetProductsController)
	http.HandleFunc("/products/one", controllers.GetProductByIDController)
	http.HandleFunc("/products/add", controllers.AddProductController)
	http.HandleFunc("/products/delete", controllers.DeleteProductController)
	http.HandleFunc("/products/update", controllers.UpdateProductController)

	http.HandleFunc("/products/count", controllers.CountProductsController)
	http.HandleFunc("/products/stock", controllers.GetProductsWithStockController)
	http.HandleFunc("/products/expensive", controllers.GetMostExpensiveProductController)
}
