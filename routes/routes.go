package routes

import (
	"Bioskop-API/controllers"

	"github.com/gin-gonic/gin"
)

func Server() *gin.Engine {
	route := gin.Default()

	bioskop := route.Group("/api/bioskop")
	{
		bioskop.POST("/create", controllers.CreateBioskop)
	}

	return route
}
