package utils

import (
	"time"

	config "github.com/Vadim-12/tszh-backend/pkg/config/utils"
	"github.com/google/uuid"
)

type Utils struct {
	Hasher
	JWTSigner
}

type Hasher interface {
	Hash(password string) (string, error)
	Verify(hash, password string) error
}

type JWTSigner interface {
	SignAccess(userID uuid.UUID, now time.Time) (string, time.Time, error)
	SignRefresh(jti, userID uuid.UUID, now time.Time) (string, time.Time, error)
	ParseAccess(tokenStr string) (uuid.UUID, error)
	ParseRefresh(tokenStr string) (uuid.UUID, uuid.UUID, error)
}

func NewUtils(hasherCost config.BcryptHasher, jwtSignerConfig config.JWTSigner) *Utils {
	return &Utils{
		Hasher:    NewBcryptHasherUtil(hasherCost),
		JWTSigner: NewJWTSignerUtil(jwtSignerConfig),
	}
}
