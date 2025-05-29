package models

import(
	"time"
)
type Order struct {
    ID        uint        `gorm:"primaryKey"`
    UserID     uint         `gorm:"not null"`    
	User       User         `gorm:"foreignKey:UserID"` 
	Total      float64      `gorm:"not null"`  
    OrderItems []OrderItem  `gorm:"constraint:OnDelete:CASCADE;"`
    Status     string       `gorm:"type:varchar(20);not null;default:'placed'"`
	CreatedAt  *time.Time   `gorm:"autoCreateTime"`
	UpdatedAt  *time.Time   `gorm:"autoUpdateTime"`
}

type OrderItem struct {
    ID        uint    `gorm:"primaryKey"`
    OrderID   uint    `gorm:"not null"`  
	ProductID uint    `gorm:"not null"`                          // Foreign key to Product
    Product   Product `gorm:"foreignKey:ProductID"`              // Association
	Quantity  int     `gorm:"not null;check:quantity > 0"`       // Must be at least 1
	Price     float64 `gorm:"not null;check:price >= 0"`         // Price per unit
}
