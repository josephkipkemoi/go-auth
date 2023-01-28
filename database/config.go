package database

import (
	"log"
	"os"

	"go-auth/go-auth-api/env"
	_ "go-auth/go-auth-api/env" // load env variables

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbName string = os.Getenv("DB_NAME")

func Connect() *gorm.DB {
	dsn := env.GetDbEnv()
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database Connection Error: %s", err)
	}

	return db
}