package repository

import (
	"context"
	"errors"

	"github.com/Vadim-12/tszh-backend/pkg/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrganizationPostgres struct {
	db *gorm.DB
}

func NewOrganizationPostgres(db *gorm.DB) *OrganizationPostgres {
	return &OrganizationPostgres{db: db}
}

func (r *OrganizationPostgres) GetAllOrganizations(ctx context.Context) ([]entity.Organization, error) {
	var organizations []entity.Organization
	if err := r.db.WithContext(ctx).Find(&organizations).Error; err != nil {
		return nil, err
	}
	return organizations, nil
}

func (r *OrganizationPostgres) GetByID(ctx context.Context, id uuid.UUID) (*entity.Organization, error) {
	var organization entity.Organization

	err := r.db.WithContext(ctx).First(&organization, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &organization, nil
}

func (r *OrganizationPostgres) GetByINN(ctx context.Context, INN string) (*entity.Organization, error) {
	var organization entity.Organization

	err := r.db.WithContext(ctx).First(&organization, "inn = ?", INN).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &organization, nil
}

func (r *OrganizationPostgres) Create(ctx context.Context, org *entity.Organization, ownerUserId uuid.UUID) (*entity.Organization, error) {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(org).Error; err != nil {
			return err
		}
		organizationStaffLink := &entity.OrganizationStaff{
			OrganizationID: org.ID,
			UserID:         ownerUserId,
			Role:           "owner",
		}
		if err := tx.Create(organizationStaffLink).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return org, nil
}
