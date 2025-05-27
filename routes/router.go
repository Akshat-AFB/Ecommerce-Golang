package routes

import(
	"net/http"
	"backend-go/handlers"
	"backend-go/middleware"
)

func RegisterRoutes() {
	// Auth
	http.HandleFunc("/auth/register", middleware.Method("POST", handlers.Register))
	http.HandleFunc("/auth/login", middleware.Method("POST", handlers.Login))

	// Products
	http.HandleFunc("/products", middleware.Method("GET", handlers.GetProducts))
	http.HandleFunc("/products/create", middleware.Method("POST", middleware.AdminMiddleware(handlers.CreateProduct)))

	// Cart
	http.HandleFunc("/cart", middleware.Method("GET", middleware.AuthMiddleware(handlers.ViewCart)))
	http.HandleFunc("/cart/add", middleware.Method("POST", middleware.AuthMiddleware(handlers.AddToCart)))
	http.HandleFunc("/cart/remove/", middleware.Method("DELETE", middleware.AuthMiddleware(handlers.RemoveFromCart)))
	http.HandleFunc("/cart/change/", middleware.Method("POST", middleware.AuthMiddleware(handlers.ChangeCartQuantity)))

	//Orders
	http.HandleFunc("/orders/place", middleware.Method("POST", middleware.AuthMiddleware(handlers.PlaceOrder)))
	http.HandleFunc("/orders", middleware.Method("GET", middleware.AuthMiddleware(handlers.ViewOrders)))
	http.HandleFunc("/orders/cancel/", middleware.Method("POST", middleware.AuthMiddleware(handlers.CancelOrder)))


}