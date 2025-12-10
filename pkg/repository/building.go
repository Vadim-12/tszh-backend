package repository

import "gorm.io/gorm"

type BuildingPostgres struct {
	db *gorm.DB
}

func NewBuildingPostgres(db *gorm.DB) *BuildingPostgres {
	return &BuildingPostgres{db: db}
}
