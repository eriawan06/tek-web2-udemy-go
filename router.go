package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/eriawan06/tek-web2-udemy-go/src/middlewares"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/auth"
)

// SetupRoutes Setup Routes
func SetupRoutes(app *gin.Engine) {
	// Check Server Status Endpoint
	app.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Server alive!",
			"data":    context,
		})
	})

	// Setup Routes Group
	authGroup := app.Group("/api/v1/auth")
	{
		auth.NewRouter(authGroup)
	}

	v1 := app.Group("/api/v1")
	{
		v1.Use(middlewares.JwtAuthMiddleware())
	}
}
