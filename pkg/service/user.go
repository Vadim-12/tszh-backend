package service

import (
	"context"

	"github.com/Vadim-12/tszh-backend/pkg/adapters"
	"github.com/Vadim-12/tszh-backend/pkg/handler/dto"
	"github.com/Vadim-12/tszh-backend/pkg/repository"
	"github.com/google/uuid"
)

type UserService struct {
	userRepo repository.User
}

func NewUserService(userRepo repository.User) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetMe(ctx context.Context, userID uuid.UUID) (*dto.GetMeResponseDto, error) {
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return adapters.FromUserEntityToDto(user), nil
}
