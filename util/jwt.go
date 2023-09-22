package util

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type jsonWebToken struct{}

func Jwt() *jsonWebToken {
	return &jsonWebToken{}
}

func (st *jsonWebToken) New(secret, sessionID string, roles []string, expTime time.Duration) (jwtToken string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(expTime).Unix()
	claims["authorized"] = true
	claims["id"] = sessionID
	claims["roles"] = strings.Join(roles, ",")
	key := []byte(secret)
	tokenString, err := token.SignedString(key)
	if err != nil {
		fmt.Println(tokenString)
		return "Signing Error", err
	}

	return tokenString, nil
}

func (st *jsonWebToken) VerifyJWT(secret, tokenString string) (parsedToken *jwt.Token, err error) {
	parsedToken, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodECDSA)
		if !ok {
			return nil, errors.New("you're unauthorized")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	return parsedToken, nil
}

func (st *jsonWebToken) ExtractClaims(secret, tokenString string) (sessionID string, roles []string, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("there's an error with the signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "Error Parsing Token: ", []string{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		username := claims["id"].(string)
		roles := strings.Split(claims["roles"].(string), ",")
		return username, roles, nil
	}
	return "", []string{}, fmt.Errorf("unable to extract claims")
}
