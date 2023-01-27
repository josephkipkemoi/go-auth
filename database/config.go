package database

import (
	"database/sql"
	"log"
	"os"

	"go-auth/go-auth-api/env"
	_ "go-auth/go-auth-api/env" // load env variables

	_ "github.com/lib/pq"
)

var dbName string = os.Getenv("DB_NAME")

func Connect() *sql.DB {
	psqlInfo := env.GetDbEnv()

	db, err := sql.Open("postgres",psqlInfo)
	if err != nil {
		log.Fatalf("Database Connection Error: %s", err)
	}

	return db
}

func Insert(body string) (sql.Result){
	res, err := Connect().Exec("SELECT * FROM users")
	if err != nil {
		log.Fatalf("Error: %v",err)
		return res
	}
	// r,err := 
	// if err != nil {
	// 	return 90, err
	// }
	return res
}

func All() {
	// res := Db.Exec("SELECT * FROM " + dbName)
	// return res
}