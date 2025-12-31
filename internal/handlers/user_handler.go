package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/geze296/simple_go_ecommerce/internal/models"
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

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request){

	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	err := h.userService.RegisterUser(&user)
	w.Header().Set("Content-Type","application/json")

	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"message":"user created successfuly",
		"data": map[string]string{
			"name":user.Name,
			"email":user.Email,
		},
	})
}