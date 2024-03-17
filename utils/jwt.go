package utils

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/yashwantsinghcode/go_backend/constants"
)

var jwtAlgo = jwt.SigningMethodHS256

func SignJwt(claims map[string]interface{}) (string, error) {
	token := jwt.New(jwtAlgo)
	tokenClaims := token.Claims.(jwt.MapClaims)

	for key, element := range claims {
		tokenClaims[key] = element
	}

	return token.SignedString(constants.SECRET_KEY)
}

func VerifyJwt(tokenStr string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return constants.SECRET_KEY, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, err
	} else {
		return nil, errors.New("Access denied!")
	}
}
