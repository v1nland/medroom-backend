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

// @Summary Lista de administradores-ti
// @Description Lista todos los administradores-ti
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Success 200 {array} SwaggerMessages.ListAdministradoresTiSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-ti [get]
func ListAdministradoresTi(c *gin.Context) {
	// model container
	var container []Models.AdministradorTi

	// query
	err := Repositories.GetAllAdministradoresTi(&container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetAdministradoresTiOutput(container))
}

// @Summary Obtiene un administrador_ti
// @Description Obtiene un administrador_ti según su UUID
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_administrador_ti     path    string     true        "UUID del administrador_ti a buscar"
// @Success 200 {object} SwaggerMessages.GetOneAdministradorTiSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-ti/{uuid_administrador_ti} [get]
func GetOneAdministradorTi(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.AdministradorTi

	// query
	err := Repositories.GetOneAdministradorTi(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetOneAdministradorTiOutput(container))
}

// @Summary Agrega un nuevo administrador_ti
// @Description Genera un nuevo administrador_ti con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_administrador_ti     body    RequestMessages.AddNewAdministradorTiPayload     true        "AdministradorTi a agregar"
// @Success 200 {object} SwaggerMessages.AddNewAdministradorTiSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-ti [post]
func AddNewAdministradorTi(c *gin.Context) {
	// input container
	var container RequestMessages.AddNewAdministradorTiPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.AddNewAdministradorTiInput(&container)

	// generate model entity
	model_container := Models.AdministradorTi{
		Id_rol:                              container.Id_rol,
		Rut_administrador_ti:                container.Rut_administrador_ti,
		Nombres_administrador_ti:            container.Nombres_administrador_ti,
		Apellidos_administrador_ti:          container.Apellidos_administrador_ti,
		Hash_contrasena_administrador_ti:    container.Hash_contrasena_administrador_ti,
		Correo_electronico_administrador_ti: container.Correo_electronico_administrador_ti,
		Telefono_fijo_administrador_ti:      container.Telefono_fijo_administrador_ti,
		Telefono_celular_administrador_ti:   container.Telefono_celular_administrador_ti,
	}

	// query
	err := Repositories.AddNewAdministradorTi(&model_container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.AddNewAdministradorTiOutput(model_container))
}

// @Summary Modifica un administrador_ti
// @Description Modifica un administrador_ti con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_administrador_ti     path    string     true        "UUID del administrador_ti a modificar"
// @Param   input_actualiza_administrador_ti     body    RequestMessages.PutOneAdministradorTiPayload     true        "AdministradorTi a modificar"
// @Success 200 {object} SwaggerMessages.PutOneAdministradorTiSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-ti/{uuid_administrador_ti} [put]
func PutOneAdministradorTi(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// input container
	var container RequestMessages.PutOneAdministradorTiPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.PutOneAdministradorTiInput(&container)

	// generate model entity
	var model_container Models.AdministradorTi

	// get query
	err := Repositories.GetOneAdministradorTi(&model_container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// replace data in model entity
	model_container = Models.AdministradorTi{
		ID:                                  model_container.ID,
		Id_rol:                              Utils.CheckUpdatedInt(container.Id_rol, model_container.Id_rol),
		Rut_administrador_ti:                Utils.CheckUpdatedString(container.Rut_administrador_ti, model_container.Rut_administrador_ti),
		Nombres_administrador_ti:            Utils.CheckUpdatedString(container.Nombres_administrador_ti, model_container.Nombres_administrador_ti),
		Apellidos_administrador_ti:          Utils.CheckUpdatedString(container.Apellidos_administrador_ti, model_container.Apellidos_administrador_ti),
		Hash_contrasena_administrador_ti:    Utils.CheckUpdatedString(container.Hash_contrasena_administrador_ti, model_container.Hash_contrasena_administrador_ti),
		Correo_electronico_administrador_ti: Utils.CheckUpdatedString(container.Correo_electronico_administrador_ti, model_container.Correo_electronico_administrador_ti),
		Telefono_fijo_administrador_ti:      Utils.CheckUpdatedString(container.Telefono_fijo_administrador_ti, model_container.Telefono_fijo_administrador_ti),
		Telefono_celular_administrador_ti:   Utils.CheckUpdatedString(container.Telefono_celular_administrador_ti, model_container.Telefono_celular_administrador_ti),
	}

	// update foreign entity
	err = Repositories.GetOneRol(&model_container.Rol_administrador_ti, Utils.ConvertIntToString(model_container.Id_rol))
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// put query
	err = Repositories.PutOneAdministradorTi(&model_container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.PutOneAdministradorTiOutput(model_container))
}

// @Summary Elimina un administrador_ti
// @Description Elimina un administrador_ti con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_administrador_ti     path    string     true        "UUID del administrador_ti a eliminar"
// @Success 200 {object} SwaggerMessages.DeleteAdministradorTiSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-ti/{uuid_administrador_ti} [delete]
func DeleteAdministradorTi(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.AdministradorTi

	// get query
	err := Repositories.GetOneAdministradorTi(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// query
	err = Repositories.DeleteAdministradorTi(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.DeleteAdministradorTiOutput(container))
}

// @Summary Obtiene el perfil del administrador ti
// @Description Obtiene el perfil del administrador ti según su token
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Success 200 {object} SwaggerMessages.GetMyAdministradorTiSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/me [get]
func GetMyAdministradorTi(c *gin.Context) {
	// params
	id_administrador_ti := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_TI")

	// model container
	var container Models.AdministradorTi

	// query
	err := Repositories.GetOneAdministradorTi(&container, id_administrador_ti)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetMyAdministradorTiOutput(container))
}

// @Summary Modifica mi perfil
// @Description Modifica el perfil de un administrador_ti con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_actualiza_administrador_ti     body    RequestMessages.PutMyAdministradorTiPayload     true        "AdministradorTi a modificar"
// @Success 200 {object} SwaggerMessages.PutMyAdministradorTiSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/me [put]
func PutMyAdministradorTi(c *gin.Context) {
	// params
	id_administrador_ti := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_TI")

	// input container
	var container RequestMessages.PutMyAdministradorTiPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.PutMyAdministradorTiInput(&container)

	// generate model entity
	var model_container Models.AdministradorTi

	// get query
	err := Repositories.GetOneAdministradorTi(&model_container, id_administrador_ti)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// replace data in model entity
	model_container = Models.AdministradorTi{
		ID:                                  model_container.ID,
		Id_rol:                              Utils.CheckUpdatedInt(container.Id_rol, model_container.Id_rol),
		Rut_administrador_ti:                Utils.CheckUpdatedString(container.Rut_administrador_ti, model_container.Rut_administrador_ti),
		Nombres_administrador_ti:            Utils.CheckUpdatedString(container.Nombres_administrador_ti, model_container.Nombres_administrador_ti),
		Apellidos_administrador_ti:          Utils.CheckUpdatedString(container.Apellidos_administrador_ti, model_container.Apellidos_administrador_ti),
		Hash_contrasena_administrador_ti:    Utils.CheckUpdatedString(container.Hash_contrasena_administrador_ti, model_container.Hash_contrasena_administrador_ti),
		Correo_electronico_administrador_ti: Utils.CheckUpdatedString(container.Correo_electronico_administrador_ti, model_container.Correo_electronico_administrador_ti),
		Telefono_fijo_administrador_ti:      Utils.CheckUpdatedString(container.Telefono_fijo_administrador_ti, model_container.Telefono_fijo_administrador_ti),
		Telefono_celular_administrador_ti:   Utils.CheckUpdatedString(container.Telefono_celular_administrador_ti, model_container.Telefono_celular_administrador_ti),
	}

	// update foreign entity
	err = Repositories.GetOneRol(&model_container.Rol_administrador_ti, Utils.ConvertIntToString(model_container.Id_rol))
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// put query
	err = Repositories.PutOneAdministradorTi(&model_container, id_administrador_ti)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.PutMyAdministradorTiOutput(model_container))
}
