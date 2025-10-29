package buildings

import (
	"errors"

	"github.com/Vadim-12/tszh-backend/internal/platform/model"
)

var ErrNotFound = errors.New("building not found")

type BuildingModel struct {
	model.BaseModel
	Number    string `gorm:"not null"             json:"number"`
	Floor     int    `gorm:"not null"             json:"floor"`
	OwnerName string `gorm:"not null"             json:"owner_name"`
}
