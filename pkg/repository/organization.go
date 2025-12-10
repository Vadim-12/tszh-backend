package repository

import "gorm.io/gorm"

type OrganizationPostgres struct {
	db *gorm.DB
}

func NewOrganizationPostgres(db *gorm.DB) *OrganizationPostgres {
	return &OrganizationPostgres{db: db}
}
