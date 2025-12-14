package repository

import (
	"context"
	"errors"

	"github.com/Vadim-12/tszh-backend/pkg/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshTokenPostgres struct {
	db *gorm.DB
}

func NewRefreshTokenPostgres(db *gorm.DB) *RefreshTokenPostgres {
	return &RefreshTokenPostgres{db: db}
}

func (r *RefreshTokenPostgres) Save(ctx context.Context, token *entity.RefreshToken) error {
	result := r.db.WithContext(ctx).Create(token)
	return result.Error
}

func (r *RefreshTokenPostgres) GetByID(ctx context.Context, id uuid.UUID) (*entity.RefreshToken, error) {
	var token entity.RefreshToken
	if err := r.db.WithContext(ctx).
		First(&token, "id = ?", id).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // "не найдено" — это не ошибка, это nil
		}
		return nil, err
	}
	return &token, nil
}

func (r *RefreshTokenPostgres) DeleteByID(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).
		Delete(&entity.RefreshToken{}, "id = ?", id).Error
}

func (r *RefreshTokenPostgres) DeleteAllByUserID(ctx context.Context, userID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Delete(&entity.RefreshToken{}, "user_id = ?", userID).Error
}
