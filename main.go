package main

import (
	"fmt"
	"net/http"

	"ecommerce-api-go/routes"
)

func main() {
	routes.RegisterRoutes()

	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fs)

	fmt.Println("SERVIDOR VERSION NUEVA")
	http.ListenAndServe(":8080", nil)
}
