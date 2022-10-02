package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tomasmiguez/hillchart-server/models"
)

type UpdateHillchartInput struct {
	Name string `json:"name"`
}

func UpdateHillchart(c *gin.Context) {
	id := c.Param("id")

	var hillchart models.Hillchart
	if err := models.DB.Preload("Scopes").Preload("Frames").Preload("Frames.FrameScopes").Find(&hillchart, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": []string{"Record not found."}})
		return
	}

	var input UpdateHillchartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"errors": []string{err.Error()}})
		return
	}

	if err := models.DB.Model(&hillchart).Updates(models.Hillchart{Name: input.Name}).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"errors": []string{err.Error()}})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": hillchart})
}
