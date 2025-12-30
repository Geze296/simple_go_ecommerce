package routes

import (
	"net/http"
	"github.com/geze296/simple_go_ecommerce/internal/handlers"
)

func SetupRoutes() *http.ServeMux{
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health",func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Welcome to Simple E-commerce API"))
	})

	mux.HandleFunc("POST /register", handlers.Register)
	mux.HandleFunc("/user",handlers.UserHandler)
	// mux.HandleFunc("/order", handlers.)
	return  mux
}