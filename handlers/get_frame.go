package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tomasmiguez/hillchart-server/models"
)

func GetFrame(c *gin.Context) {
	id := c.Param("id")

	frame := models.Frame{}
	if err := models.DB.Preload("Scopes").Take(&frame, id).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": []string{"Record not found."}})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": frame})
}
