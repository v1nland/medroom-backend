package auth

import (
	"medroom-backend/api_helpers"
	"medroom-backend/messages/Response"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type authEvaluadorRequest struct {
	Correo_electronico_evaluador string `json:"correo_electronico_evaluador"`
	Hash_contrasena_evaluador    string `json:"hash_contrasena_evaluador"`
}

func authEvaluadorRequestParse(msg *authEvaluadorRequest) {
	msg.Hash_contrasena_evaluador = strings.TrimSpace(msg.Hash_contrasena_evaluador)
	msg.Correo_electronico_evaluador = strings.ToUpper(msg.Correo_electronico_evaluador)
}

// @Summary Autenticación de evaluador
// @Description Ingresa usuario y contraseña para iniciar sesión
// @Tags 01 - Autenticación
// @Accept  json
// @Produce  json
// @Param   input_credentials     body    Request.LoginEvaluador     true        "Credenciales de acceso"
// @Success 200 {array} Swagger.AuthenticationSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /evaluadores/login [post]
func AuthenticateEvaluador(c *gin.Context) {
	var evaluador models.Evaluador

	var login_message authEvaluadorRequest
	var token_response Response.Authentication

	if err := c.ShouldBind(&login_message); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	authEvaluadorRequestParse(&login_message)

	if err := repositories.AuthenticateEvaluador(&evaluador, login_message.Correo_electronico_evaluador, login_message.Hash_contrasena_evaluador); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	} else {
		encoder := jwt.New(jwt.SigningMethodHS256)
		claims := encoder.Claims.(jwt.MapClaims)

		claims["id"] = evaluador.Id
		claims["perfil"] = "evaluador"
		claims["correo_electronico_evaluador"] = evaluador.Correo_electronico_evaluador

		token, _ := encoder.SignedString([]byte(os.Getenv("SECRET_KEY_EVALUADOR")))
		token_response.Token = token

		api_helpers.RespondJSON(c, 200, token_response)
		return
	}
}
