package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tomasmiguez/hillchart-server/models"
)

type CreateHillchartInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateHillchart(c *gin.Context) {
	var input CreateHillchartInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"errors": []string{err.Error()}})
		return
	}

	hillchart := models.Hillchart{Name: input.Name, Scopes: []models.Scope{}, Frames: []models.Frame{}}
	if err := models.DB.Select("Name").Create(&hillchart).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"errors": []string{err.Error()}})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": hillchart})
}
