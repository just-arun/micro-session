package util

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type jsonWebToken struct{}

func Jwt() *jsonWebToken {
	return &jsonWebToken{}
}

func (st *jsonWebToken) New(secret, sessionID string, expTime time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(expTime).Unix()
	claims["authorized"] = true
	claims["id"] = sessionID
	key := []byte(secret) 
	tokenString, err := token.SignedString(key)
	if err != nil {
		fmt.Println(tokenString)
		return "Signing Error", err
	}

	return tokenString, nil
}

func (st *jsonWebToken) VerifyJWT(secret, tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodECDSA)
		if !ok {
			return nil, errors.New("You're Unauthorized")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return false, err
	}

	return token.Valid, nil
}

func (st *jsonWebToken) ExtractClaims(secret, tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("there's an error with the signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "Error Parsing Token: ", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		username := claims["id"].(string)
		return username, nil
	}
	return "", fmt.Errorf("unable to extract claims")
}
