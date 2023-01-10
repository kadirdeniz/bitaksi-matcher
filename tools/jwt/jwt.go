package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

const JWT_SECRET = "secret"
const JWT_TTL = 3600

type JWT struct {
	DumpClaim jwt.MapClaims
	Token     string
}

func NewJWT(token string) *JWT {
	return &JWT{
		DumpClaim: jwt.MapClaims{
			"exp": JWT_TTL,
		},
		Token: token,
	}
}

func (_jwt *JWT) GetToken() string {
	return _jwt.Token
}

func (_jwt *JWT) GetIsAuthenticated() bool {
	return _jwt.DecodeToken().DumpClaim["authenticated"] == "true"
}

func (_jwt *JWT) DecodeToken() *JWT {
	jwt.ParseWithClaims(_jwt.GetToken(), _jwt.DumpClaim, keyFunc)
	return _jwt
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return JWT_SECRET, nil
}
