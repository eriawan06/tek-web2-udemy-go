package main

import (
	"github.com/eriawan06/tek-web2-udemy-go/src/cores/database"
	"github.com/eriawan06/tek-web2-udemy-go/src/middlewares"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/auth"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/category"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/user"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	// Setup Database Connection
	db := database.SetupDatabase()
	database.MigrateDb(db)

	// initialize modules/apps
	auth.New(db).InitModule()
	user.New(db).InitModule()
	category.New(db).InitModule()

	// Get Gin Mode from ENV
	mode := os.Getenv("GIN_MODE")

	// Set Gin Mode
	gin.SetMode(mode)

	// Create New App Instance
	app := gin.Default()

	// Setup CORS
	// app.Use(cors.Default())
	app.Use(middlewares.CORSMiddleware())
	//app.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"https://foo.com"},
	//	AllowMethods:     []string{"PUT", "POST", "GET"},
	//	AllowHeaders:     []string{"Origin"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	AllowOriginFunc: func(origin string) bool {
	//		return origin == "https://github.com"
	//	},
	//	MaxAge: 12 * time.Hour,
	//}))

	// Setup Routes
	SetupRoutes(app)

	// Run App at 3000
	err := app.Run(":3000")
	if err != nil {
		return
	}
}
