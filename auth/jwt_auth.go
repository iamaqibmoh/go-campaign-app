package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type JWTAuth interface {
	GenerateToken(userID int) (string, error)
}

var SECRET_KEY = []byte("inisecretkeyyangsulitsyekali")

type JWTAuthImpl struct {
}

func NewJWTAuth() JWTAuth {
	return &JWTAuthImpl{}
}

func (auth *JWTAuthImpl) GenerateToken(userID int) (string, error) {
	payload := jwt.MapClaims{}
	payload["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
