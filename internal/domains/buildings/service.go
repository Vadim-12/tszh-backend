package buildings

import (
	"context"

	"gorm.io/gorm"
)

type Service struct{ DB *gorm.DB }

func (service *Service) Create(ctx context.Context, creationDto BuildingModel) (*BuildingModel, error) {
	if err := service.DB.WithContext(ctx).Create(&creationDto).Error; err != nil {
		return nil, err
	}
	return &creationDto, nil
}

func (service *Service) GetOne(ctx context.Context, buildingId string) (*BuildingModel, error) {
	var building BuildingModel
	if err := service.DB.WithContext(ctx).First(&building, "id = ?", buildingId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &building, nil
}

func (service *Service) GetList(ctx context.Context, limit, offset int) ([]BuildingModel, error) {
	if limit <= 0 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}
	var result []BuildingModel
	if err := service.DB.WithContext(ctx).Order("created_at desc").Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (service *Service) UpdateOne(ctx context.Context, buildingId string, patch BuildingModel) (*BuildingModel, error) {
	result := service.DB.WithContext(ctx).Model(&BuildingModel{}).
		Where("id = ?", buildingId).
		Updates(map[string]any{
			"number":     patch.Number,
			"floor":      patch.Floor,
			"owner_name": patch.OwnerName,
		})
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, ErrNotFound
	}
	return service.GetOne(ctx, buildingId)
}

func (service *Service) DeleteOne(ctx context.Context, buildingId string) error {
	result := service.DB.WithContext(ctx).Delete(&BuildingModel{}, "id = ?", buildingId)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

// На будущее — пример транзакции:
/*
func (s *Service) DoInTx(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return s.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}
*/
