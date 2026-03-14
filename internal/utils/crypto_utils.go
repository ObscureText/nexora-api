package utils

//go:generate mockgen -source=crypto_util.go -destination=mocks/mock_crypto_util.go -package=mocks

import (
	nexora_error "nexora-api/internal/errors"

	"golang.org/x/crypto/bcrypt"
)

type CryptoUtil interface {
	HashPassword(password string) (string, *nexora_error.NexoraError)
	ComparePasswordAndHash(password string, hash string) bool
}

type cryptoUtil struct{}

func NewCryptoUtils() CryptoUtil {
	return &cryptoUtil{}
}

func (_ *cryptoUtil) HashPassword(password string) (string, *nexora_error.NexoraError) {
	hash, hashErr := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if hashErr != nil {
		return "", nexora_error.NewInternalServerNexoraError("CryptoUtil: HashPassword: " + hashErr.Error())
	}

	return string(hash), nil
}

func (_ *cryptoUtil) ComparePasswordAndHash(password string, hash string) bool {
	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	) == nil
}
