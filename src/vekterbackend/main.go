package main

import (
	"VekterBackend/src/api"
	"VekterBackend/src/initializers"
	"VekterBackend/src/middleware"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.Connect()
	initializers.Migrate()
}

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}                                       // Adjust this to specify the allowed origins
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"} // Allowed HTTP methods
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}

	// Use the CORS middleware
	r.Use(cors.New(config))

	r.POST("/register", api.CreateUser)
	r.POST("/login", api.Login)

	r.Use(middleware.JWTAuthMiddleware())
	{
		r.GET("/ping", api.Pong)
	}

	err := r.Run()
	if err != nil {
		fmt.Print(err)
	}
}
