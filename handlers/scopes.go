package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tomasmiguez/hillchart-server/models"
)

type CreateScopeInput struct {
	Color       string `json:"color"`
	HillchartID string   `json:"hillchart_id" binding:"required"`
}

func CreateScope(c *gin.Context) {
	var input CreateScopeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"errors": []string{err.Error()}})
		return
	}

	scope := models.Scope{Color: input.Color, HillchartID: input.HillchartID}
	if err := models.DB.Select("Color", "HillchartID").Create(&scope).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"errors": []string{err.Error()}})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": scope})
}

func GetScopes(c *gin.Context) {
	var scopes []models.Scope
	if err := models.DB.Find(&scopes).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": []string{err.Error()}})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": scopes})
}

func GetScope(c *gin.Context) {
	id := c.Param("id")

	scope := models.Scope{}
	if err := models.DB.Take(&scope, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": []string{"Record not found."}})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": scope})
}

type UpdateScopeInput struct {
	Color string `json:"color"`
}

func UpdateScope(c *gin.Context) {
	id := c.Param("id")

	var scope models.Scope
	if err := models.DB.Take(&scope, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": []string{"Record not found."}})
		return
	}

	var input UpdateScopeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"errors": []string{err.Error()}})
		return
	}

	if err := models.DB.Model(&scope).Updates(models.Scope{Color: input.Color}).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"errors": []string{err.Error()}})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": scope})
}

func DeleteScope(c *gin.Context) {
	id := c.Param("id")

	var scope models.Scope
	if err := models.DB.Take(&scope, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": []string{"Record not found."}})
		return
	}

	if err := models.DB.Delete(&models.Scope{}, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"errors": []string{err.Error()}})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": true})
}
