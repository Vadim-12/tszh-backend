package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`

	FullName     string `gorm:"not null"             json:"full_name"`
	Email        string `gorm:"uniqueIndex"          json:"email"`
	PhoneNumber  string `gorm:"uniqueIndex;not null" json:"phone_number"`
	PasswordHash string `gorm:"not null"             json:"password_hash"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	RefreshTokens      []*RefreshToken      `gorm:"foreignKey:UserID" json:"-"`
	OrganizationsLinks []*OrganizationStaff `gorm:"foreignKey:UserID" json:"-"`
}
