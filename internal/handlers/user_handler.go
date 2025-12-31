package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/geze296/simple_go_ecommerce/internal/models"
	"github.com/geze296/simple_go_ecommerce/internal/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	
	users, err := h.userService.GetAllUsers()

	w.Header().Set("Content-Type","application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"data":    users,
		"message": "Successfuly return user data",
	})
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {

	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	err := h.userService.RegisterUser(&user)

	w.Header().Set("Content-Type","application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"message": "user created successfuly",
		"data": map[string]string{
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
