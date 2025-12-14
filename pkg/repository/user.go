package repository

import (
	"context"
	"errors"

	"github.com/Vadim-12/tszh-backend/pkg/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(ctx context.Context, u *entity.User) (*entity.User, error) {
	if err := r.db.WithContext(ctx).Create(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserPostgres) FindByID(ctx context.Context, userId uuid.UUID) (*entity.User, error) {
	var user entity.User

	if err := r.db.WithContext(ctx).
		First(&user, "id = ?", userId).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserPostgres) FindByPhoneNumber(ctx context.Context, phoneNumber string) (*entity.User, error) {
	var user entity.User

	if err := r.db.WithContext(ctx).
		Where("phone_number = ?", phoneNumber).
		First(&user).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserPostgres) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User

	if err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
