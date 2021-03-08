package auth

import (
	"medroom-backend/app/Messages/Request"
	"medroom-backend/app/Messages/Response"
	"medroom-backend/app/api_helpers"
	"medroom-backend/app/formats/f_input"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// @Summary Autenticación de evaluador
// @Description Ingresa usuario y contraseña para iniciar sesión
// @Tags 01 - Autenticación
// @Accept  json
// @Produce  json
// @Param   input_credentials     body    Request.LoginEvaluador     true        "Credenciales de acceso"
// @Success 200 {array} Swagger.AuthenticationSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /evaluadores/login [post]
func AutenticarEvaluador(c *gin.Context) {
	var evaluador models.Evaluador
	var login_message Request.LoginEvaluador
	var token_response Response.Authentication

	if err := c.ShouldBind(&login_message); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	f_input.LoginEvaluador(&login_message)

	err := repositories.AuthenticateEvaluador(&evaluador, login_message.Correo_electronico_evaluador, login_message.Hash_contrasena_evaluador)
	if err != nil {
		api_helpers.RespondError(c, 500, "default")
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
