package database

import (
	"database/sql"
	"log"

	// "os"
	"fmt"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// func getEnv(key, fallback string) string {
// 	val := os.Getenv(key)
//     log.Printf("Environment variable %s: %s", key, val)
// 	if val == "" {
// 		return fallback
// 	}
// 	return val
// }

func ConnectDatabase() {
	// Read environment variables or use defaults
	// host := getEnv("DB_HOST", "localhost")
	// port := getEnv("DB_PORT", "5432")
	// user := getEnv("DB_USER", "postgres")
	// password := getEnv("DB_PASSWORD", "postgres")
	// dbname := getEnv("DB_NAME", "ecommerce")
	host := "localhost"
	port := "5432"
	user := "akshat"
	password := "password"
	dbname := "ecommerce-golang"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

}

func GetDB() *sql.DB {
	return DB
}
