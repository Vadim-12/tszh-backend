package entity

import (
	"time"

	"github.com/google/uuid"
)

type UsersAndBuildingUnits struct {
	UserID         uuid.UUID `gorm:"type:uuid;primaryKey;index"`
	BuildingUnitID uuid.UUID `gorm:"type:uuid;primaryKey;index"`
	Role           string    `gorm:"not null;size:32"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`

	User         *User         `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	BuildingUnit *BuildingUnit `gorm:"foreignKey:BuildingUnitID;constraint:OnDelete:CASCADE"`
}

func (UsersAndBuildingUnits) TableName() string {
	return "users_and_building_units"
}
