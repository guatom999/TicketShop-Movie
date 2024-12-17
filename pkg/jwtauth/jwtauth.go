package jwtauth

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

type (
	CustomerClaims struct {
		Id      string `json:"customer_id"`
		Email   string `json:"email"`
		UserNam string `json:"username"`
		jwt.RegisteredClaims
	}
)

func ParseToken(secret string, tokenString string) (*CustomerClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &CustomerClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error: unexpected signing method")
		}

		return []byte(secret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, errors.New("error: token format is invalid")
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("error: token is expired")
		} else {
			return nil, errors.New("error: token is invalid")
		}
	}

	if claims, ok := token.Claims.(*CustomerClaims); ok {
		return claims, nil
	} else {
		return nil, errors.New("error: claims type is invalid")

	}

}
