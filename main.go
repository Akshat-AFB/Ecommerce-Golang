package main

import (
	"backend-go/database"
	"backend-go/redis"
	"backend-go/routes"
	"fmt"
	"net/http"
)

func main() {
	database.ConnectDatabase()
	redis.InitRedis()
	routes.RegisterRoutes()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, E - Commerce!")
	})
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
