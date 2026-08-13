package controllers

import (
	"Bioskop-API/database"
	"Bioskop-API/models"
	"Bioskop-API/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBioskop(c *gin.Context) {
	var bioskop *models.Bioskop
	var allBioskop []models.Bioskop

	if err := c.ShouldBindJSON(&bioskop); err != nil {
		utils.Fail(c, http.StatusBadRequest, "400", "Invalid request format / Invalid data")
		return
	}

	if bioskop.Name == "" && bioskop.Location == "" {
		utils.Fail(c, http.StatusInternalServerError, "500", "Failed to process data")
	} else {
		newBioskop := []*models.Bioskop{
			{
				Name:     bioskop.Name,
				Location: bioskop.Location,
				Rating:   bioskop.Rating,
			},
		}

		database.DB.Create(&newBioskop)
		database.DB.Last(&allBioskop)

		utils.OK(c, http.StatusCreated, allBioskop)
	}
}
