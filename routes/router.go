package routes

import(
	"net/http"
	// "backend-go/handlers"
	"backend-go/middleware"
	"backend-go/controllers"
)

func RegisterRoutes() {
	http.HandleFunc("/api/v1/auth/register", middleware.Method("POST", controllers.Register))
	http.HandleFunc("/api/v1/auth/login", middleware.Method("POST", controllers.Login))

	
	http.HandleFunc("/api/v1/products", middleware.Method("GET", controllers.GetProducts))
	http.HandleFunc("/api/v1/products/create", middleware.Method("POST", middleware.AdminMiddleware(controllers.CreateProduct)))


	http.HandleFunc("/api/v1/cart", middleware.Method("GET", middleware.AuthMiddleware(controllers.ViewCart)))
	http.HandleFunc("/api/v1/cart/add", middleware.Method("POST", middleware.AuthMiddleware(controllers.AddToCart)))
	http.HandleFunc("/api/v1/cart/remove/", middleware.Method("DELETE", middleware.AuthMiddleware(controllers.RemoveFromCart)))
	http.HandleFunc("/api/v1/cart/change/", middleware.Method("POST", middleware.AuthMiddleware(controllers.ChangeQuantity)))


	http.HandleFunc("/api/v1/orders/place", middleware.Method("POST", middleware.AuthMiddleware(controllers.PlaceOrder)))
	http.HandleFunc("/api/v1/orders", middleware.Method("GET", middleware.AuthMiddleware(controllers.ViewOrders)))
	http.HandleFunc("/api/v1/orders/cancel/", middleware.Method("POST", middleware.AuthMiddleware(controllers.CancelOrder)))
}