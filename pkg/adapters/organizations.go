package adapters

import (
	"github.com/Vadim-12/tszh-backend/pkg/entity"
	"github.com/Vadim-12/tszh-backend/pkg/handler/dto"
)

func FromEntityToOrganizationDto(entity *entity.Organization) *dto.OrganizationDto {
	return &dto.OrganizationDto{
		ID:           entity.ID,
		FullTitle:    entity.FullTitle,
		ShortTitle:   entity.ShortTitle,
		INN:          entity.INN,
		Email:        *entity.Email,
		LegalAddress: entity.LegalAddress,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
	}
}

func FromCreationDtoToEntity(dto *dto.CreateOrganizationPayload) *entity.Organization {
	return &entity.Organization{
		FullTitle:    dto.FullTitle,
		ShortTitle:   dto.ShortTitle,
		INN:          dto.INN,
		Email:        dto.Email,
		LegalAddress: dto.LegalAddress,
		StaffLinks:   []*entity.OrganizationStaff{},
	}
}
