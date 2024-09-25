package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/hrabit64/shortlink/app/api/handlers"
	"github.com/hrabit64/shortlink/app/api/middleware"
	"github.com/hrabit64/shortlink/app/core"
	"github.com/hrabit64/shortlink/app/utils"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	log.Println("Starting the application...")

	// Load the environment variables
	err := godotenv.Load("./resource/.env")
	if err != nil {
		panic(err)
	}

	// Initialize the database
	log.Println("Initializing the database start...")
	err = core.InitDB("./resource/init.sql")
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err = v.RegisterValidation("url", utils.CheckUrlRegexFunc)

		if err != nil {
			panic(err)
		}

	}

	r.Use(middleware.ErrorHandlingMiddleware)

	// Register the routes
	r.GET("/api/v1/auth", handlers.ProcessLogin)
	r.DELETE("/api/v1/auth", handlers.ProcessLogout)

	protected := r.Group("/api/v1")
	protected.Use(middleware.AuthRequired)
	{
		protected.GET("/item", handlers.GetItems)
		protected.GET("/item/:id", handlers.GetItem)
		protected.POST("/item/perm", handlers.CreatePermItem)
		protected.POST("/item/temp", handlers.CreateTempItem)
		protected.POST("/item/count", handlers.CreateCountItem)
		protected.PUT("/item/perm/:id", handlers.UpdatePermItem)
		protected.PUT("/item/temp/:id", handlers.UpdateTempItem)
		protected.PUT("/item/count/:id", handlers.UpdateCountItem)
		protected.PUT("/item/:id", handlers.ConvertItem)
		protected.DELETE("/item/:id", handlers.DeleteItem)

		protected.PUT("/user", handlers.UpdateUser)
	}

	r.GET("/s/:shortUrl", handlers.ToOriginalUrl)
	r.GET("/login", handlers.GetLoginPage)
	err = r.Run(":8080")

	if err != nil {
		panic(err)
	}
}
