package models

import(
	"time"
)
type Order struct {
    ID        uint        
    UserID     uint        
	User       User         
	Total      float64     
    OrderItems []OrderItem  
    Status     string       
	CreatedAt  *time.Time   
	UpdatedAt  *time.Time   
}

type OrderItem struct {
    ID        uint    
    OrderID   uint    
	ProductID uint    
    Product   Product 
	Quantity  int     
	Price     float64 
}
