package entity

import (
	"time"

	"github.com/google/uuid"
)

type OrganizationStaff struct {
	OrganizationID uuid.UUID `gorm:"type:uuid;primaryKey;index" json:"organization_id"`
	UserID         uuid.UUID `gorm:"type:uuid;primaryKey;index" json:"user_id"`
	Role           string    `gorm:"not null"             json:"role"`

	CreatedAt time.Time `gorm:"autoCreateTime"            json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"            json:"updated_at"`

	User         *User         `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"         json:"-"`
	Organization *Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE" json:"-"`
}

func (OrganizationStaff) TableName() string {
	return "organizations_staff"
}
