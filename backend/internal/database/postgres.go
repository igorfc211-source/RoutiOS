package database

import (
	"log"

	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"project-api/internal/shared/config"
)

func Connect() *gorm.DB{
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.GetEnv("DB_HOST"),
		config.GetEnv("DB_USER"),
		config.GetEnv("DB_PASSWORD"),
		config.GetEnv("DB_NAME"),
		config.GetEnv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})


	if err != nil {

		log.Fatal(err)
	}
	fmt.Println("connected to " + dsn)

	return db
}