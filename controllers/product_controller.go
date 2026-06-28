package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"ecommerce-api-go/models"
	"ecommerce-api-go/services"
)

func GetProductsController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products := services.GetProducts()

	json.NewEncoder(w).Encode(products)
}

func GetProductByIDController(w http.ResponseWriter, r *http.Request) {

	id := 1

	product, found := services.GetProductByID(id)

	if !found {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func AddProductController(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	body, _ := io.ReadAll(r.Body)

	var product models.Product

	json.Unmarshal(body, &product)

	services.AddProduct(product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(product)
}
func DeleteProductController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	id := 1

	deleted := services.DeleteProductByID(id)

	if !deleted {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Producto eliminado correctamente"))
}
func UpdateProductController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	id := 2

	body, _ := io.ReadAll(r.Body)

	var product models.Product
	json.Unmarshal(body, &product)

	updated := services.UpdateProductByID(id, product)

	if !updated {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
func CountProductsController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	total := services.CountProducts()

	json.NewEncoder(w).Encode(map[string]int{
		"total": total,
	})
}

func GetProductsWithStockController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products := services.GetProductsWithStock()

	json.NewEncoder(w).Encode(products)
}

func GetMostExpensiveProductController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	product := services.GetMostExpensiveProduct()

	json.NewEncoder(w).Encode(product)
}
func SearchProductController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	name := r.URL.Query().Get("name")

	products := services.SearchProductByName(name)

	json.NewEncoder(w).Encode(products)
}

func GetLowStockProductsController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products := services.GetLowStockProducts(30)

	json.NewEncoder(w).Encode(products)
}
