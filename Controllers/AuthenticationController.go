package Controllers

import (
	"medroom-backend/ApiHelpers"
	"medroom-backend/Formats/Input"
	"medroom-backend/Messages/Request"
	"medroom-backend/Messages/Response"
	"medroom-backend/Models"
	"medroom-backend/Repositories"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// @Summary Autenticación de estudiante
// @Description Ingresa usuario y contraseña para iniciar sesión
// @Tags 01 - Autenticación
// @Accept  json
// @Produce  json
// @Param   input_credentials     body    Request.LoginEstudiantePayload     true        "Credenciales de acceso"
// @Success 200 {array} Swagger.AuthenticationSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes/login [post]
func AutenticarEstudiante(c *gin.Context) {
	var estudiante Models.Estudiante
	var login_message Request.LoginEstudiantePayload
	var token_response Response.Authentication

	if err := c.ShouldBind(&login_message); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	Input.FormatLoginEstudianteMessage(&login_message)

	err := Repositories.AuthenticateEstudiante(&estudiante, login_message.Correo_electronico_estudiante, login_message.Hash_contrasena_estudiante)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	} else {
		encoder := jwt.New(jwt.SigningMethodHS256)
		claims := encoder.Claims.(jwt.MapClaims)

		claims["id"] = estudiante.Id
		claims["perfil"] = "estudiante"
		claims["correo_electronico_estudiante"] = estudiante.Correo_electronico_estudiante

		token, _ := encoder.SignedString([]byte(os.Getenv("SECRET_KEY_ESTUDIANTE")))
		token_response.Token = token

		ApiHelpers.RespondJSON(c, 200, token_response)
		return
	}
}

// @Summary Autenticación de evaluador
// @Description Ingresa usuario y contraseña para iniciar sesión
// @Tags 01 - Autenticación
// @Accept  json
// @Produce  json
// @Param   input_credentials     body    Request.LoginEvaluadorPayload     true        "Credenciales de acceso"
// @Success 200 {array} Swagger.AuthenticationSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluadores/login [post]
func AutenticarEvaluador(c *gin.Context) {
	var evaluador Models.Evaluador
	var login_message Request.LoginEvaluadorPayload
	var token_response Response.Authentication

	if err := c.ShouldBind(&login_message); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	Input.FormatLoginEvaluadorMessage(&login_message)

	err := Repositories.AuthenticateEvaluador(&evaluador, login_message.Correo_electronico_evaluador, login_message.Hash_contrasena_evaluador)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	} else {
		encoder := jwt.New(jwt.SigningMethodHS256)
		claims := encoder.Claims.(jwt.MapClaims)

		claims["id"] = evaluador.Id
		claims["perfil"] = "evaluador"
		claims["correo_electronico_evaluador"] = evaluador.Correo_electronico_evaluador

		token, _ := encoder.SignedString([]byte(os.Getenv("SECRET_KEY_EVALUADOR")))
		token_response.Token = token

		ApiHelpers.RespondJSON(c, 200, token_response)
		return
	}
}

// @Summary Autenticación de administrador académico
// @Description Ingresa usuario y contraseña para iniciar sesión
// @Tags 01 - Autenticación
// @Accept  json
// @Produce  json
// @Param   input_credentials     body    Request.LoginAdministradorAcademicoPayload     true        "Credenciales de acceso"
// @Success 200 {array} Swagger.AuthenticationSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-academica/login [post]
func AutenticarAdministradorAcademico(c *gin.Context) {
	var administrador_academico Models.AdministradorAcademico
	var login_message Request.LoginAdministradorAcademicoPayload
	var token_response Response.Authentication

	if err := c.ShouldBind(&login_message); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	Input.FormatLoginAdministradorAcademicoMessage(&login_message)

	err := Repositories.AuthenticateAdministradorAcademico(&administrador_academico, login_message.Correo_electronico_administrador_academico, login_message.Hash_contrasena_administrador_academico)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	} else {
		encoder := jwt.New(jwt.SigningMethodHS256)
		claims := encoder.Claims.(jwt.MapClaims)

		claims["id"] = administrador_academico.Id
		claims["perfil"] = "administrador_academico"
		claims["correo_electronico_administrador_academico"] = administrador_academico.Correo_electronico_administrador_academico

		token, _ := encoder.SignedString([]byte(os.Getenv("SECRET_KEY_ADMINISTRADOR_ACADEMICO")))
		token_response.Token = token

		ApiHelpers.RespondJSON(c, 200, token_response)
		return
	}
}

// @Summary Autenticación de administrador ti
// @Description Ingresa usuario y contraseña para iniciar sesión
// @Tags 01 - Autenticación
// @Accept  json
// @Produce  json
// @Param   input_credentials     body    Request.LoginAdministradorTiPayload     true        "Credenciales de acceso"
// @Success 200 {array} Swagger.AuthenticationSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/login [post]
func AutenticarAdministradorTi(c *gin.Context) {
	var administrador_ti Models.AdministradorTi
	var login_message Request.LoginAdministradorTiPayload
	var token_response Response.Authentication

	if err := c.ShouldBind(&login_message); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	Input.FormatLoginAdministradorTiMessage(&login_message)

	err := Repositories.AuthenticateAdministradorTi(&administrador_ti, login_message.Correo_electronico_administrador_ti, login_message.Hash_contrasena_administrador_ti)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	} else {
		encoder := jwt.New(jwt.SigningMethodHS256)
		claims := encoder.Claims.(jwt.MapClaims)

		claims["id"] = administrador_ti.Id
		claims["perfil"] = "administrador_ti"
		claims["correo_electronico_administrador_ti"] = administrador_ti.Correo_electronico_administrador_ti

		token, _ := encoder.SignedString([]byte(os.Getenv("SECRET_KEY_ADMINISTRADOR_TI")))
		token_response.Token = token

		ApiHelpers.RespondJSON(c, 200, token_response)
		return
	}
}
