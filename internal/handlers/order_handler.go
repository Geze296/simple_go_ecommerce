package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/geze296/simple_go_ecommerce/internal/models"
	"github.com/geze296/simple_go_ecommerce/internal/services"
)

type OrderHandler struct{
	orderService services.OrderService
}

func NewOrderHandler(orderService services.OrderService) *OrderHandler{
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order

	json.NewDecoder(r.Body).Decode(&order)
	
	err := h.orderService.CreateOrder(&order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"Error":err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Successfuly Order is created",
	})
}

func (h *OrderHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	orders, err := h.orderService.GetAllOrders()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"Error":err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"message":"Success",
		"data":orders,
	})
}