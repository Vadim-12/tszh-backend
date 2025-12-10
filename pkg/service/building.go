package service

import "github.com/Vadim-12/tszh-backend/pkg/repository"

type BuildingService struct {
	buildingRepo repository.Building
}

func NewBuildingService(buildingRepo repository.Building) *BuildingService {
	return &BuildingService{buildingRepo: buildingRepo}
}
