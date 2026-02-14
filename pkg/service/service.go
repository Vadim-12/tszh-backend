package service

import (
	"context"

	"github.com/Vadim-12/tszh-backend/pkg/handler/dto"
	"github.com/Vadim-12/tszh-backend/pkg/repository"
	"github.com/Vadim-12/tszh-backend/pkg/utils"
	"github.com/google/uuid"
)

type Authorization interface {
	SignUp(ctx context.Context, dto *dto.SignUpRequestDto) (*dto.SignUpResponseDto, error)
	SignIn(ctx context.Context, dto *dto.SignInRequestDto) (*dto.SignInResponseDto, error)
	Refresh(ctx context.Context, dto *dto.RefreshRequestDto) (*dto.RefreshResponseDto, error)
	Logout(ctx context.Context, dto *dto.LogoutRequestDto) (*dto.LogoutResponseDto, error)
}

type User interface {
	GetMe(ctx context.Context, userID uuid.UUID) (*dto.GetMeResponseDto, error)
}

type Building interface{}

type Organization interface {
	GetAll(ctx context.Context) ([]dto.OrganizationDto, error)
	GetByID(ctx context.Context, id uuid.UUID) (*dto.OrganizationDto, error)
	Create(ctx context.Context, dto *dto.CreateOrganizationPayload, userID uuid.UUID) (*dto.OrganizationDto, error)
}

type Health interface {
	Ping(ctx context.Context) error
}

type Service struct {
	Authorization
	User
	Building
	Organization
	Health
}

func NewService(repos *repository.Repository, utils *utils.Utils) *Service {
	return &Service{
		Authorization: NewAuthService(repos.User, repos.RefreshTokens, utils.Hasher, utils.JWTSigner),
		User:          NewUserService(repos.User),
		Building:      NewBuildingService(repos.Building),
		Organization:  NewOrganizationService(repos.Organization),
		Health:        NewHealthService(repos.Health),
	}
}
