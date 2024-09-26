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

	//init session store
	core.InitSessionStore()

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
	r.POST("/api/v1/auth", handlers.ProcessLogin)
	r.DELETE("/api/v1/auth", handlers.ProcessLogout)

	protected := r.Group("")
	protected.Use(middleware.AuthRequired)
	{
		protected.GET("/api/v1/item", handlers.GetItems)
		protected.GET("/api/v1/item/:id", handlers.GetItem)
		protected.POST("/api/v1/item/perm", handlers.CreatePermItem)
		protected.POST("/api/v1/item/temp", handlers.CreateTempItem)
		protected.POST("/api/v1/item/count", handlers.CreateCountItem)
		protected.PUT("/api/v1/item/perm/:id", handlers.UpdatePermItem)
		protected.PUT("/api/v1/item/temp/:id", handlers.UpdateTempItem)
		protected.PUT("/api/v1/item/count/:id", handlers.UpdateCountItem)
		protected.PUT("/api/v1/item/:id", handlers.ConvertItem)
		protected.DELETE("/api/v1/item/:id", handlers.DeleteItem)

		protected.PUT("/api/v1/user", handlers.UpdateUser)

		protected.GET("/", handlers.GetMainPage)
	}

	r.GET("/s/:shortUrl", handlers.ToOriginalUrl)
	r.GET("/login", handlers.GetLoginPage)
	err = r.Run(":8080")

	if err != nil {
		panic(err)
	}
}
