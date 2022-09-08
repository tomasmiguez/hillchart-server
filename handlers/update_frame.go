package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tomasmiguez/hillchart-server/models"
)

type UpdateFrameInput struct {
}

func UpdateFrame(c *gin.Context) {
	id := c.Param("id")

	var frame models.Frame
	if err := models.DB.Preload("Scopes").Take(&frame, id).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": []string{"Record not found."}})
		return
	}

	var input UpdateFrameInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"errors": []string{err.Error()}})
		return
	}

	if err := models.DB.Model(&frame).Updates(models.Frame{}).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"errors": []string{err.Error()}})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": frame})
}
