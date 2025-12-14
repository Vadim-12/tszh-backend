package service

import (
	"context"

	"github.com/Vadim-12/tszh-backend/pkg/repository"
)

type HealthService struct {
	healthRepo repository.Health
}

func NewHealthService(healthRepo repository.Health) *HealthService {
	return &HealthService{
		healthRepo: healthRepo,
	}
}

func (s *HealthService) Ping(ctx context.Context) error {
	return s.healthRepo.Ping(ctx)
}
