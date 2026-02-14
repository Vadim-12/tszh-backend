package adapters

import (
	"github.com/Vadim-12/tszh-backend/pkg/entity"
	"github.com/Vadim-12/tszh-backend/pkg/handler/dto"
)

func FromUserEntityToDto(entity *entity.User) *dto.GetMeResponseDto {
	return &dto.GetMeResponseDto{
		Id:          entity.ID,
		FullName:    entity.FullName,
		PhoneNumber: entity.PhoneNumber,
		Email:       *entity.Email,
	}
}
