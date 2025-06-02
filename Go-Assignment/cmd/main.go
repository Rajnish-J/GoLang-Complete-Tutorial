package main

import (
    "go-assignment/config"
    "go-assignment/db"
    "go-assignment/handlers"
    "go-assignment/repository"
    "go-assignment/routes"
    "go-assignment/services"

    "github.com/gin-gonic/gin"
)

func main() {
    cfg := config.LoadConfig()
    db.InitDB(cfg)

    userRepo := repository.NewUserRepository(db.DB)
    userService := services.NewUserService(userRepo)
    userHandler := handlers.NewUserHandler(userService)

    router := gin.Default()
    routes.RegisterRoutes(router, userHandler)

    router.Run(":8080")
}
