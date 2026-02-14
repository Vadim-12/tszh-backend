package service

import (
	"context"

	"github.com/Vadim-12/tszh-backend/pkg/adapters"
	appErrors "github.com/Vadim-12/tszh-backend/pkg/errors"
	"github.com/Vadim-12/tszh-backend/pkg/handler/dto"
	"github.com/Vadim-12/tszh-backend/pkg/repository"
	"github.com/google/uuid"
)

type OrganizationService struct {
	organizationRepo repository.Organization
}

func NewOrganizationService(organizationRepo repository.Organization) *OrganizationService {
	return &OrganizationService{organizationRepo: organizationRepo}
}

func (s *OrganizationService) GetAll(ctx context.Context) ([]dto.OrganizationDto, error) {
	organizations, err := s.organizationRepo.GetAllOrganizations(ctx)
	if err != nil {
		return nil, err
	}

	clearOrganizations := make([]dto.OrganizationDto, 0, len(organizations))
	for _, org := range organizations {
		clearOrg := adapters.FromEntityToOrganizationDto(&org)
		clearOrganizations = append(clearOrganizations, *clearOrg)
	}
	return clearOrganizations, nil
}

func (s *OrganizationService) GetByID(ctx context.Context, id uuid.UUID) (*dto.OrganizationDto, error) {
	org, err := s.organizationRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if org == nil {
		return nil, nil
	}
	result := adapters.FromEntityToOrganizationDto(org)
	return result, nil
}

func (s *OrganizationService) Create(ctx context.Context, dto *dto.CreateOrganizationPayload, userID uuid.UUID) (*dto.OrganizationDto, error) {
	org, err := s.organizationRepo.GetByINN(ctx, dto.INN)
	if err != nil {
		return nil, err
	}
	if org != nil {
		return nil, appErrors.ErrINNAlreadyExists
	}

	creationEntity := adapters.FromCreationDtoToEntity(dto)
	org, err = s.organizationRepo.Create(ctx, creationEntity, userID)
	if err != nil {
		return nil, err
	}
	return adapters.FromEntityToOrganizationDto(org), nil
}
