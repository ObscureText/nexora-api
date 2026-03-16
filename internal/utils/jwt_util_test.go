package utils_test

import (
	"os"
	"testing"
	"time"

	"nexora-api/internal/domain"
	"nexora-api/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/suite"
)

type JwtUtilTestSuite struct {
	suite.Suite
	jwtUtil utils.JwtUtil
}

func (suite *JwtUtilTestSuite) SetupTest() {
	os.Setenv("JWT_SECRET", "test-secret")
	suite.jwtUtil = utils.NewJWTUtil()
}

func TestJwtUtilSuite(t *testing.T) {
	suite.Run(t, new(JwtUtilTestSuite))
}

func (suite *JwtUtilTestSuite) TestGenerateAndParseToken() {
	token, err := suite.jwtUtil.GenerateToken("test-user-id", domain.EstablishmentOwner)

	suite.Require().NoError(err)
	suite.Require().NotEmpty(token)

	parsedUserID, parsedRole, err := suite.jwtUtil.ParseToken(token)

	suite.Require().NoError(err)
	suite.Equal("test-user-id", parsedUserID)
	suite.Equal(domain.EstablishmentOwner, parsedRole)
}

func (suite *JwtUtilTestSuite) TestParseToken_InvalidSignature() {
	os.Setenv("JWT_SECRET", "secret1")
	jwtUtil1 := utils.NewJWTUtil()

	token, err := jwtUtil1.GenerateToken("user1", domain.Individual)
	suite.Require().NoError(err)

	os.Setenv("JWT_SECRET", "secret2")
	jwtUtil2 := utils.NewJWTUtil()

	_, _, err = jwtUtil2.ParseToken(token)

	suite.Error(err)
}

func (suite *JwtUtilTestSuite) TestParseToken_ExpiredToken() {
	claims := jwt.MapClaims{
		"user_id":   "user1",
		"user_role": domain.Individual,
		"exp":       time.Now().Add(-1 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte("test-secret"))
	suite.Require().NoError(err)

	_, _, err = suite.jwtUtil.ParseToken(signed)

	suite.Error(err)
}

func (suite *JwtUtilTestSuite) TestParseToken_MalformedToken() {
	_, _, err := suite.jwtUtil.ParseToken("invalid.token.value")
	suite.Error(err)
}

func (suite *JwtUtilTestSuite) TestParseToken_InvalidSigningMethod() {
	claims := jwt.MapClaims{
		"user_id":   "user1",
		"user_role": domain.Individual,
		"exp":       time.Now().Add(1 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodNone, claims)

	signed, err := token.SignedString(jwt.UnsafeAllowNoneSignatureType)
	suite.Require().NoError(err)

	_, _, err = suite.jwtUtil.ParseToken(signed)
	suite.Error(err)
}
