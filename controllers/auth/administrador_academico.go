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

type authAdministradorAcademicoRequest struct {
	Correo_electronico_administrador_academico string `json:"correo_electronico_administrador_academico"`
	Hash_contrasena_administrador_academico    string `json:"hash_contrasena_administrador_academico"`
}

func authAdministradorAcademicoRequestParse(msg *authAdministradorAcademicoRequest) {
	msg.Hash_contrasena_administrador_academico = strings.TrimSpace(msg.Hash_contrasena_administrador_academico)
	msg.Correo_electronico_administrador_academico = strings.ToUpper(msg.Correo_electronico_administrador_academico)
}

// @Summary Autenticación de administrador académico
// @Description Ingresa usuario y contraseña para iniciar sesión
// @Tags 01 - Autenticación
// @Accept  json
// @Produce  json
// @Param   input_credentials     body    Request.LoginAdministradorAcademico     true        "Credenciales de acceso"
// @Success 200 {array} Swagger.AuthenticationSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-academica/login [post]
func AuthenticateAdministradorAcademico(c *gin.Context) {
	var administrador_academico models.AdministradorAcademico

	var login_message authAdministradorAcademicoRequest
	var token_response Response.Authentication

	if err := c.ShouldBind(&login_message); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	authAdministradorAcademicoRequestParse(&login_message)

	if err := repositories.AuthenticateAdministradorAcademico(&administrador_academico, login_message.Correo_electronico_administrador_academico, login_message.Hash_contrasena_administrador_academico); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	} else {
		encoder := jwt.New(jwt.SigningMethodHS256)
		claims := encoder.Claims.(jwt.MapClaims)

		claims["id"] = administrador_academico.Id
		claims["perfil"] = "administrador_academico"
		claims["correo_electronico_administrador_academico"] = administrador_academico.Correo_electronico_administrador_academico

		token, _ := encoder.SignedString([]byte(os.Getenv("SECRET_KEY_ADMINISTRADOR_ACADEMICO")))
		token_response.Token = token

		api_helpers.RespondJSON(c, 200, token_response)
		return
	}
}
