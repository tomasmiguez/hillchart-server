package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tomasmiguez/hillchart-server/models"
)

func GetFrames(c *gin.Context) {
	var frames []models.Frame
	if err := models.DB.Preload("Scopes").Find(&frames).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": []string{err.Error()}})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": frames})
}
