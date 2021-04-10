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

type authEstudianteRequest struct {
	Correo_electronico_estudiante string `json:"correo_electronico_estudiante"`
	Hash_contrasena_estudiante    string `json:"hash_contrasena_estudiante"`
}

func authEstudianteRequestParse(msg *authEstudianteRequest) {
	msg.Hash_contrasena_estudiante = strings.TrimSpace(msg.Hash_contrasena_estudiante)
	msg.Correo_electronico_estudiante = strings.ToUpper(msg.Correo_electronico_estudiante)
}

// @Summary Autenticación de estudiante
// @Description Ingresa usuario y contraseña para iniciar sesión
// @Tags 01 - Autenticación
// @Accept  json
// @Produce  json
// @Param   input_credentials     body    Request.LoginEstudiante     true        "Credenciales de acceso"
// @Success 200 {array} Swagger.AuthenticationSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /estudiantes/login [post]
func AuthenticateEstudiante(c *gin.Context) {
	var estudiante models.Estudiante

	var login_message authEstudianteRequest
	var token_response Response.Authentication

	if err := c.ShouldBind(&login_message); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	authEstudianteRequestParse(&login_message)

	if err := repositories.AuthenticateEstudiante(&estudiante, login_message.Correo_electronico_estudiante, login_message.Hash_contrasena_estudiante); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
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
