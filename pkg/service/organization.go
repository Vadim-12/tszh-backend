package service

import "github.com/Vadim-12/tszh-backend/pkg/repository"

type OrganizationService struct {
	organizationRepo repository.Organization
}

func NewOrganizationService(organizationRepo repository.Organization) *OrganizationService {
	return &OrganizationService{organizationRepo: organizationRepo}
}
