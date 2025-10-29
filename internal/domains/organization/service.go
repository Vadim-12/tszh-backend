package organizations

import (
	"context"

	"gorm.io/gorm"
)

type Service struct{ DB *gorm.DB }

func (service *Service) Create(ctx context.Context, creationDto OrganizationModel) (*OrganizationModel, error) {
	if err := service.DB.WithContext(ctx).Create(&creationDto).Error; err != nil {
		return nil, err
	}
	return &creationDto, nil
}

func (service *Service) GetOne(ctx context.Context, organizationId string) (*OrganizationModel, error) {
	var organization OrganizationModel
	if err := service.DB.WithContext(ctx).First(&organization).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &organization, nil
}

func (service *Service) GetList(ctx context.Context, limit, offset int) ([]OrganizationModel, error) {
	if limit <= 0 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}
	var result []OrganizationModel
	if err := service.DB.WithContext(ctx).Order("created_at desc").Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (service *Service) UpdateOne(ctx context.Context, organizationId string, patch OrganizationModel) (*OrganizationModel, error) {
	result := service.DB.WithContext(ctx).Model(&OrganizationModel{}).
		Where("id = ?", organizationId).
		Updates(map[string]any{
			"name":          patch.Name,
			"email_address": patch.EmailAddress,
			"post_address":  patch.PostAddress,
			"inn":           patch.INN,
		})
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, ErrNotFound
	}
	return service.GetOne(ctx, organizationId)
}

func (service *Service) DeleteOne(ctx context.Context, organizationId string) error {
	result := service.DB.WithContext(ctx).Delete(&OrganizationModel{}, "id = ?", organizationId)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}
