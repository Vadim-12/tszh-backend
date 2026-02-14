package repository

import (
	"context"

	"github.com/Vadim-12/tszh-backend/pkg/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User interface {
	CreateUser(ctx context.Context, dto *entity.User) (*entity.User, error)
	FindByID(ctx context.Context, userId uuid.UUID) (*entity.User, error)
	FindByPhoneNumber(ctx context.Context, phoneNumber string) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}

type RefreshTokens interface {
	Save(ctx context.Context, token *entity.RefreshToken) error
	GetByID(ctx context.Context, id uuid.UUID) (*entity.RefreshToken, error)
	DeleteByID(ctx context.Context, tokenId uuid.UUID) error
	DeleteAllByUserID(ctx context.Context, userId uuid.UUID) error
}

type Building interface{}

type Organization interface {
	GetAllOrganizations(ctx context.Context) ([]entity.Organization, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Organization, error)
	Create(ctx context.Context, creationDto *entity.Organization, ownerUserID uuid.UUID) (*entity.Organization, error)
	GetByINN(ctx context.Context, INN string) (*entity.Organization, error)
}

type Health interface {
	Ping(ctx context.Context) error
}

type Repository struct {
	User
	RefreshTokens
	Building
	Organization
	Health
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:          NewUserPostgres(db),
		RefreshTokens: NewRefreshTokenPostgres(db),
		Building:      NewBuildingPostgres(db),
		Organization:  NewOrganizationPostgres(db),
		Health:        NewHealthPostgres(db),
	}
}
