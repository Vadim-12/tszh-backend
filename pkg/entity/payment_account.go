package entity

import (
	"time"

	"github.com/google/uuid"
)

type PaymentAccount struct {
	ID                   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Number               int       `gorm:"not null"`
	BIC                  string    `gorm:"not null"`
	BankName             string    `gorm:"not null"`
	CorrespondentAccount string    `gorm:"not null"`
	CreatedAt            time.Time `gorm:"autoCreateTime"`
	UpdatedAt            time.Time `gorm:"autoUpdateTime"`
}

// это все потом (все банковские приколы и оплаты)
// со счетами на оплату тоже пока ничего не делаем
