package service

import (
	"context"
	"time"

	"github.com/Vadim-12/tszh-backend/pkg/entity"
	"github.com/Vadim-12/tszh-backend/pkg/errors"
	"github.com/Vadim-12/tszh-backend/pkg/handler/dto"
	"github.com/Vadim-12/tszh-backend/pkg/repository"
	"github.com/Vadim-12/tszh-backend/pkg/utils"
	"github.com/google/uuid"
)

type AuthService struct {
	userRepo  repository.User
	tokenRepo repository.RefreshTokens
	hasher    utils.Hasher
	jwtSigner utils.JWTSigner
}

func NewAuthService(userRepo repository.User, tokenRepo repository.RefreshTokens, hasher utils.Hasher, jwtSigner utils.JWTSigner) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
		hasher:    hasher,
		jwtSigner: jwtSigner,
	}
}

func (s *AuthService) SignUp(ctx context.Context, signUpDto *dto.SignUpRequestDto) (*dto.SignUpResponseDto, error) {
	existedUser, err := s.userRepo.FindByPhoneNumber(ctx, signUpDto.PhoneNumber)
	if err != nil {
		return nil, err
	}
	if existedUser != nil {
		return nil, errors.ErrUserWithPhoneNumberAlreadyExists
	}

	existedUser, err = s.userRepo.FindByEmail(ctx, signUpDto.Email)
	if err != nil {
		return nil, err
	}
	if existedUser != nil {
		return nil, errors.ErrUserWithEmailAlreadyExists
	}

	passwordHash, err := s.hasher.Hash(signUpDto.Password)
	if err != nil {
		return nil, err
	}
	user := &entity.User{
		FullName:     signUpDto.FullName,
		PhoneNumber:  signUpDto.PhoneNumber,
		Email:        signUpDto.Email,
		PasswordHash: passwordHash,
	}
	user, err = s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	accessToken, _, err := s.jwtSigner.SignAccess(user.ID, now)
	if err != nil {
		return nil, err
	}

	jti := uuid.New()
	refreshToken, refreshExp, err := s.jwtSigner.SignRefresh(jti, user.ID, now)
	if err != nil {
		return nil, err
	}

	rt := &entity.RefreshToken{
		ID:        jti,
		UserID:    user.ID,
		ExpiresAt: refreshExp,
	}
	if err := s.tokenRepo.Save(ctx, rt); err != nil {
		return nil, err
	}

	response := dto.SignUpResponseDto{
		User: dto.UserResponseDto{
			Id:       user.ID.String(),
			FullName: user.FullName,
		},
		Tokens: dto.TokensResponseDto{
			Access:  accessToken,
			Refresh: refreshToken,
		},
	}
	return &response, nil
}

func (s *AuthService) SignIn(ctx context.Context, signInDto *dto.SignInRequestDto) (*dto.SignInResponseDto, error) {
	user, err := s.userRepo.FindByPhoneNumber(ctx, signInDto.PhoneNumber)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.ErrInvalidCredentials
	}

	if err := s.hasher.Verify(user.PasswordHash, signInDto.Password); err != nil {
		return nil, errors.ErrInvalidCredentials
	}

	now := time.Now()
	accessToken, _, err := s.jwtSigner.SignAccess(user.ID, now)
	if err != nil {
		return nil, err
	}

	jti := uuid.New()
	refreshToken, refreshExp, err := s.jwtSigner.SignRefresh(jti, user.ID, now)
	if err != nil {
		return nil, err
	}

	rt := &entity.RefreshToken{
		ID:        jti,
		UserID:    user.ID,
		ExpiresAt: refreshExp,
	}
	if err := s.tokenRepo.Save(ctx, rt); err != nil {
		return nil, err
	}

	response := dto.SignInResponseDto{
		User: dto.UserResponseDto{
			Id:       user.ID.String(),
			FullName: user.FullName,
		},
		Tokens: dto.TokensResponseDto{
			Access:  accessToken,
			Refresh: refreshToken,
		},
	}
	return &response, nil
}

func (s *AuthService) Refresh(ctx context.Context, refreshDto *dto.RefreshRequestDto) (*dto.RefreshResponseDto, error) {
	jti, userId, err := s.jwtSigner.ParseRefresh(refreshDto.RefreshToken)
	if err != nil {
		return nil, errors.ErrRefreshInvalid
	}

	stored, err := s.tokenRepo.GetByID(ctx, jti)
	if err != nil {
		return nil, err
	}
	if stored == nil || stored.UserID != userId || stored.ExpiresAt.Before(time.Now()) {
		return nil, errors.ErrRefreshInvalid
	}

	if err := s.tokenRepo.DeleteByID(ctx, jti); err != nil {
		return nil, err
	}

	now := time.Now()
	access, _, err := s.jwtSigner.SignAccess(userId, now)
	if err != nil {
		return nil, err
	}

	newJTI := uuid.New()
	newRefresh, newRefreshExp, err := s.jwtSigner.SignRefresh(newJTI, userId, now)
	if err != nil {
		return nil, err
	}

	rt := &entity.RefreshToken{
		ID:        newJTI,
		UserID:    userId,
		ExpiresAt: newRefreshExp,
	}
	if err := s.tokenRepo.Save(ctx, rt); err != nil {
		return nil, err
	}

	response := &dto.RefreshResponseDto{
		Tokens: dto.TokensResponseDto{
			Access:  access,
			Refresh: newRefresh,
		},
	}

	return response, nil
}

func (s *AuthService) Logout(ctx context.Context, logoutDto *dto.LogoutRequestDto) (*dto.LogoutResponseDto, error) {
	jti, _, err := s.jwtSigner.ParseRefresh(logoutDto.RefreshToken)
	if err != nil {
		return nil, errors.ErrRefreshInvalid
	}

	if err := s.tokenRepo.DeleteByID(ctx, jti); err != nil {
		return nil, err
	}

	resp := &dto.LogoutResponseDto{
		Status: "success",
	}
	return resp, nil
}
