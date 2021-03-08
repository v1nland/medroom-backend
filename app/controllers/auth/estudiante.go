package auth

import (
	"medroom-backend/app/api_helpers"
	"medroom-backend/app/formats/f_input"
	"medroom-backend/app/messages/Request"
	"medroom-backend/app/messages/Response"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// @Summary Autenticación de estudiante
// @Description Ingresa usuario y contraseña para iniciar sesión
// @Tags 01 - Autenticación
// @Accept  json
// @Produce  json
// @Param   input_credentials     body    Request.LoginEstudiante     true        "Credenciales de acceso"
// @Success 200 {array} Swagger.AuthenticationSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /estudiantes/login [post]
func AutenticarEstudiante(c *gin.Context) {
	var estudiante models.Estudiante
	var login_message Request.LoginEstudiante
	var token_response Response.Authentication

	if err := c.ShouldBind(&login_message); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	f_input.LoginEstudiante(&login_message)

	err := repositories.AuthenticateEstudiante(&estudiante, login_message.Correo_electronico_estudiante, login_message.Hash_contrasena_estudiante)
	if err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	} else {
		encoder := jwt.New(jwt.SigningMethodHS256)
		claims := encoder.Claims.(jwt.MapClaims)

		claims["id"] = estudiante.Id
		claims["perfil"] = "estudiante"
		claims["correo_electronico_estudiante"] = estudiante.Correo_electronico_estudiante

		token, _ := encoder.SignedString([]byte(os.Getenv("SECRET_KEY_ESTUDIANTE")))
		token_response.Token = token

		api_helpers.RespondJSON(c, 200, token_response)
		return
	}
}
