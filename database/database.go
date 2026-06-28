package database

import "ecommerce-api-go/models"

var Products = []models.Product{
	{ID: 1, Name: "Laptop Lenovo", Description: "Laptop empresarial Lenovo", Price: 1200.50, Stock: 10},
	{ID: 2, Name: "Mouse Logitech", Description: "Mouse inalámbrico Logitech", Price: 25.99, Stock: 50},
	{ID: 3, Name: "Teclado Mecánico RGB", Description: "Teclado gamer con luces RGB", Price: 75.50, Stock: 20},
	{ID: 4, Name: "Monitor Samsung 27", Description: "Monitor Full HD de 27 pulgadas", Price: 299.99, Stock: 12},
	{ID: 5, Name: "Monitor LG UltraWide", Description: "Monitor panorámico para productividad", Price: 349.99, Stock: 9},
	{ID: 6, Name: "SSD Samsung 1TB", Description: "Disco sólido de alta velocidad", Price: 129.99, Stock: 40},
	{ID: 7, Name: "Memoria RAM Corsair 16GB", Description: "Memoria DDR4 para alto rendimiento", Price: 79.99, Stock: 35},
	{ID: 8, Name: "Procesador Ryzen 7", Description: "Procesador AMD Ryzen 7", Price: 420.00, Stock: 8},
	{ID: 9, Name: "Tarjeta Gráfica RTX 4070", Description: "GPU NVIDIA para gaming y diseño", Price: 699.99, Stock: 5},
	{ID: 10, Name: "Webcam Logitech Brio", Description: "Cámara web 4K profesional", Price: 169.99, Stock: 15},
	{ID: 11, Name: "Audífonos HyperX", Description: "Auriculares gamer con micrófono", Price: 89.90, Stock: 30},
	{ID: 12, Name: "Silla Gamer Secretlab", Description: "Silla ergonómica para escritorio", Price: 249.99, Stock: 7},
}
