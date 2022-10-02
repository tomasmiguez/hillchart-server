package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tomasmiguez/hillchart-server/models"
)

func DeleteHillchart(c *gin.Context) {
	id := c.Param("id")

	var hillchart models.Hillchart
	if err := models.DB.Take(&hillchart, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": []string{"Record not found."}})
		return
	}

	if err := models.DB.Delete(&models.Hillchart{}, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"errors": []string{err.Error()}})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": true})
}
