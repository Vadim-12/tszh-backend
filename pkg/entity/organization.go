package entity

import (
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	FullTitle    string    `gorm:"not null"`
	ShortTitle   string    `gorm:"not null;index"`
	INN          string    `gorm:"not null;uniqueIndex;size:12"`
	Email        *string   `gorm:"type:text"`
	LegalAddress string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`

	StaffLinks []*OrganizationStaff `gorm:"foreignKey:OrganizationID"`
}
