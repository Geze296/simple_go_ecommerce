package services

import (

	"github.com/geze296/simple_go_ecommerce/internal/repositories"
)


type UserService struct{
	Repo *repositories.UserRepository
}

func NewUserSevice(repo *repositories.UserRepository) *UserService{
	return &UserService{Repo: repo}
}