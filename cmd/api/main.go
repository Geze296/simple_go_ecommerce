package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/geze296/simple_go_ecommerce/internal/config"
	"github.com/geze296/simple_go_ecommerce/internal/database"
	"github.com/geze296/simple_go_ecommerce/internal/routes"
)

func main() {
	fmt.Println("Hello, Server")

	cfg := config.LoadConfig()

	_, err := database.ConnectToMongo(cfg)
	if err !=nil {
		log.Fatal("Error with database connection")
	}

	fmt.Println("Mongo is connected")

	mux := routes.SetupRoutes()
	http.ListenAndServe(":8080", mux)
}
