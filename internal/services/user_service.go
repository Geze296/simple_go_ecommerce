package services

import (

	"github.com/geze296/simple_go_ecommerce/internal/repositories"
)


type UserService struct{
	Repo *repositories.UserRepo
}

func NewUserSevice(r *repositories.UserRepo) *UserService{
	return &UserService{Repo: r}
}