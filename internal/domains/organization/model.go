package organizations

import (
	"errors"

	"github.com/Vadim-12/tszh-backend/internal/platform/model"
)

var ErrNotFound = errors.New("organization not found")

type OrganizationModel struct {
	model.BaseModel
	Name         string `gorm:"not null" json:"name"`
	EmailAddress string `gorm:"not null" json:"email_address"`
	PostAddress  string `gorm:"not null" json:"post_address"`
	INN          string `gorm:"type:char(11);not null" json:"inn"`
}
