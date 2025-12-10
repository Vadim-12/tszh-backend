package entity

import (
	"time"

	"github.com/google/uuid"
)

type BuildingUnit struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`

	Number     string   `gorm:"not null"         json:"number"`    // № квартиры / гаража / бокса
	UnitType   string   `gorm:"not null;size:32" json:"unit_type"` // "apartment", "garage", "storage", etc
	TotalArea  float64  `gorm:"not null"         json:"total_area"`
	LivingArea *float64 `gorm:"type:numeric"     json:"living_area,omitempty"` // только для жилых помещений

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	BuildingID uuid.UUID `gorm:"type:uuid;index;not null"                          json:"building_id"`
	Building   *Building `gorm:"foreignKey:BuildingID;constraint:OnDelete:CASCADE" json:"-"`
}
