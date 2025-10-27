package users

import (
	"context"

	"gorm.io/gorm"
)

type Service struct{ DB *gorm.DB }

func (service *Service) Create(ctx context.Context, in User) (User, error) {
	if err := service.DB.WithContext(ctx).Create(&in).Error; err != nil {
		return User{}, err
	}
	return in, nil
}

func (service *Service) Get(ctx context.Context, userId string) (*User, error) {
	var user User
	if err := service.DB.WithContext(ctx).First(&user, "id = ?", userId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (service *Service) List(ctx context.Context, limit, offset int) ([]User, error) {
	if limit <= 0 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}
	var result []User
	if err := service.DB.WithContext(ctx).Order("created_at desc").Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (service *Service) Update(ctx context.Context, userId string, updateDto User) (*User, error) {
	result := service.DB.WithContext(ctx).Model(&User{}).Where("id = ?", userId).Updates(map[string]any{
		"full_name":    updateDto.FullName,
		"email":        updateDto.Email,
		"phone_number": updateDto.PhoneNumber,
	})
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, ErrNotFound
	}
	return service.Get(ctx, userId)
}

func (service *Service) Delete(ctx context.Context, userId string) error {
	result := service.DB.WithContext(ctx).Delete(&User{}, "id = ?", userId)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}
