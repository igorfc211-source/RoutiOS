package main

import (
	"project-api/internal/database"
	"project-api/internal/modules/user"
	"project-api/internal/shared/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {




	config.LoadEnv()

	db := database.Connect()

	db.AutoMigrate(&user.User{})


	r := gin.Default()

	r.Use(cors.Default())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	
		AllowCredentials: true,
	}))
	repo := user.NewRepository(db)
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	user.RegisterRoutes(r, handler)

	r.Run(":8080")
}