package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tomasmiguez/hillchart-server/models"
)

type CreateHillchartInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateHillchart(c *gin.Context) {
	var input CreateHillchartInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"errors": []string{err.Error()}})
		return
	}

	hillchart := models.Hillchart{Name: input.Name}
	models.DB.Create(&hillchart)

	c.IndentedJSON(http.StatusOK, gin.H{"data": hillchart})
}
