package utils

import (
	"nexora-api/internal/domain"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtClaims struct {
	UserId   string          `json:"user_id"`
	UserRole domain.UserRole `json:"user_role"`

	jwt.RegisteredClaims
}

type JwtUtil interface {
	GenerateToken(userId string, role domain.UserRole) (string, error)
	ParseToken(token string) (string, domain.UserRole, error)
}

type jwtUtil struct {
	secret []byte
}

func NewJWTUtil() JwtUtil {
	secret := os.Getenv("JWT_SECRET")

	return &jwtUtil{
		secret: []byte(secret),
	}
}

func (util *jwtUtil) GenerateToken(userId string, role domain.UserRole) (string, error) {
	claims := jwtClaims{
		UserId:   userId,
		UserRole: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(util.secret)
}

func (util *jwtUtil) ParseToken(tokenString string) (string, domain.UserRole, error) {
	claims := &jwtClaims{}

	_, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenInvalidClaims
			}
			return util.secret, nil
		},
	)

	if err != nil {
		return "", "", err
	}

	return claims.UserId, claims.UserRole, nil
}
