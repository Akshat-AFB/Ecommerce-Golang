package models

type CartItem struct {
    ID        uint    `gorm:"primaryKey"`
    UserID    uint    `gorm:"not null"` // User must be set
    ProductID uint    `gorm:"not null"` // Product must be set
    Quantity  int     `gorm:"not null;check:quantity > 0"` // Ensure quantity is always > 0
    Product   Product `gorm:"foreignKey:ProductID"` // Auto preload if needed
}