package organizations

import (
	"errors"

	"github.com/google/uuid"
)

var ErrNotFound = errors.New("organization not found")

type Organization struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name         string    `gorm:"not null" json:"name"`
	EmailAddress string    `gorm:"not null" json:"email_address"`
	PostAddress  string    `gorm:"not null" json:"post_address"`
	INN          string    `gorm:"type:char(11);not null" json:"inn"`
}
