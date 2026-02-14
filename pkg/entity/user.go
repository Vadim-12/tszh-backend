package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	FullName     string    `gorm:"not null"`
	Email        *string   `gorm:"uniqueIndex"`
	PhoneNumber  string    `gorm:"uniqueIndex;not null"`
	PasswordHash string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`

	RefreshTokens      []*RefreshToken      `gorm:"foreignKey:UserID"`
	OrganizationsLinks []*OrganizationStaff `gorm:"foreignKey:UserID"`
}
