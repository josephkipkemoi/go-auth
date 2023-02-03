package env

import (
	"fmt"
	"os"
)

// Database ENV
const (
	host = "127.0.0.1"
	port = "5432"
	user = "jkemboe"
	password = "commandme007!@~"
	dbname = "pinacle_db"
)

func setEnv() {
	// APP URL
	os.Setenv("GO_AUTH_API_DEV_URL", "http://localhost:8080/")
	//  Database env
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s	password=%s	dbname=%s sslmode=disable", host, port, user, password, dbname)
	os.Setenv("GOAUTH_DB", psqlInfo)
	os.Setenv("DB_NAME", dbname)
	// Token Env
	// os.Setenv("API_SECRET", "maasai")
}

// Get development app url
func GetDevAppUrl() string {
	return os.Getenv("GO_AUTH_API_DEV_URL")
}
// Get Database env
func GetDbEnv() string {
	return os.Getenv("GOAUTH_DB")
}

func init() {
	setEnv()
}

