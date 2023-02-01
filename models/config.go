package models

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-auth/go-auth-api/env"
)

var DB *gorm.DB

func ConnectDB() {

	psqlInfo := env.GetDbEnv()

	Db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		fmt.Println("Database connection error")
		log.Fatalf("Error Message: %s", err)
	} else {
		fmt.Println("********************************")
		fmt.Println("Database connected succesfully")
		fmt.Println("********************************")
	}
	Db.AutoMigrate(
		&User{},
	)
	DB = Db
}