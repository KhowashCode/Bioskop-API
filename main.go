package main

import (
	"Bioskop-API/database"
	"Bioskop-API/routes"
	"fmt"
)

func main() {
	const (
		PORT = ":8080"
		URL  = "http://localhost:8080"
	)

	database.DatabaseConnection()
	routes.Server().Run(PORT)
	fmt.Printf("Server Running On %v\n", URL)
}
