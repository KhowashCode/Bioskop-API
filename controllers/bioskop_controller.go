package controllers

import (
	"Bioskop-API/database"
	"Bioskop-API/models"
	"Bioskop-API/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBioskop(c *gin.Context) {
	var allBioskop []models.Bioskop

	if err := database.DB.Find(&allBioskop).Error; err != nil {
		utils.Fail(c, http.StatusNotFound, "404", "Data not found")
		return
	}

	utils.OK(c, http.StatusOK, allBioskop)
}

func GetBioskopByID(c *gin.Context) {
	bioskopId := c.Param("id")
	var bioskop models.Bioskop

	if err := database.DB.First(&bioskop, bioskopId).Error; err != nil {
		utils.Fail(c, http.StatusNotFound, "404", "Data not found")
		return
	}

	utils.OK(c, http.StatusOK, bioskop)
}

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

func UpdateBioskop(c *gin.Context) {
	bioskopId := c.Param("id")
	var newDataBioskop *models.Bioskop
	var oldDataBioskop []*models.Bioskop
	var bioskop models.Bioskop

	if err := c.ShouldBindJSON(&newDataBioskop); err != nil {
		utils.Fail(c, http.StatusBadRequest, "400", "Invalid request format / Invalid data")
		return
	}

	database.DB.Model(&oldDataBioskop).Where("id = ?", bioskopId).Updates(models.Bioskop{Name: newDataBioskop.Name, Location: newDataBioskop.Location, Rating: newDataBioskop.Rating})
	database.DB.First(&bioskop, bioskopId)

	utils.OK(c, http.StatusCreated, bioskop)
}

func DeleteBioskop(c *gin.Context) {
	bioskopId := c.Param("id")
	var bioskop models.Bioskop

	if err := database.DB.First(&bioskop, bioskopId).Error; err != nil {
		utils.Fail(c, http.StatusNotFound, "404", err.Error())
		return
	}

	if err := database.DB.Delete(&bioskop).Error; err != nil {
		utils.Fail(c, http.StatusInternalServerError, "500", "Failed to delete data")
		return
	}

	utils.OK(c, http.StatusNoContent, "Data deleted successfully")
}
