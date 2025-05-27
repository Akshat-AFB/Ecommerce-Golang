package main

import (
	"backend-go/routes"
	"fmt"
	"net/http"
	"backend-go/database"
	// "github.com/gin-gonic/gin"
	// "github.com/Akshat-AFB/ecommerce-backend/routes"
)

func main(){
	database.ConnectDatabase()
	routes.RegisterRoutes()
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, E - Commerce!")
	// })
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
	// r:= gin.Default()
	// routes.setUpRoutes(r)
	// r.Run(":8080")
}