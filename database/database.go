package database

import(
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"backend-go/models"
)

var DB *gorm.DB
func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("ecommerce.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate Product table, User table
    db.AutoMigrate(
		&models.Product{},
		&models.User{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
	)


    DB = db
}
func GetDB() *gorm.DB {
    return DB
}