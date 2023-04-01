package auth

import (
	"bwa-campaign-app/helper"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type JWTAuth interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

var SecretKey = []byte(helper.GetEnv("SECRET_KEY"))

type JWTAuthImpl struct {
}

func (auth *JWTAuthImpl) ValidateToken(token string) (*jwt.Token, error) {
	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}

		return SecretKey, nil
	})

	if err != nil {
		return parseToken, err
	}

	return parseToken, nil
}

func (auth *JWTAuthImpl) GenerateToken(userID int) (string, error) {
	payload := jwt.MapClaims{}
	payload["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signedToken, err := token.SignedString(SecretKey)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func NewJWTAuth() JWTAuth {
	return &JWTAuthImpl{}
}
