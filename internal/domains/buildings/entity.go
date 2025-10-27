package buildings

import (
	"errors"

	"github.com/google/uuid"
)

var ErrNotFound = errors.New("building not found")

type Building struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Number    string    `gorm:"not null"             json:"number"`
	Floor     int       `gorm:"not null"             json:"floor"`
	OwnerName string    `gorm:"not null"             json:"owner_name"`

	// unix timestamps (удобно для простоты)
	CreatedAt int64 `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64 `gorm:"autoUpdateTime" json:"updated_at"`
}
