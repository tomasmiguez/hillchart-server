package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tomasmiguez/hillchart-server/models"
)

type CreateFrameInput struct {
	HillchartID string `json:"hillchart_id" binding:"required"`
}

func CreateFrame(c *gin.Context) {
	var input CreateFrameInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"errors": []string{err.Error()}})
		return
	}

	frame := models.Frame{HillchartID: input.HillchartID, FrameScopes: []models.FrameScope{}}
	if err := models.DB.Select("HillchartID").Create(&frame).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"errors": []string{err.Error()}})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": frame})
}
