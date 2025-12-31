package handlers

import (
	"fmt"
	"net/http"

	"github.com/geze296/simple_go_ecommerce/internal/services"
)

type UserHandler struct{
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler{
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Get All users Using Postgresql")
}

func Register(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Registration Handler")
}