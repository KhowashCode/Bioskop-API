package main

import (
	"Bioskop-API/database"
	"Bioskop-API/routes"
	"os"
)

func main() {
	database.DatabaseConnection()
	routes.Server().Run(":" + os.Getenv("PGPORT"))
}
