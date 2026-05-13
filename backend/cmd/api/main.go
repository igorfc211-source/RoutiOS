package main

import (
	"project-api/internal/database"
	"project-api/internal/modules/user"
	"project-api/internal/shared/config"
	

	"github.com/gin-gonic/gin"
)

func main() {




	config.LoadEnv()

	db := database.Connect()

	db.AutoMigrate(&user.User{})


	r := gin.Default()



	repo := user.NewRepository(db)
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	user.RegisterRoutes(r, handler)

	r.Run(":8080")
}