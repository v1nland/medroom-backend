package Controllers

import (
	"medroom-backend/ApiHelpers"
	"medroom-backend/Formats/Input"
	"medroom-backend/Formats/Output"
	"medroom-backend/Messages/Request"
	"medroom-backend/Models"
	"medroom-backend/Repositories"
	"medroom-backend/Utils"

	"github.com/gin-gonic/gin"
)

// @Summary Lista de administradores-ti
// @Description Lista todos los administradores-ti
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Success 200 {array} Swagger.ListAdministradoresTiSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-ti [get]
func ListAdministradoresTi(c *gin.Context) {
	var administradores_ti []Models.AdministradorTi
	if err := Repositories.GetAllAdministradoresTi(&administradores_ti); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	ApiHelpers.RespondJSON(c, 200, Output.ListAdministradoresTiOutput(administradores_ti))
}

// @Summary Obtiene un administrador_ti
// @Description Obtiene un administrador_ti según su UUID
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_administrador_ti     path    string     true        "UUID del administrador_ti a buscar"
// @Success 200 {object} Swagger.GetOneAdministradorTiSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-ti/{uuid_administrador_ti} [get]
func GetOneAdministradorTi(c *gin.Context) {
	id := c.Params.ByName("id")

	var administrador_ti Models.AdministradorTi
	if err := Repositories.GetOneAdministradorTi(&administrador_ti, id); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	ApiHelpers.RespondJSON(c, 200, Output.GetOneAdministradorTiOutput(administrador_ti))
}

// @Summary Agrega un nuevo administrador_ti
// @Description Genera un nuevo administrador_ti con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_administrador_ti     body    Request.AddNewAdministradorTi     true        "AdministradorTi a agregar"
// @Success 200 {object} Swagger.AddNewAdministradorTiSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-ti [post]
func AddNewAdministradorTi(c *gin.Context) {
	var input Request.AddNewAdministradorTi
	if err := c.ShouldBind(&input); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	Input.AddNewAdministradorTi(&input)

	model_container := Models.AdministradorTi{
		Id_rol:                              *input.Id_rol,
		Rut_administrador_ti:                *input.Rut_administrador_ti,
		Nombres_administrador_ti:            *input.Nombres_administrador_ti,
		Apellidos_administrador_ti:          *input.Apellidos_administrador_ti,
		Hash_contrasena_administrador_ti:    *input.Hash_contrasena_administrador_ti,
		Correo_electronico_administrador_ti: *input.Correo_electronico_administrador_ti,
		Telefono_fijo_administrador_ti:      *input.Telefono_fijo_administrador_ti,
		Telefono_celular_administrador_ti:   *input.Telefono_celular_administrador_ti,
	}

	if err := Repositories.AddNewAdministradorTi(&model_container); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	ApiHelpers.RespondJSON(c, 200, Output.AddNewAdministradorTiOutput(model_container))
}

// @Summary Modifica un administrador_ti
// @Description Modifica un administrador_ti con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_administrador_ti     path    string     true        "UUID del administrador_ti a modificar"
// @Param   input_actualiza_administrador_ti     body    Request.PutOneAdministradorTi     true        "AdministradorTi a modificar"
// @Success 200 {object} Swagger.PutOneAdministradorTiSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-ti/{uuid_administrador_ti} [put]
func PutOneAdministradorTi(c *gin.Context) {
	id := c.Params.ByName("id")

	var input Request.PutOneAdministradorTi
	if err := c.ShouldBind(&input); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	Input.PutOneAdministradorTi(&input)

	var administrador_ti Models.AdministradorTi
	if err := Repositories.GetOneAdministradorTi(&administrador_ti, id); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	administrador_ti = Models.AdministradorTi{
		Id:                                  administrador_ti.Id,
		Id_rol:                              administrador_ti.Id_rol,
		Rut_administrador_ti:                Utils.CheckNullString(input.Rut_administrador_ti, administrador_ti.Rut_administrador_ti),
		Nombres_administrador_ti:            Utils.CheckNullString(input.Nombres_administrador_ti, administrador_ti.Nombres_administrador_ti),
		Apellidos_administrador_ti:          Utils.CheckNullString(input.Apellidos_administrador_ti, administrador_ti.Apellidos_administrador_ti),
		Hash_contrasena_administrador_ti:    Utils.CheckNullString(input.Hash_contrasena_administrador_ti, administrador_ti.Hash_contrasena_administrador_ti),
		Correo_electronico_administrador_ti: Utils.CheckNullString(input.Correo_electronico_administrador_ti, administrador_ti.Correo_electronico_administrador_ti),
		Telefono_fijo_administrador_ti:      Utils.CheckNullString(input.Telefono_fijo_administrador_ti, administrador_ti.Telefono_fijo_administrador_ti),
		Telefono_celular_administrador_ti:   Utils.CheckNullString(input.Telefono_celular_administrador_ti, administrador_ti.Telefono_celular_administrador_ti),
	}

	if err := Repositories.PutOneAdministradorTi(&administrador_ti, id); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	ApiHelpers.RespondJSON(c, 200, Output.PutOneAdministradorTiOutput(administrador_ti))
}

// @Summary Elimina un administrador_ti
// @Description Elimina un administrador_ti con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_administrador_ti     path    string     true        "UUID del administrador_ti a eliminar"
// @Success 200 {object} Swagger.DeleteAdministradorTiSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-ti/{uuid_administrador_ti} [delete]
func DeleteAdministradorTi(c *gin.Context) {
	id := c.Params.ByName("id")

	var container Models.AdministradorTi

	if err := Repositories.GetOneAdministradorTi(&container, id); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	if err := Repositories.DeleteAdministradorTi(&container, id); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	ApiHelpers.RespondJSON(c, 200, Output.DeleteAdministradorTiOutput(container))
}

// @Summary Obtiene el perfil del administrador ti
// @Description Obtiene el perfil del administrador ti según su token
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Success 200 {object} Swagger.GetMyAdministradorTiSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/me [get]
func GetMyAdministradorTi(c *gin.Context) {
	id := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_TI")

	var administrador_ti Models.AdministradorTi
	if err := Repositories.GetOneAdministradorTi(&administrador_ti, id); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	ApiHelpers.RespondJSON(c, 200, Output.GetMyAdministradorTiOutput(administrador_ti))
}

// @Summary Modifica mi perfil
// @Description Modifica el perfil de un administrador_ti con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_actualiza_administrador_ti     body    Request.PutMyAdministradorTi     true        "AdministradorTi a modificar"
// @Success 200 {object} Swagger.PutMyAdministradorTiSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/me [put]
func PutMyAdministradorTi(c *gin.Context) {
	id := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_TI")

	var input Request.PutMyAdministradorTi
	if err := c.ShouldBind(&input); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	Input.PutMyAdministradorTi(&input)

	var administrador_ti Models.AdministradorTi
	if err := Repositories.GetOneAdministradorTi(&administrador_ti, id); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// replace data in model entity
	administrador_ti = Models.AdministradorTi{
		Id:                                  administrador_ti.Id,
		Id_rol:                              administrador_ti.Id_rol,
		Rut_administrador_ti:                Utils.CheckNullString(input.Rut_administrador_ti, administrador_ti.Rut_administrador_ti),
		Nombres_administrador_ti:            Utils.CheckNullString(input.Nombres_administrador_ti, administrador_ti.Nombres_administrador_ti),
		Apellidos_administrador_ti:          Utils.CheckNullString(input.Apellidos_administrador_ti, administrador_ti.Apellidos_administrador_ti),
		Hash_contrasena_administrador_ti:    Utils.CheckNullString(input.Hash_contrasena_administrador_ti, administrador_ti.Hash_contrasena_administrador_ti),
		Correo_electronico_administrador_ti: Utils.CheckNullString(input.Correo_electronico_administrador_ti, administrador_ti.Correo_electronico_administrador_ti),
		Telefono_fijo_administrador_ti:      Utils.CheckNullString(input.Telefono_fijo_administrador_ti, administrador_ti.Telefono_fijo_administrador_ti),
		Telefono_celular_administrador_ti:   Utils.CheckNullString(input.Telefono_celular_administrador_ti, administrador_ti.Telefono_celular_administrador_ti),
	}

	if err := Repositories.PutOneAdministradorTi(&administrador_ti, id); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	ApiHelpers.RespondJSON(c, 200, Output.PutMyAdministradorTiOutput(administrador_ti))
}
