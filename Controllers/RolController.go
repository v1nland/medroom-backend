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

// @Summary Lista de roles
// @Description Lista todos los roles
// @Tags 00 - Rutas públicas
// @Accept  json
// @Produce  json
// @Success 200 {array} Swagger.ListRolesSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /roles [get]
func ListRoles(c *gin.Context) {
	// model container
	var container []Models.Rol

	// query
	err := Repositories.GetAllRoles(&container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.ListRolesOutput(container))
}

// @Summary Obtiene un rol
// @Description Obtiene un rol según su Id
// @Tags 00 - Rutas públicas
// @Accept  json
// @Produce  json
// @Param   id_rol     path    string     true        "Id del rol a buscar"
// @Success 200 {object} Swagger.GetOneRolSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /roles/{id_rol} [get]
func GetOneRol(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.Rol

	// query
	err := Repositories.GetOneRol(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.GetOneRolOutput(container))
}

// @Summary Agrega un nuevo rol
// @Description Genera un nuevo rol con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_rol     body    Request.AddNewRolPayload     true        "Rol a agregar"
// @Success 200 {object} Swagger.AddNewRolSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/roles [post]
func AddNewRol(c *gin.Context) {
	// input container
	var container Request.AddNewRolPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	Input.AddNewRolInput(&container)

	// generate model entity
	model_container := Models.Rol{
		Nombre_rol: container.Nombre_rol,
	}

	// query
	err := Repositories.AddNewRol(&model_container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.AddNewRolOutput(model_container))
}

// @Summary Modifica un rol
// @Description Modifica un rol con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_rol     path    string     true        "Id del rol a modificar"
// @Param   input_actualiza_rol     body    Request.PutOneRolPayload     true        "Rol a modificar"
// @Success 200 {object} Swagger.PutOneRolSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/roles/{id_rol} [put]
func PutOneRol(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// input container
	var container Request.PutOneRolPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	Input.PutOneRolInput(&container)

	// generate model entity
	var model_container Models.Rol

	// get query
	err := Repositories.GetOneRol(&model_container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// replace data in model entity
	model_container = Models.Rol{
		Id:         model_container.Id,
		Nombre_rol: Utils.CheckUpdatedString(container.Nombre_rol, model_container.Nombre_rol),
	}

	// put query
	err = Repositories.PutOneRol(&model_container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.PutOneRolOutput(model_container))
}

// @Summary Elimina un rol
// @Description Elimina un rol con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_rol     path    string     true        "Id del rol a eliminar"
// @Success 200 {object} Swagger.DeleteRolSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/roles/{id_rol} [delete]
func DeleteRol(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.Rol

	// get query
	err := Repositories.GetOneRol(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// query
	err = Repositories.DeleteRol(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.DeleteRolOutput(container))
}
