package database

import (
    "database/sql"
    "log"
    "os"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    _ "github.com/lib/pq"
)

var DB *sql.DB

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
    log.Printf("Environment variable %s: %s", key, val)
	if val == "" {
		return fallback
	}
	return val
}

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

    createTables()
}

func createTables() {
    queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username TEXT UNIQUE NOT NULL,
			email TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			role TEXT NOT NULL DEFAULT 'user'
		);`,
		`CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL DEFAULT 'Unknown Product',
			price REAL NOT NULL DEFAULT 0.0,
			description TEXT NOT NULL DEFAULT 'No description',
			image_url TEXT NOT NULL DEFAULT 'https://example.com/default.jpg',
			quantity INTEGER NOT NULL DEFAULT 0
		);`,
		`CREATE TABLE IF NOT EXISTS cart_items (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL,
			product_id INTEGER NOT NULL,
			quantity INTEGER NOT NULL CHECK(quantity > 0),
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
			FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
		);`,
		`CREATE TABLE IF NOT EXISTS orders (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL,
			total REAL NOT NULL,
			status TEXT NOT NULL DEFAULT 'placed',
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);`,
		`CREATE TABLE IF NOT EXISTS order_items (
			id SERIAL PRIMARY KEY,
			order_id INTEGER NOT NULL,
			product_id INTEGER NOT NULL,
			quantity INTEGER NOT NULL CHECK(quantity > 0),
			price REAL NOT NULL CHECK(price >= 0),
			FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
			FOREIGN KEY (product_id) REFERENCES products(id)
		);`,
	}

    for _, query := range queries {
        if _, err := DB.Exec(query); err != nil {
            log.Fatalf("Failed to execute table creation query: %v", err)
        }
    }
}

func GetDB() *sql.DB {
    return DB
}
