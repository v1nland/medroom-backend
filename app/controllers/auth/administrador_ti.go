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

// @Summary Autenticación de administrador ti
// @Description Ingresa usuario y contraseña para iniciar sesión
// @Tags 01 - Autenticación
// @Accept  json
// @Produce  json
// @Param   input_credentials     body    Request.LoginAdministradorTi     true        "Credenciales de acceso"
// @Success 200 {array} Swagger.AuthenticationSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/login [post]
func AutenticarAdministradorTi(c *gin.Context) {
	var administrador_ti models.AdministradorTi
	var login_message Request.LoginAdministradorTi
	var token_response Response.Authentication

	if err := c.ShouldBind(&login_message); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	f_input.LoginAdministradorTi(&login_message)

	err := repositories.AuthenticateAdministradorTi(&administrador_ti, login_message.Correo_electronico_administrador_ti, login_message.Hash_contrasena_administrador_ti)
	if err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	} else {
		encoder := jwt.New(jwt.SigningMethodHS256)
		claims := encoder.Claims.(jwt.MapClaims)

		claims["id"] = administrador_ti.Id
		claims["perfil"] = "administrador_ti"
		claims["correo_electronico_administrador_ti"] = administrador_ti.Correo_electronico_administrador_ti

		token, _ := encoder.SignedString([]byte(os.Getenv("SECRET_KEY_ADMINISTRADOR_TI")))
		token_response.Token = token

		api_helpers.RespondJSON(c, 200, token_response)
		return
	}
}
