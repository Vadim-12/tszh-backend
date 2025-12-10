package service

import "github.com/Vadim-12/tszh-backend/pkg/repository"

type UserService struct {
	userRepo repository.User
}

func NewUserService(userRepo repository.User) *UserService {
	return &UserService{userRepo: userRepo}
}
