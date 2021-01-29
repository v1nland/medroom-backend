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
	var roles []Models.Rol

	if err := Repositories.GetAllRoles(&roles); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	ApiHelpers.RespondJSON(c, 200, Output.ListRolesOutput(roles))
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
	id := c.Params.ByName("id")

	var rol Models.Rol
	if err := Repositories.GetOneRol(&rol, id); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	ApiHelpers.RespondJSON(c, 200, Output.GetOneRolOutput(rol))
}

// @Summary Agrega un nuevo rol
// @Description Genera un nuevo rol con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_rol     body    Request.AddNewRol     true        "Rol a agregar"
// @Success 200 {object} Swagger.AddNewRolSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/roles [post]
func AddNewRol(c *gin.Context) {
	var input Request.AddNewRol

	if err := c.ShouldBind(&input); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	Input.AddNewRol(&input)

	rol := Models.Rol{
		Nombre_rol: *input.Nombre_rol,
	}

	if err := Repositories.AddNewRol(&rol); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	ApiHelpers.RespondJSON(c, 200, Output.AddNewRolOutput(rol))
}

// @Summary Modifica un rol
// @Description Modifica un rol con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_rol     path    string     true        "Id del rol a modificar"
// @Param   input_actualiza_rol     body    Request.PutOneRol     true        "Rol a modificar"
// @Success 200 {object} Swagger.PutOneRolSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/roles/{id_rol} [put]
func PutOneRol(c *gin.Context) {
	id := c.Params.ByName("id")

	var input Request.PutOneRol
	if err := c.ShouldBind(&input); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	Input.PutOneRol(&input)

	var rol Models.Rol
	if err := Repositories.GetOneRol(&rol, id); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	rol = Models.Rol{
		Id:         rol.Id,
		Nombre_rol: Utils.CheckNullString(input.Nombre_rol, rol.Nombre_rol),
	}

	if err := Repositories.PutOneRol(&rol, id); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	ApiHelpers.RespondJSON(c, 200, Output.PutOneRolOutput(rol))
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
	id := c.Params.ByName("id")

	var rol Models.Rol
	if err := Repositories.GetOneRol(&rol, id); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	if err := Repositories.DeleteRol(&rol, id); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	ApiHelpers.RespondJSON(c, 200, Output.DeleteRolOutput(rol))
}
