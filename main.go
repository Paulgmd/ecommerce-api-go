package main

import (
	"fmt"
	"net/http"

	"ecommerce-api-go/routes"
)

func main() {
	routes.RegisterRoutes()

	fmt.Println("SERVIDOR VERSION NUEVA")
	http.ListenAndServe(":8080", nil)
}
