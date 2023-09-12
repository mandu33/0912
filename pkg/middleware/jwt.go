package middleware

import (
	"time"

	"github.com/form3tech-oss/jwt-go"
)

var jwtKey = []byte("this is a secret")

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

func CreateToken(userID int64) (string, error) {
	expireTime := time.Now().Add(7200 * 24 * time.Hour)

	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "douyin",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
