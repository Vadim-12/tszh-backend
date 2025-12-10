package entity

import (
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`

	FullTitle    string `gorm:"not null"             json:"full_title"`
	ShortTitle   string `gorm:"uniqueIndex"          json:"short_title"`
	INN          string `gorm:"uniqueIndex;not null" json:"inn"`
	Email        string `gorm:"not null"             json:"email"`
	LegalAddress string `gorm:"not null"             json:"legal_address"`

	CreatedAt time.Time `gorm:"autoCreateTime"       json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"       json:"updated_at"`

	StaffLinks []*OrganizationStaff `gorm:"foreignKey:OrganizationID" json:"-"`
}
