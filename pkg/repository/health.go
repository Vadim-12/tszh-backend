package repository

import (
	"context"

	"gorm.io/gorm"
)

type HealthPostgres struct {
	db *gorm.DB
}

func NewHealthPostgres(db *gorm.DB) *HealthPostgres {
	return &HealthPostgres{db: db}
}

func (r *HealthPostgres) Ping(ctx context.Context) error {
	return r.db.WithContext(ctx).Raw("SELECT 1").Error
}
