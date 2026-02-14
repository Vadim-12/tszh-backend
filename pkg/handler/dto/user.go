package dto

import "github.com/google/uuid"

type GetMeResponseDto struct {
	Id          uuid.UUID `json:"id"`
	FullName    string    `json:"full_name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
}
