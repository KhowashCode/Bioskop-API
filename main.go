package main

import (
	"Bioskop-API/database"
	"Bioskop-API/routes"
	"os"
)

func main() {
	database.DatabaseConnection()

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	routes.Server().Run(":" + port)
}
