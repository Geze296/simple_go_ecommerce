package routes

import (
	"database/sql"
	"net/http"

	"github.com/geze296/simple_go_ecommerce/internal/handlers"
	"github.com/geze296/simple_go_ecommerce/internal/repositories"
	"github.com/geze296/simple_go_ecommerce/internal/services"
)

func SetupRoutes(db *sql.DB) *http.ServeMux{
	mux := http.NewServeMux()

	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserSevice(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	mux.HandleFunc("GET /health",func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Welcome to Simple E-commerce API"))
	})

	mux.HandleFunc("POST /user", userHandler.Register)
	mux.HandleFunc("/user",userHandler.GetUsers)
	mux.HandleFunc("/user/{id}",userHandler.GetSingleUser)
	// mux.HandleFunc("/order", handlers.)
	return  mux
}