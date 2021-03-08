package Utils

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

/*
	*
	*  FUNCIÓN DecodificarToken
	*
    *  ESTA FUNCIÓN RETORNA id_tienda EMBEBIDO EN EL TOKEN
	*  EN CASO DE TOKEN INVALIDO, RETORNA UN ERROR
	*
    *
*/
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

/*
	*
	*  FUNCIÓN ValidarTokenEstudiante
	*
    *  ESTA FUNCIÓN RETORNA true SI ES QUE UN TOKEN RECIBIDO ES VÁLIDO
	*  Y RETORNA false SI ES QUE ES UN TOKEN INVÁLIDO.
	*  ESTÁ DISEÑADA PARA USO INTERNO DE LOS PATHS.
    *
*/
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
