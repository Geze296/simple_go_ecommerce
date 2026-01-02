package services

import (
	"github.com/geze296/simple_go_ecommerce/internal/models"
	"github.com/geze296/simple_go_ecommerce/internal/repositories"
)


type OrderService struct{
	orderRepository *repositories.OrderRepository
}

func NewOrderService(orderRepo *repositories.OrderRepository) *OrderService{
	return &OrderService{orderRepository: orderRepo}
}

func (s *OrderService) CreateOrder(order *models.Order) error{
	err := s.orderRepository.Create(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *OrderService) GetAllOrders() ([]models.Order, error){
	orders, err := s.orderRepository.GetAll()

	if err !=nil {
		return nil, err
	}
	return orders, nil
}