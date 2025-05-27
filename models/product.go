package models

type Product struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"not null;default:'Unknown Product'"`     // Default product name
	Price       float64 `gorm:"not null;default:0.0"`                   // Default price is 0.0
	Description string  `gorm:"not null;default:'No description'"`      // Default description
	ImageURL    string  `gorm:"not null;default:'https://example.com/default.jpg'"` // Default image URL
	Quantity    int     `gorm:"not null;default:0"`                     // Default quantity is 0
}