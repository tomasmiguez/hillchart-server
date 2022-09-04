package main

import (
	"github.com/tomasmiguez/hillchart-server/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/hillcharts", handlers.CreateHillchart)
	router.GET("/hillcharts", handlers.GetHillcharts)
	router.GET("/hillcharts/:id", handlers.GetHillchart)
	router.PATCH("/hillcharts/:id", handlers.UpdateHillchart)
	router.DELETE("hillcharts/:id", handlers.DeleteHillchart)

	router.Run("localhost:8080")
}
