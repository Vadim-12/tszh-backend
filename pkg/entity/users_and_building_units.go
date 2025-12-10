package entity

import (
	"time"

	"github.com/google/uuid"
)

type UsersAndBuildingUnits struct {
	UserID         uuid.UUID `gorm:"type:uuid;primaryKey;index" json:"user_id"`
	BuildingUnitID uuid.UUID `gorm:"type:uuid;primaryKey;index" json:"building_unit_id"`
	Role           string    `gorm:"not null;size:32"           json:"role"`

	CreatedAt time.Time `gorm:"autoCreateTime"            json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"            json:"updated_at"`

	User         *User         `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"         json:"-"`
	BuildingUnit *BuildingUnit `gorm:"foreignKey:BuildingUnitID;constraint:OnDelete:CASCADE" json:"-"`
}

func (UsersAndBuildingUnits) TableName() string {
	return "users_and_building_units"
}
