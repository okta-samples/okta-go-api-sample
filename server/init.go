package server

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/joho/godotenv"
)

func Init() {

	godotenv.Load("./.okta.env")

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080" // default when missing
	}

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Apply the middleware to the router (works with groups too)
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	// protect all routes
	// router.Use(AuthMiddleware())

	// setup public routes
	router.GET("/", IndexHandler)

	api := router.Group("/api")
	api.GET("/hello", HelloHandler)

	// setup private routes
	authorized := router.Group("/api", AuthMiddleware())
	authorized.GET("/whoami", WhoAmIHandler)

	// Start and run the server
	log.Printf("Running on http://localhost:" + port)
	router.Run(":" + port)
}
