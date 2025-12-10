package utils

import (
	"errors"
	"time"

	config "github.com/Vadim-12/tszh-backend/pkg/config/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTSignerUtil struct {
	AccessSecret  []byte
	RefreshSecret []byte
	AccessTTL     time.Duration
	RefreshTTL    time.Duration
}

func NewJWTSignerUtil(config config.JWTSigner) *JWTSignerUtil {
	return &JWTSignerUtil{
		AccessSecret:  config.AccessSecret,
		RefreshSecret: config.RefreshSecret,
		AccessTTL:     config.AccessTTL,
		RefreshTTL:    config.RefreshTTL,
	}
}

func (s *JWTSignerUtil) SignAccess(userID uuid.UUID, now time.Time) (string, time.Time, error) {
	exp := now.Add(s.AccessTTL)
	claims := jwt.MapClaims{
		"sub": userID.String(),
		"exp": exp.Unix(),
		"typ": "access",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(s.AccessSecret)
	return signed, exp, err
}

func (s *JWTSignerUtil) SignRefresh(jti, userID uuid.UUID, now time.Time) (string, time.Time, error) {
	exp := now.Add(s.RefreshTTL)
	claims := jwt.MapClaims{
		"jti": jti.String(),
		"sub": userID.String(),
		"exp": exp.Unix(),
		"typ": "refresh",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(s.RefreshSecret)
	return signed, exp, err
}

func (s *JWTSignerUtil) ParseAccess(tokenStr string) (uuid.UUID, error) {
	t, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("alg mismatch")
		}
		return s.AccessSecret, nil
	})
	if err != nil || !t.Valid {
		return uuid.Nil, errors.New("invalid access token")
	}
	sub, err := t.Claims.GetSubject()
	if err != nil {
		return uuid.Nil, err
	}
	return uuid.Parse(sub)
}

func (s *JWTSignerUtil) ParseRefresh(tokenStr string) (uuid.UUID, uuid.UUID, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("alg mismatch")
		}
		return s.RefreshSecret, nil
	})
	if err != nil || !token.Valid {
		return uuid.Nil, uuid.Nil, errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, uuid.Nil, errors.New("invalid claims")
	}

	sub, _ := claims["sub"].(string)
	jti, _ := claims["jti"].(string)

	uid, err := uuid.Parse(sub)
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}
	jid, err := uuid.Parse(jti)
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}
	return jid, uid, nil
}
