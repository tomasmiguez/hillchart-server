package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/tomasmiguez/hillchart-server/models"
)

type CreateHillchartInput struct {
	Name string `json:"name" binding:"required"`
	UUID string `json:"uuid"`
}

func CreateHillchart(c *gin.Context) {
	var input CreateHillchartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"errors": []string{err.Error()}})
		return
	}

	if _, err := uuid.Parse(input.UUID); input.UUID != "" && err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"errors": []string{err.Error()}})
		return
	}

	hillchart := models.Hillchart{Name: input.Name, UUID: input.UUID, Scopes: []models.Scope{}, Frames: []models.Frame{}}

	if hillchart.UUID == "" {
		hillchart.UUID = uuid.NewString()
	}

	if err := models.DB.Create(&hillchart).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"errors": []string{err.Error()}})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": hillchart})
}
