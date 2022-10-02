package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tomasmiguez/hillchart-server/models"
)

type CreateFrameScopeInput struct {
	Title    string  `json:"title" binding:"required"`
	Position float32 `json:"position"`
	FrameID  string    `json:"frame_id" binding:"required"`
	ScopeID  string    `json:"scope_id" binding:"required"`
}

func CreateFrameScope(c *gin.Context) {
	var input CreateFrameScopeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"errors": []string{err.Error()}})
		return
	}

	frameScope := models.FrameScope{Title: input.Title, Position: input.Position, FrameID: input.FrameID, ScopeID: input.ScopeID}
	if err := models.DB.Select("Title", "Position", "FrameID", "ScopeID").Create(&frameScope).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"errors": []string{err.Error()}})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": frameScope})
}

func GetFrameScopes(c *gin.Context) {
	var frameScopes []models.FrameScope
	if err := models.DB.Find(&frameScopes).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": []string{err.Error()}})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": frameScopes})
}

func GetFrameScope(c *gin.Context) {
	id := c.Param("id")

	frameScope := models.FrameScope{}
	if err := models.DB.Take(&frameScope, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": []string{"Record not found."}})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": frameScope})
}

type UpdateFrameScopeInput struct {
	Title    string  `json:"title"`
	Position float32 `json:"position"`
}

func UpdateFrameScope(c *gin.Context) {
	id := c.Param("id")

	var frameScope models.FrameScope
	if err := models.DB.Take(&frameScope, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": []string{"Record not found."}})
		return
	}

	var input UpdateFrameScopeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"errors": []string{err.Error()}})
		return
	}

	if err := models.DB.Model(&frameScope).Updates(models.FrameScope{Title: input.Title, Position: input.Position}).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"errors": []string{err.Error()}})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": frameScope})
}

func DeleteFrameScope(c *gin.Context) {
	id := c.Param("id")

	var frameScope models.FrameScope
	if err := models.DB.Take(&frameScope, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errors": []string{"Record not found."}})
		return
	}

	if err := models.DB.Delete(&models.FrameScope{}, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"errors": []string{err.Error()}})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": true})
}
