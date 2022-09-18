package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tomasmiguez/hillchart-server/models"
)

func GetHillchart(c *gin.Context) {
	id := c.Param("id")

	hillchart := models.Hillchart{}
	if err := models.DB.Preload("Scopes").Preload("Frames").Preload("Frames.FrameScopes").Find(&hillchart, id).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": []string{"Record not found."}})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": hillchart})
}
