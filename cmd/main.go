package main

import (
	"cms/Internals/handlers"
	"cms/Internals/repository"
	"cms/Internals/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// We will load the env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot find env file")
		return
	}

	r := gin.Default()

	// Handler Object
	repo := repository.NewInMemory()
	jwtService := &services.JWTService{}
	handler := handlers.NewHandler(repo, jwtService)
	v1 := r.Group("/api/v1")
	// Two types of groups // auth routes
	auth := v1.Group("/auth") // /api/v1/auth

	auth.POST("/signup", handler.Signup)
	auth.POST("/login", handler.Login)

	// user routes
	user := v1.Group("/user") // /api/v1/user
	user.GET("/getUsers", handler.GetAllUsers)

	err = r.Run("localhost:8080")
	if err != nil {
		return
	}

}
