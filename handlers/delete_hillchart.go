package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tomasmiguez/hillchart-server/models"
)

func DeleteHillchart(c *gin.Context) {
    id := c.Param("id")

    var hillchart models.Hillchart
    if err := models.DB.Take(&hillchart, id).Error; err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": []string{"Record not found."}})
        return
    }

    models.DB.Delete(&models.Hillchart{}, id)

    c.IndentedJSON(http.StatusOK, gin.H{"data": true})
}

