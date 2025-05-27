package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"

    "backend-go/database"
    "backend-go/middleware"
    "backend-go/models"
)

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
    userID := middleware.GetUserIDFromContext(r.Context())
    if userID == 0 {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    var cartItems []models.CartItem
    if err := database.DB.Preload("Product").Where("user_id = ?", userID).Find(&cartItems).Error; err != nil || len(cartItems) == 0 {
        http.Error(w, "Cart is empty or failed to retrieve", http.StatusBadRequest)
        return
    }

    var total float64
    var orderItems []models.OrderItem
    for _, item := range cartItems {
        total += float64(item.Quantity) * item.Product.Price
        orderItems = append(orderItems, models.OrderItem{
            ProductID: item.ProductID,
            Quantity:  item.Quantity,
            Price:     item.Product.Price,
        })
    }

    order := models.Order{
        UserID:     userID,
        Total:      total,
        OrderItems: orderItems,
    }

    if err := database.DB.Create(&order).Error; err != nil {
        http.Error(w, "Failed to place order", http.StatusInternalServerError)
        return
    }

    database.DB.Where("user_id = ?", userID).Delete(&models.CartItem{})

    // Reload full order with nested relations
    var fullOrder models.Order
    if err := database.DB.
        Preload("User").
        Preload("OrderItems").
        Preload("OrderItems.Product").
        First(&fullOrder, order.ID).Error; err != nil {
        http.Error(w, "Failed to load order", http.StatusInternalServerError)
        return
    }

    // Remove password before returning
    fullOrder.User.Password = ""

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(fullOrder)
}

func ViewOrders(w http.ResponseWriter, r *http.Request) {
    userID := middleware.GetUserIDFromContext(r.Context())
    if userID == 0 {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    var orders []models.Order
    if err := database.DB.
        Preload("User").
        Preload("OrderItems").
        Preload("OrderItems.Product").
        Where("user_id = ?", userID).
        Find(&orders).Error; err != nil {
        http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
        return
    }

    // Remove passwords from each user
    for i := range orders {
        orders[i].User.Password = ""
    }

    json.NewEncoder(w).Encode(orders)
}

func CancelOrder(w http.ResponseWriter, r *http.Request) {
    userID := middleware.GetUserIDFromContext(r.Context())
    if userID == 0 {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    orderIDStr := strings.TrimPrefix(r.URL.Path, "/orders/cancel/")
    orderID, err := strconv.Atoi(orderIDStr)
    if err != nil {
        http.Error(w, "Invalid order ID", http.StatusBadRequest)
        return
    }

    var order models.Order
    if err := database.DB.First(&order, orderID).Error; err != nil || order.UserID != userID {
        http.Error(w, "Order not found or unauthorized", http.StatusForbidden)
        return
    }

    order.Status = "cancelled"
    database.DB.Save(&order)

    // Reload full order with related data
    var fullOrder models.Order
    if err := database.DB.
        Preload("User").
        Preload("OrderItems").
        Preload("OrderItems.Product").
        First(&fullOrder, orderID).Error; err != nil {
        http.Error(w, "Failed to reload order", http.StatusInternalServerError)
        return
    }

    fullOrder.User.Password = ""

    json.NewEncoder(w).Encode(fullOrder)
}

// package handlers

// import (
//     "encoding/json"
//     "net/http"
//     "backend-go/models"
//     "backend-go/database"
// 	"backend-go/middleware"
// 	// "backend-go/types"
//     "strings"
//     // "github.com/golang-jwt/jwt/v5"
// 	"strconv"
	
// )

// func PlaceOrder(w http.ResponseWriter, r *http.Request) {
//     userID := middleware.GetUserIDFromContext(r.Context())
//     if userID == 0 {
//         http.Error(w, "Unauthorized", http.StatusUnauthorized)
//         return
//     }

//     var cartItems []models.CartItem
//     if err := database.DB.Preload("Product").Where("user_id = ?", userID).Find(&cartItems).Error; err != nil || len(cartItems) == 0 {
//         http.Error(w, "Cart is empty or failed to retrieve", http.StatusBadRequest)
//         return
//     }

//     var total float64
//     var orderItems []models.OrderItem
//     for _, item := range cartItems {
//         total += float64(item.Quantity) * item.Product.Price
//         orderItems = append(orderItems, models.OrderItem{
//             ProductID: item.ProductID,
//             Quantity:  item.Quantity,
//             Price:     item.Product.Price,
//         })
//     }

//     order := models.Order{
//         UserID: userID,
//         Total: total,
//         OrderItems: orderItems,
//     }

//     if err := database.DB.Create(&order).Error; err != nil {
//         http.Error(w, "Failed to place order", http.StatusInternalServerError)
//         return
//     }

//     database.DB.Where("user_id = ?", userID).Delete(&models.CartItem{})

// 	// Reload full order with nested relations
// 	var fullOrder models.Order
// 	if err := database.DB.
// 		Preload("User").
// 		Preload("OrderItems").
// 		Preload("OrderItems.Product").
// 		First(&fullOrder, order.ID).Error; err != nil {
// 		http.Error(w, "Failed to load order", http.StatusInternalServerError)
// 		return
// 	}

//     w.WriteHeader(http.StatusCreated)
//     json.NewEncoder(w).Encode(fullOrder)
// }

// func ViewOrders(w http.ResponseWriter, r *http.Request) {
//     userID := middleware.GetUserIDFromContext(r.Context())
//     if userID == 0 {
//         http.Error(w, "Unauthorized", http.StatusUnauthorized)
//         return
//     }

//     var orders []models.Order
//     if err := database.DB.Preload("OrderItems").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
//         http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
//         return
//     }

//     json.NewEncoder(w).Encode(orders)
// }

// func CancelOrder(w http.ResponseWriter, r *http.Request) {
//     userID := middleware.GetUserIDFromContext(r.Context())
//     if userID == 0 {
//         http.Error(w, "Unauthorized", http.StatusUnauthorized)
//         return
//     }

//     orderIDStr := strings.TrimPrefix(r.URL.Path, "/orders/cancel/")
//     orderID, err := strconv.Atoi(orderIDStr)
//     if err != nil {
//         http.Error(w, "Invalid order ID", http.StatusBadRequest)
//         return
//     }

//     var order models.Order
//     if err := database.DB.First(&order, orderID).Error; err != nil || order.UserID != userID {
//         http.Error(w, "Order not found or unauthorized", http.StatusForbidden)
//         return
//     }

//     order.Status = "cancelled"
//     database.DB.Save(&order)

//     json.NewEncoder(w).Encode(map[string]string{"message": "Order cancelled"})
// }
