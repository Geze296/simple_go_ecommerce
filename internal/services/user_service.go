package services

import (
	"errors"
	"github.com/geze296/simple_go_ecommerce/internal/models"
	"github.com/geze296/simple_go_ecommerce/internal/repositories"
)


type UserService struct{
	Repo *repositories.UserRepository
}

func NewUserSevice(repo *repositories.UserRepository) *UserService{
	return &UserService{Repo: repo}
}

func (s *UserService) RegisterUser(user *models.User) error{
	if user.Name == "" {
		return errors.New("Name is required")
	}

	err := s.Repo.Create(user)
	return err
}

func (s *UserService) GetAllUsers() ([]models.User, error){
	users, err := s.Repo.GetAll()
	if err!=nil {
		return nil,err
	}
	return users,nil
}