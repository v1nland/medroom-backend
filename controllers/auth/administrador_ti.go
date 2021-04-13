package auth

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type authAdministradorTiRequest struct {
	Correo_electronico_administrador_ti string `json:"correo_electronico_administrador_ti"`
	Hash_contrasena_administrador_ti    string `json:"hash_contrasena_administrador_ti"`
}

func authAdministradorTiRequestParse(msg *authAdministradorTiRequest) {
	msg.Hash_contrasena_administrador_ti = strings.TrimSpace(msg.Hash_contrasena_administrador_ti)
	msg.Correo_electronico_administrador_ti = strings.ToUpper(msg.Correo_electronico_administrador_ti)
}

type authAdministradorTiResponse struct {
	Token string `json:"token"`
}

// @Summary Autenticación de administrador ti
// @Description Ingresa usuario y contraseña para iniciar sesión
// @Tags Autenticación
// @Accept  json
// @Produce  json
// @Param   input_credentials     body    authAdministradorTiRequest     true        "Credenciales de acceso"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/login [post]
func AuthenticateAdministradorTi(c *gin.Context) {
	var administrador_ti models.AdministradorTi

	var login_message authAdministradorTiRequest
	var login_response authAdministradorTiResponse

	if err := c.ShouldBind(&login_message); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	authAdministradorTiRequestParse(&login_message)

	if err := repositories.AuthenticateAdministradorTi(&administrador_ti, login_message.Correo_electronico_administrador_ti, login_message.Hash_contrasena_administrador_ti); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	} else {
		encoder := jwt.New(jwt.SigningMethodHS256)
		claims := encoder.Claims.(jwt.MapClaims)

		claims["id"] = administrador_ti.Id
		claims["perfil"] = "administrador_ti"
		claims["correo_electronico_administrador_ti"] = administrador_ti.Correo_electronico_administrador_ti

		token, _ := encoder.SignedString([]byte(os.Getenv("SECRET_KEY_ADMINISTRADOR_TI")))
		login_response.Token = token

		api_helpers.RespondJson(c, 200, login_response)
		return
	}
}
