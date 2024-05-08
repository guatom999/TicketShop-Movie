package jwtauth

import "github.com/golang-jwt/jwt/v4"

type (
	customerClaims struct {
		Id      string `json:"customer_id"`
		Email   string `json:"email"`
		UserNam string `json:"username"`
		jwt.RegisteredClaims
	}
)

// func NewAccessToken(secret string, expireAt int64, claims *customerClaims) string {

// 	// claims := &customerClaims{}

// 	return ""
// }
