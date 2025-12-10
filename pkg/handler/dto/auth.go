package dto

type SignInRequestDto struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UserResponseDto struct {
	Id       string `json:"id"`
	FullName string `json:"full_name"`
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
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type SignUpResponseDto struct {
	User   UserResponseDto   `json:"user"`
	Tokens TokensResponseDto `json:"tokens"`
}

type RefreshRequestDto struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponseDto struct {
	Tokens TokensResponseDto `json:"tokens"`
}

type LogoutRequestDto struct {
	RefreshToken string `json:"refresh_token"`
}

type LogoutResponseDto struct {
	Status string `json:"status"`
}
