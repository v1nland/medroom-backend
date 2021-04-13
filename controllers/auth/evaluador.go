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

type authEvaluadorRequest struct {
	Correo_electronico_evaluador string `json:"correo_electronico_evaluador"`
	Hash_contrasena_evaluador    string `json:"hash_contrasena_evaluador"`
}

func authEvaluadorRequestParse(msg *authEvaluadorRequest) {
	msg.Hash_contrasena_evaluador = strings.TrimSpace(msg.Hash_contrasena_evaluador)
	msg.Correo_electronico_evaluador = strings.ToUpper(msg.Correo_electronico_evaluador)
}

type authEvaluadorResponse struct {
	Token string `json:"token"`
}

// @Summary Autenticación de evaluador
// @Description Ingresa usuario y contraseña para iniciar sesión
// @Tags Autenticación
// @Accept  json
// @Produce  json
// @Param   input_credentials     body    authEvaluadorRequest     true        "Credenciales de acceso"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /evaluadores/login [post]
func AuthenticateEvaluador(c *gin.Context) {
	var evaluador models.Evaluador

	var login_message authEvaluadorRequest
	var login_response authEvaluadorResponse

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
		login_response.Token = token

		api_helpers.RespondJson(c, 200, login_response)
		return
	}
}
