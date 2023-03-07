package authonticator

import (
	"INi-Wallet2/dto"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	GenerateToken(cred dto.LoginRequestBody) (string, error)
	ValidateToken(token string) error
}

type MyClaims struct {
	jwt.StandardClaims
	Email string `json:"Email"`
}

var jwtKey = []byte(("SECRET_KEY"))

func GenerateToken(cred dto.LoginRequestBody) (string, error) {
	now := time.Now().UTC()
	end := now.Add(60 * time.Second)
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: end.Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    "ISSUER",
		},
		Email: cred.Email,
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, claims,
	)
	fmt.Println(token.SignedString(jwtKey))
	return token.SignedString(jwtKey)

}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&MyClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*MyClaims)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
