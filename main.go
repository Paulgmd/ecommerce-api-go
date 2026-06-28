package main

import (
	"fmt"
	"net/http"

	"ecommerce-api-go/routes"
)

func main() {
	routes.RegisterRoutes()

	fmt.Println("Servidor actualizado con ruta PUT")

	http.ListenAndServe(":8080", nil)
}
