package jwtauth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/guatom999/TicketShop-Movie/utils"
)

type (
	AuthInterface interface {
		SignToken() string
	}
	Claims struct {
		CustomerId string `json:"customer_id"`
	}

	CustomerClaims struct {
		*Claims
		jwt.RegisteredClaims
	}

	authConcrete struct {
		Secret []byte
		Claims *CustomerClaims
	}

	accessToken struct {
		*authConcrete
	}

	refreshToken struct {
		*authConcrete
	}
)

func (a *authConcrete) SignToken() string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a.Claims)

	tokenString, _ := token.SignedString(a.Claims)

	return tokenString
}

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

func NewAccessToken(secret string, expireAt int64, claims *Claims) AuthInterface {
	return &accessToken{
		authConcrete: &authConcrete{
			Secret: []byte(secret),
			Claims: &CustomerClaims{
				Claims: claims,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "seeyarnmirttown.com",
					Subject:   "access-token",
					Audience:  []string{"seeyarnmirttown.com"},
					ExpiresAt: jwt.NewNumericDate(utils.GetLocaltime().Add(time.Second * 20)),
					NotBefore: jwt.NewNumericDate(utils.GetLocaltime()),
					IssuedAt:  jwt.NewNumericDate(utils.GetLocaltime()),
				},
			},
		},
	}
}

func NewRefreshToken(secret string, expireAt int64, claims *Claims) AuthInterface {
	return &refreshToken{
		authConcrete: &authConcrete{
			Secret: []byte(secret),
			Claims: &CustomerClaims{
				Claims: claims,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "seeyarnmirttown.com",
					Subject:   "access-token",
					Audience:  []string{"seeyarnmirttown.com"},
					ExpiresAt: jwt.NewNumericDate(utils.GetLocaltime().Add(time.Second * 20)),
					NotBefore: jwt.NewNumericDate(utils.GetLocaltime()),
					IssuedAt:  jwt.NewNumericDate(utils.GetLocaltime()),
				},
			},
		},
	}
}

func ReloadToken(secret string, expireAt int64, claims *Claims) AuthInterface {
	return &refreshToken{
		authConcrete: &authConcrete{
			Secret: []byte(secret),
			Claims: &CustomerClaims{
				Claims: claims,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "seeyarnmirttown.com",
					Subject:   "access-token",
					Audience:  []string{"seeyarnmirttown.com"},
					ExpiresAt: jwt.NewNumericDate(utils.GetLocaltime().Add(time.Second * 20)),
					NotBefore: jwt.NewNumericDate(utils.GetLocaltime()),
					IssuedAt:  jwt.NewNumericDate(utils.GetLocaltime()),
				},
			},
		},
	}
}
