package main

import (
	"github.com/tomasmiguez/hillchart-server/handlers"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"*"},
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
	}))

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

	router.POST("/scopes", handlers.CreateScope)
	router.GET("/scopes", handlers.GetScopes)
	router.GET("/scopes/:id", handlers.GetScope)
	router.PATCH("/scopes/:id", handlers.UpdateScope)
	router.DELETE("/scopes/:id", handlers.DeleteScope)

	router.POST("/frame-scopes", handlers.CreateFrameScope)
	router.GET("/frame-scopes", handlers.GetFrameScopes)
	router.GET("/frame-scopes/:id", handlers.GetFrameScope)
	router.PATCH("/frame-scopes/:id", handlers.UpdateFrameScope)
	router.DELETE("/frame-scopes/:id", handlers.DeleteFrameScope)

	router.Run("localhost:3000")
}
