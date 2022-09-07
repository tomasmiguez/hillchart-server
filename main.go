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
	router.DELETE("/hillcharts/:id", handlers.DeleteHillchart)

	router.POST("/frames", handlers.CreateFrame)
  router.GET("/frames", handlers.GetFrames)
	router.GET("/frames/:id", handlers.GetFrame)
	router.PATCH("/frames/:id", handlers.UpdateFrame)
	router.DELETE("/frames/:id", handlers.DeleteFrame)

	router.Run("localhost:8080")
}
