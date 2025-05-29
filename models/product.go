package models

type Product struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"not null;default:'Unknown Product'"`    
	Price       float64 `gorm:"not null;default:0.0"`                 
	Description string  `gorm:"not null;default:'No description'"`      
	ImageURL    string  `gorm:"not null;default:'https://example.com/default.jpg'"` 
	Quantity    int     `gorm:"not null;default:0"`                    
}