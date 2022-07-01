package main

import (
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/category"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/course"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/general/upload"
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

	v1 := app.Group("/api/v1")
	{
		public := v1.Group("")
		{
			auth.NewRouter(public.Group("/auth"))
			category.NewPublicRouter(public.Group("/categories"))
			course.NewPublicRouter(public.Group("/courses"))
		}

		private := v1.Group("")
		{
			private.Use(middlewares.JwtAuthMiddleware())

			category.NewRouter(private.Group("/categories"))
			course.NewRouter(private.Group("/courses"))
			upload.NewRouter(private.Group("/upload"))
		}
	}
}
