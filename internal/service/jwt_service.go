package service

import (
	"errors"
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
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString([]byte(j.secret)) //sign the token using the secret and then return it
}

func (j *JWTService) Validate(tokenString string) (uint, error) { //Function to verify the jwt and excract the UserID
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { //parse function is used to check if the token was signed by the secret, also decodes the signed jwt
		return []byte(j.secret), nil //returns the secret key
	})

	if err != nil {
		return 0, err
	}

	//token.Claims contains the JWT payload. .(jwt.MapClaims) converts it to a MapClaims so you can access values like user_id.
	claims, ok := token.Claims.(jwt.MapClaims) //Extracts the jwt payload
	if !ok || !token.Valid {                   //ensures if the payload is valid
		return 0, errors.New("Invalid Token")
	}

	userId := uint(claims["user_id"].(float64)) //read the user id from the payload

	return userId, nil

}
