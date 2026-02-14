package dto

import (
	"time"

	"github.com/google/uuid"
)

type OrganizationDto struct {
	ID           uuid.UUID `json:"id"`
	FullTitle    string    `json:"full_title"`
	ShortTitle   string    `json:"short_title"`
	INN          string    `json:"INN"`
	Email        string    `json:"email"`
	LegalAddress string    `json:"legal_address"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateOrganizationPayload struct {
	FullTitle    string  `json:"full_title" validate:"required"`
	ShortTitle   string  `json:"short_title" validate:"required"`
	INN          string  `json:"inn" validate:"required,len=10|len=12,numeric"`
	Email        *string `json:"email" validate:"omitempty,email"`
	LegalAddress string  `json:"legal_address" validate:"required"`
}
