package main

import (
	"backend-go/database"
	"backend-go/redis"
	"backend-go/routes"
	"fmt"
	"net/http"
	"log"
)

func main() {
	database.ConnectDatabase()
	redis.InitRedis()
	routes.RegisterRoutes()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, E - Commerce!")
	})


	port := ":8080"
	fmt.Println("Starting server on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Printf("Port %s unavailable: %v. Trying fallback port 8081...\n", port, err)
		fallbackPort := ":8081"
		fmt.Println("Trying to run on fallback port", fallbackPort)
		err = http.ListenAndServe(fallbackPort, nil)
		if err != nil {
			log.Fatalf("Failed to start server on fallback port %s: %v", fallbackPort, err)
		}
	}
}
