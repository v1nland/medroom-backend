package Controllers

import (
	"medroom-backend/ApiHelpers"
	"medroom-backend/InputFormats"
	"medroom-backend/Models"
	"medroom-backend/Repositories"
	"medroom-backend/RequestMessages"
	"medroom-backend/ResponseMessages"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

/*
	*
	*  FUNCIÓN AutenticarEstudiante
	*
    *  ESTA FUNCIÓN RECIBE EL JSON BODY DE LA REQUEST, Y RETORNA
	*  UN TOKEN JWT CON LOS DATOS DE LA TIENDA CIFRADOS.
	*  SI NO SE ENCUENTRA EN LA DB LA TIENDA, RETORNA ERROR.
    *
*/
// @Summary Autenticación de estudiante
// @Description Ingresa usuario y contraseña para iniciar sesión
// @Tags Autenticación
// @Accept  json
// @Produce  json
// @Param   input_credentials     body    RequestMessages.LoginEstudiantePayload     true        "Credenciales de acceso"
// @Success 200 {array} SwaggerMessages.AuthenticationSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /login [post]
func AutenticarEstudiante(c *gin.Context) {
	var estudiante Models.Estudiante
	var login_message RequestMessages.LoginEstudiantePayload
	var token_response ResponseMessages.Authentication

	if err := c.ShouldBind(&login_message); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	InputFormats.FormatLoginEstudianteMessage(&login_message)

	err := Repositories.AuthEstudiante(&estudiante, login_message.Correo_electronico_estudiante, login_message.Hash_contrasena_estudiante)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	} else {
		encoder := jwt.New(jwt.SigningMethodHS256)
		claims := encoder.Claims.(jwt.MapClaims)

		claims["id"] = estudiante.ID
		claims["perfil"] = "estudiante"
		claims["correo_electronico_estudiante"] = estudiante.Correo_electronico_estudiante

		token, _ := encoder.SignedString([]byte(os.Getenv("SECRET_KEY_ESTUDIANTE")))
		token_response.Token = token

		ApiHelpers.RespondJSON(c, 200, token_response)
		return
	}
}

/*
	*
	*  FUNCIÓN AutenticarEvaluador
	*
    *  ESTA FUNCIÓN RECIBE EL JSON BODY DE LA REQUEST, Y RETORNA
	*  UN TOKEN JWT CON LOS DATOS DE LA TIENDA CIFRADOS.
	*  SI NO SE ENCUENTRA EN LA DB LA TIENDA, RETORNA ERROR.
    *
*/
// @Summary Autenticación de evaluador
// @Description Ingresa usuario y contraseña para iniciar sesión
// @Tags Autenticación
// @Accept  json
// @Produce  json
// @Param   input_credentials     body    RequestMessages.LoginEvaluadorPayload     true        "Credenciales de acceso"
// @Success 200 {array} SwaggerMessages.AuthenticationSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /login-evaluador [post]
func AutenticarEvaluador(c *gin.Context) {
	var evaluador Models.Evaluador
	var login_message RequestMessages.LoginEvaluadorPayload
	var token_response ResponseMessages.Authentication

	if err := c.ShouldBind(&login_message); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	InputFormats.FormatLoginEvaluadorMessage(&login_message)

	err := Repositories.AuthEvaluador(&evaluador, login_message.Correo_electronico_evaluador, login_message.Hash_contrasena_evaluador)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	} else {
		encoder := jwt.New(jwt.SigningMethodHS256)
		claims := encoder.Claims.(jwt.MapClaims)

		claims["id"] = evaluador.ID
		claims["perfil"] = "evaluador"
		claims["correo_electronico_evaluador"] = evaluador.Correo_electronico_evaluador

		token, _ := encoder.SignedString([]byte(os.Getenv("SECRET_KEY_EVALUADOR")))
		token_response.Token = token

		ApiHelpers.RespondJSON(c, 200, token_response)
		return
	}
}

/*
	*
	*  FUNCIÓN AutenticarAdministradorAcademico
	*
    *  ESTA FUNCIÓN RECIBE EL JSON BODY DE LA REQUEST, Y RETORNA
	*  UN TOKEN JWT CON LOS DATOS DE LA TIENDA CIFRADOS.
	*  SI NO SE ENCUENTRA EN LA DB LA TIENDA, RETORNA ERROR.
    *
*/
// @Summary Autenticación de administrador académico
// @Description Ingresa usuario y contraseña para iniciar sesión
// @Tags Autenticación
// @Accept  json
// @Produce  json
// @Param   input_credentials     body    RequestMessages.LoginAdministradorAcademicoPayload     true        "Credenciales de acceso"
// @Success 200 {array} SwaggerMessages.AuthenticationSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /login-admiministrador-academico [post]
func AutenticarAdministradorAcademico(c *gin.Context) {
	var administrador_academico Models.AdministradorAcademico
	var login_message RequestMessages.LoginAdministradorAcademicoPayload
	var token_response ResponseMessages.Authentication

	if err := c.ShouldBind(&login_message); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	InputFormats.FormatLoginAdministradorAcademicoMessage(&login_message)

	err := Repositories.AuthAdministradorAcademico(&administrador_academico, login_message.Correo_electronico_administrador_academico, login_message.Hash_contrasena_administrador_academico)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	} else {
		encoder := jwt.New(jwt.SigningMethodHS256)
		claims := encoder.Claims.(jwt.MapClaims)

		claims["id"] = administrador_academico.ID
		claims["perfil"] = "administrador_academico"
		claims["correo_electronico_administrador_academico"] = administrador_academico.Correo_electronico_administrador_academico

		token, _ := encoder.SignedString([]byte(os.Getenv("SECRET_KEY_ADMINISTRADOR_ACADEMICO")))
		token_response.Token = token

		ApiHelpers.RespondJSON(c, 200, token_response)
		return
	}
}

/*
	*
	*  FUNCIÓN AutenticarAdministradorTi
	*
    *  ESTA FUNCIÓN RECIBE EL JSON BODY DE LA REQUEST, Y RETORNA
	*  UN TOKEN JWT CON LOS DATOS DE LA TIENDA CIFRADOS.
	*  SI NO SE ENCUENTRA EN LA DB LA TIENDA, RETORNA ERROR.
    *
*/
// @Summary Autenticación de administrador ti
// @Description Ingresa usuario y contraseña para iniciar sesión
// @Tags Autenticación
// @Accept  json
// @Produce  json
// @Param   input_credentials     body    RequestMessages.LoginAdministradorTiPayload     true        "Credenciales de acceso"
// @Success 200 {array} SwaggerMessages.AuthenticationSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /login-administrador-ti [post]
func AutenticarAdministradorTi(c *gin.Context) {
	var administrador_ti Models.AdministradorTi
	var login_message RequestMessages.LoginAdministradorTiPayload
	var token_response ResponseMessages.Authentication

	if err := c.ShouldBind(&login_message); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	InputFormats.FormatLoginAdministradorTiMessage(&login_message)

	err := Repositories.AuthAdministradorTi(&administrador_ti, login_message.Correo_electronico_administrador_ti, login_message.Hash_contrasena_administrador_ti)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	} else {
		encoder := jwt.New(jwt.SigningMethodHS256)
		claims := encoder.Claims.(jwt.MapClaims)

		claims["id"] = administrador_ti.ID
		claims["perfil"] = "administrador_ti"
		claims["correo_electronico_administrador_ti"] = administrador_ti.Correo_electronico_administrador_ti

		token, _ := encoder.SignedString([]byte(os.Getenv("SECRET_KEY_ADMINISTRADOR_TI")))
		token_response.Token = token

		ApiHelpers.RespondJSON(c, 200, token_response)
		return
	}
}
