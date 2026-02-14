package entity

import (
	"time"

	"github.com/google/uuid"
)

type BuildingUnit struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Number     string    `gorm:"not null"`         // № квартиры / гаража / бокса
	UnitType   string    `gorm:"not null;size:32"` // "apartment", "garage", "storage", etc
	TotalArea  float64   `gorm:"not null"`
	LivingArea *float64  `gorm:"type:numeric"` // только для жилых помещений
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`

	BuildingID uuid.UUID `gorm:"type:uuid;index;not null"`
	Building   *Building `gorm:"foreignKey:BuildingID;constraint:OnDelete:CASCADE"`
}
