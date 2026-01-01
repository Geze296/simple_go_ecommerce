package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/geze296/simple_go_ecommerce/internal/models"
	"github.com/geze296/simple_go_ecommerce/internal/services"
)

type ProductHandler struct{
	productService *services.ProductService
}

func NewProductHandler(productService *services.ProductService)  *ProductHandler{
	return &ProductHandler{productService: productService}
}


func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request){
	var product models.Product

	json.NewDecoder(r.Body).Decode(&product)
	err := h.productService.CreateProduct(&product)

	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"Error": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"Message":"Product Created Successfuly",
	})
}