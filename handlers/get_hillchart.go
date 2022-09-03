package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tomasmiguez/hillchart-server/models"
)

func GetHillchart(c *gin.Context) {
  	id := c.Param("id")

    hillchart := models.Hillchart{}
    models.DB.Take(&hillchart, id)

    c.IndentedJSON(http.StatusOK, gin.H{"data": hillchart})
}

