package Controllers

import (
	"medroom-backend/ApiHelpers"
	"medroom-backend/InputFormats"
	"medroom-backend/Models"
	"medroom-backend/OutputFormats"
	"medroom-backend/Repositories"
	"medroom-backend/RequestMessages"
	"medroom-backend/Utils"

	"github.com/gin-gonic/gin"
)

// @Summary Lista de administradores-academicos
// @Description Lista todos los administradores-academicos
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Success 200 {array} SwaggerMessages.ListAdministradoresAcademicosSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-academicos [get]
func ListAdministradoresAcademicos(c *gin.Context) {
	// model container
	var container []Models.AdministradorAcademico

	// query
	err := Repositories.GetAllAdministradoresAcademicos(&container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetAdministradoresAcademicosOutput(container))
}

// @Summary Obtiene un administrador_academico
// @Description Obtiene un administrador_academico según su UUID
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_administrador_academico     path    string     true        "UUID del administrador_academico a buscar"
// @Success 200 {object} SwaggerMessages.GetOneAdministradorAcademicoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-academicos/{uuid_administrador_academico} [get]
func GetOneAdministradorAcademico(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.AdministradorAcademico

	// query
	err := Repositories.GetOneAdministradorAcademico(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetOneAdministradorAcademicoOutput(container))
}

// @Summary Agrega un nuevo administrador_academico
// @Description Genera un nuevo administrador_academico con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_administrador_academico     body    RequestMessages.AddNewAdministradorAcademicoPayload     true        "AdministradorAcademico a agregar"
// @Success 200 {object} SwaggerMessages.AddNewAdministradorAcademicoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-academicos [post]
func AddNewAdministradorAcademico(c *gin.Context) {
	// input container
	var container RequestMessages.AddNewAdministradorAcademicoPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.AddNewAdministradorAcademicoInput(&container)

	// generate model entity
	model_container := Models.AdministradorAcademico{
		Id_rol:                                     container.Id_rol,
		Rut_administrador_academico:                container.Rut_administrador_academico,
		Nombres_administrador_academico:            container.Nombres_administrador_academico,
		Apellidos_administrador_academico:          container.Apellidos_administrador_academico,
		Hash_contrasena_administrador_academico:    container.Hash_contrasena_administrador_academico,
		Correo_electronico_administrador_academico: container.Correo_electronico_administrador_academico,
		Telefono_fijo_administrador_academico:      container.Telefono_fijo_administrador_academico,
		Telefono_celular_administrador_academico:   container.Telefono_celular_administrador_academico,
	}

	// query
	err := Repositories.AddNewAdministradorAcademico(&model_container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.AddNewAdministradorAcademicoOutput(model_container))
}

// @Summary Modifica un administrador_academico
// @Description Modifica un administrador_academico con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_administrador_academico     path    string     true        "UUID del administrador_academico a modificar"
// @Param   input_actualiza_administrador_academico     body    RequestMessages.PutOneAdministradorAcademicoPayload     true        "AdministradorAcademico a modificar"
// @Success 200 {object} SwaggerMessages.PutOneAdministradorAcademicoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-academicos/{uuid_administrador_academico} [put]
func PutOneAdministradorAcademico(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// input container
	var container RequestMessages.PutOneAdministradorAcademicoPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.PutOneAdministradorAcademicoInput(&container)

	// generate model entity
	var model_container Models.AdministradorAcademico

	// get query
	err := Repositories.GetOneAdministradorAcademico(&model_container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// replace data in model entity
	model_container = Models.AdministradorAcademico{
		ID:                                      model_container.ID,
		Id_rol:                                  Utils.CheckUpdatedInt(container.Id_rol, model_container.Id_rol),
		Rut_administrador_academico:             Utils.CheckUpdatedString(container.Rut_administrador_academico, model_container.Rut_administrador_academico),
		Nombres_administrador_academico:         Utils.CheckUpdatedString(container.Nombres_administrador_academico, model_container.Nombres_administrador_academico),
		Apellidos_administrador_academico:       Utils.CheckUpdatedString(container.Apellidos_administrador_academico, model_container.Apellidos_administrador_academico),
		Hash_contrasena_administrador_academico: Utils.CheckUpdatedString(container.Hash_contrasena_administrador_academico, model_container.Hash_contrasena_administrador_academico),
		Correo_electronico_administrador_academico: Utils.CheckUpdatedString(container.Correo_electronico_administrador_academico, model_container.Correo_electronico_administrador_academico),
		Telefono_fijo_administrador_academico:      Utils.CheckUpdatedString(container.Telefono_fijo_administrador_academico, model_container.Telefono_fijo_administrador_academico),
		Telefono_celular_administrador_academico:   Utils.CheckUpdatedString(container.Telefono_celular_administrador_academico, model_container.Telefono_celular_administrador_academico),
	}

	// update foreign entity
	err = Repositories.GetOneRol(&model_container.Rol_administrador_academico, Utils.ConvertIntToString(model_container.Id_rol))
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// put query
	err = Repositories.PutOneAdministradorAcademico(&model_container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.PutOneAdministradorAcademicoOutput(model_container))
}

// @Summary Elimina un administrador_academico
// @Description Elimina un administrador_academico con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_administrador_academico     path    string     true        "UUID del administrador_academico a eliminar"
// @Success 200 {object} SwaggerMessages.DeleteAdministradorAcademicoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-academicos/{uuid_administrador_academico} [delete]
func DeleteAdministradorAcademico(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.AdministradorAcademico

	// get query
	err := Repositories.GetOneAdministradorAcademico(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// query
	err = Repositories.DeleteAdministradorAcademico(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.DeleteAdministradorAcademicoOutput(container))
}

// @Summary Obtiene el perfil del administrador academico
// @Description Obtiene el perfil del administrador academico según su token
// @Tags 04 - Administración Academica
// @Accept  json
// @Produce  json
// @Success 200 {object} SwaggerMessages.GetMyAdministradorAcademicoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-academica/me [get]
func GetMyAdministradorAcademico(c *gin.Context) {
	// params
	id_administrador_academico := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_ACADEMICO")

	// model container
	var container Models.AdministradorAcademico

	// query
	err := Repositories.GetOneAdministradorAcademico(&container, id_administrador_academico)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetMyAdministradorAcademicoOutput(container))
}

// @Summary Modifica mi perfil
// @Description Modifica el perfil de un administrador_academico con los datos entregados
// @Tags 04 - Administración Academica
// @Accept  json
// @Produce  json
// @Param   input_actualiza_administrador_academico     body    RequestMessages.PutMyAdministradorAcademicoPayload     true        "AdministradorAcademico a modificar"
// @Success 200 {object} SwaggerMessages.PutMyAdministradorAcademicoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-academica/me [put]
func PutMyAdministradorAcademico(c *gin.Context) {
	// params
	id_administrador_academico := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_ACADEMICO")

	// input container
	var container RequestMessages.PutMyAdministradorAcademicoPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.PutMyAdministradorAcademicoInput(&container)

	// generate model entity
	var model_container Models.AdministradorAcademico

	// get query
	err := Repositories.GetOneAdministradorAcademico(&model_container, id_administrador_academico)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// replace data in model entity
	model_container = Models.AdministradorAcademico{
		ID:                                      model_container.ID,
		Id_rol:                                  Utils.CheckUpdatedInt(container.Id_rol, model_container.Id_rol),
		Rut_administrador_academico:             Utils.CheckUpdatedString(container.Rut_administrador_academico, model_container.Rut_administrador_academico),
		Nombres_administrador_academico:         Utils.CheckUpdatedString(container.Nombres_administrador_academico, model_container.Nombres_administrador_academico),
		Apellidos_administrador_academico:       Utils.CheckUpdatedString(container.Apellidos_administrador_academico, model_container.Apellidos_administrador_academico),
		Hash_contrasena_administrador_academico: Utils.CheckUpdatedString(container.Hash_contrasena_administrador_academico, model_container.Hash_contrasena_administrador_academico),
		Correo_electronico_administrador_academico: Utils.CheckUpdatedString(container.Correo_electronico_administrador_academico, model_container.Correo_electronico_administrador_academico),
		Telefono_fijo_administrador_academico:      Utils.CheckUpdatedString(container.Telefono_fijo_administrador_academico, model_container.Telefono_fijo_administrador_academico),
		Telefono_celular_administrador_academico:   Utils.CheckUpdatedString(container.Telefono_celular_administrador_academico, model_container.Telefono_celular_administrador_academico),
	}

	// update foreign entity
	err = Repositories.GetOneRol(&model_container.Rol_administrador_academico, Utils.ConvertIntToString(model_container.Id_rol))
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// put query
	err = Repositories.PutOneAdministradorAcademico(&model_container, id_administrador_academico)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.PutMyAdministradorAcademicoOutput(model_container))
}
