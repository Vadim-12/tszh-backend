package users

import (
	"errors"

	"github.com/google/uuid"
)

var ErrNotFound = errors.New("user not found")

type User struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	FullName    string    `gorm:"not null"             json:"full_name"`
	Email       string    `gorm:"not null"             json:"email"`
	PhoneNumber string    `gorm:"not null"             json:"phone_number"`

	// unix timestamps (удобно для простоты)
	CreatedAt int64 `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64 `gorm:"autoUpdateTime" json:"updated_at"`
}
