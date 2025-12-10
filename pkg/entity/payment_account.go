package entity

import (
	"time"

	"github.com/google/uuid"
)

type PaymentAccount struct {
	ID                   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Number               int       `gorm:"not null"       json:"number"`
	BIC                  string    `gorm:"not null"       json:"bic"`
	BankName             string    `gorm:"not null"       json:"bank_name"`
	CorrespondentAccount string    `gorm:"not null"       json:"correspondent_account"`
	CreatedAt            time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt            time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// это все потом (все банковские приколы и оплаты)
// со счетами на оплату тоже пока ничего не делаем
