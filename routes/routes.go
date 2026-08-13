package routes

import (
	"Bioskop-API/controllers"

	"github.com/gin-gonic/gin"
)

func Server() *gin.Engine {
	route := gin.Default()

	bioskop := route.Group("/api")
	{
		bioskop.GET("/bioskop", controllers.GetAllBioskop)
		bioskop.GET("/bioskop/:id", controllers.GetBioskopByID)
		bioskop.POST("/bioskop", controllers.CreateBioskop)
		bioskop.PUT("/bioskop/:id", controllers.UpdateBioskop)
		bioskop.DELETE("/bioskop/:id", controllers.DeleteBioskop)
	}

	return route
}
