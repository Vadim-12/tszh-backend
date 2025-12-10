package utils

import (
	config "github.com/Vadim-12/tszh-backend/pkg/config/utils"
	"golang.org/x/crypto/bcrypt"
)

type BcryptHasherUtil struct {
	cost int
}

func NewBcryptHasherUtil(config config.BcryptHasher) *BcryptHasherUtil {
	if config.Cost == 0 {
		config.Cost = bcrypt.DefaultCost
	}
	return &BcryptHasherUtil{cost: config.Cost}
}

func (h *BcryptHasherUtil) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	return string(hash), err
}

func (h *BcryptHasherUtil) Verify(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
