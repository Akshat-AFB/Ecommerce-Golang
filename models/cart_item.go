package models

type CartItem struct {
    ID        uint    
    UserID    uint    
    ProductID uint    
    Quantity  int     
    Product   Product 
}