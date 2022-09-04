package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tomasmiguez/hillchart-server/models"
)

type CreateStatusInput struct {
	HillchartID uint `json:"hillchart_id" binding:"required"`
}

func CreateStatus(c *gin.Context) {
	var input CreateStatusInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"errors": []string{err.Error()}})
		return
	}

	status := models.Status{HillchartID: input.HillchartID}
	if err := models.DB.Create(&status).Error; err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"errors": []string{err.Error()}})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": status})
}
