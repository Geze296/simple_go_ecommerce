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
	
	productRepo := repositories.NewProductRepo(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	mux.HandleFunc("GET /health",func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Welcome to Simple E-commerce API"))
	})

	mux.HandleFunc("POST /user", userHandler.Register)
	mux.HandleFunc("/user",userHandler.GetUsers)
	mux.HandleFunc("/user/{id}",userHandler.GetSingleUser)
	mux.HandleFunc("PATCH /user/{id}",userHandler.EditUser)
	mux.HandleFunc("POST /product", productHandler.CreateProduct)
	// mux.HandleFunc("/order", handlers.)
	return  mux
}