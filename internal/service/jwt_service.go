package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secret string //Secret used to sign and verify JWTs
}

func NewJWTService(secret string) *JWTService {
	return &JWTService{
		secret: secret,
	}
}

func (j *JWTService) Generate(userID uint) (string, error) {
	claims := jwt.MapClaims{ //standard jwt payload
		"user_id": userID,                           //store the logged-in User's ID
		"exp":     time.Now().Add(time.Hour).Unix(), //token expires in an hour
	}

	token := jwt.NewWithClaims( //Create a jwt using the HS256 signing algorithm
		jwt.SigningMethodES256,
		claims,
	)

	return token.SignedString([]byte(j.secret)) //sign the token using the secret and then return it
}
