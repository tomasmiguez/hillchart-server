package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tomasmiguez/hillchart-server/models"
)

func GetHillcharts(c *gin.Context) {
    var hillcharts []models.Hillchart
    models.DB.Find(&hillcharts)

    c.IndentedJSON(http.StatusOK, gin.H{"data": hillcharts})
}

