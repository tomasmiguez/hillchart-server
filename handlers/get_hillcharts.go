package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tomasmiguez/hillchart-server/models"
)

func GetHillcharts(c *gin.Context) {
	var hillcharts []models.Hillchart
	if err := models.DB.Preload("Frames").Find(&hillcharts).Error; err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"errors": []string{err.Error()}})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": hillcharts})
}
