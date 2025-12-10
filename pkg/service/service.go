package service

import (
	"context"

	"github.com/Vadim-12/tszh-backend/pkg/handler/dto"
	"github.com/Vadim-12/tszh-backend/pkg/repository"
	"github.com/Vadim-12/tszh-backend/pkg/utils"
)

type Authorization interface {
	SignUp(ctx context.Context, dto *dto.SignUpRequestDto) (*dto.SignUpResponseDto, error)
	SignIn(ctx context.Context, dto *dto.SignInRequestDto) (*dto.SignInResponseDto, error)
	Refresh(ctx context.Context, dto *dto.RefreshRequestDto) (*dto.RefreshResponseDto, error)
	Logout(ctx context.Context, dto *dto.LogoutRequestDto) (*dto.LogoutResponseDto, error)
}

type User interface{}

type Building interface{}

type Organization interface{}

type Service struct {
	Authorization
	User
	Building
	Organization
}

func NewService(repos *repository.Repository, utils *utils.Utils) *Service {
	return &Service{
		Authorization: NewAuthService(repos.User, repos.RefreshTokens, utils.Hasher, utils.JWTSigner),
		User:          NewUserService(repos.User),
		Building:      NewBuildingService(repos.Building),
		Organization:  NewOrganizationPostgres(repos.Organization),
	}
}
