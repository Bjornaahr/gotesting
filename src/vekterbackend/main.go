package main

import (
	"VekterBackend/src/api"
	"VekterBackend/src/initializers"
	"VekterBackend/src/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.Connect()
	initializers.Migrate()
}

func main() {
	r := gin.Default()
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
