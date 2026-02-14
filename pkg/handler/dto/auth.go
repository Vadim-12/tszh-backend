package dto

type SignInRequestDto struct {
	PhoneNumber string `json:"phone_number" validate:"required,e164"`
	Password    string `json:"password" validate:"required,min=8,max=72"`
}

type UserResponseDto struct {
	Id          string  `json:"id"`
	FullName    string  `json:"full_name"`
	PhoneNumber string  `json:"phone_number"`
	Email       *string `json:"email"`
}

type TokensResponseDto struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type SignInResponseDto struct {
	User   UserResponseDto   `json:"user"`
	Tokens TokensResponseDto `json:"tokens"`
}

type SignUpRequestDto struct {
	FullName    string  `json:"full_name" validate:"required,min=2,max=100"`
	PhoneNumber string  `json:"phone_number" validate:"required,e164"`
	Email       *string `json:"email" validate:"omitempty,email"`
	Password    string  `json:"password" validate:"required,min=8,max=72"`
}

type SignUpResponseDto struct {
	User   UserResponseDto   `json:"user"`
	Tokens TokensResponseDto `json:"tokens"`
}

type RefreshRequestDto struct {
	RefreshToken string `json:"refresh_token" validate:"required,min=20"`
}

type RefreshResponseDto struct {
	Tokens TokensResponseDto `json:"tokens"`
}

type LogoutRequestDto struct {
	RefreshToken string `json:"refresh_token" validate:"required,min=20"`
}

type LogoutResponseDto struct {
	Status string `json:"status"`
}
