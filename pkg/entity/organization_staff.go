package entity

import (
	"time"

	"github.com/google/uuid"
)

type OrganizationStaff struct {
	OrganizationID uuid.UUID `gorm:"type:uuid;primaryKey;index"`
	UserID         uuid.UUID `gorm:"type:uuid;primaryKey;index"`
	Role           string    `gorm:"not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`

	User         *User         `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Organization *Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE"`
}

func (OrganizationStaff) TableName() string {
	return "organizations_staff"
}
