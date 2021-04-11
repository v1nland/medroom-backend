package utils

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func DecodificarToken(token_str string, env_entity string) string {
	cnt := 0
	for i := 0; i < len(token_str); i++ {
		if token_str[i] == '.' {
			cnt++
		}
	}

	if cnt != 2 {
		return "Invalid JWT token format"
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(token_str, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv(env_entity)), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["id"].(string)
	} else {
		return err.Error()
	}
}

func ValidarToken(token_str string, env_entity string) bool {
	cnt := 0
	for i := 0; i < len(token_str); i++ {
		if token_str[i] == '.' {
			cnt++
		}
	}

	if cnt != 2 {
		return false
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(token_str, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv(env_entity)), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid && err == nil {
		return true
	} else {
		return false
	}
}
