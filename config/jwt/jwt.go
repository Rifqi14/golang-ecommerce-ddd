package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtCredential struct {
	TokenSecret         string
	ExpiredToken        int
	RefreshTokenSecret  string
	ExpiredRefreshToken int
}

type CustomClaims struct {
	jwt.StandardClaims
	Payload string `json:"payload"`
}

func (c JwtCredential) GetToken(issuer, payload string) (string, int64, error) {
	expirationTime := time.Now().Add(time.Duration(c.ExpiredToken) * time.Minute).Unix()

	claims := &CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
			Issuer:    issuer,
		},
		Payload: payload,
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawToken.SignedString([]byte(c.TokenSecret))
	if err != nil {
		return "", 0, err
	}

	return token, expirationTime, nil
}

func (c JwtCredential) GetRefreshToken(issuer, payload string) (string, int64, error) {
	expirationTime := time.Now().Add(time.Duration(c.ExpiredRefreshToken) * time.Minute).Unix()

	claims := &CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
			Issuer:    issuer,
		},
		Payload: payload,
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawToken.SignedString([]byte(c.RefreshTokenSecret))
	if err != nil {
		return "", 0, err
	}

	return token, expirationTime, nil
}
