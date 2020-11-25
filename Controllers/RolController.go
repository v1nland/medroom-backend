package Controllers

import (
	"github.com/gin-gonic/gin"
	"medroom-backend/ApiHelpers"
	"medroom-backend/InputFormats"
	"medroom-backend/Models"
	"medroom-backend/OutputFormats"
	"medroom-backend/Repositories"
	"medroom-backend/RequestMessages"
	"medroom-backend/Utils"
)

/*
	*
	*  FUNCIÓN ListRol
	*
    *
	*
	*
    *
*/

// @Summary Lista de rols
// @Description Lista todos los rols
// @Accept  json
// @Produce  json
// @Success 200 {array} ResponseMessages.ListRolsResponse "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /rols [get]
func ListRols(c *gin.Context) {
	// model container
	var container []Models.Rol

	// query
	err := Repositories.GetAllRols(&container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetRolsOutput(container))
}

/*
	*
	*  FUNCIÓN GetOneRol
	*
    *
	*
	*
    *
*/

// @Summary Obtiene un rol
// @Description Obtiene un rol según su UUID
// @Accept  json
// @Produce  json
// @Param   uuid_rol     path    string     true        "UUID del rol a buscar"
// @Success 200 {object} ResponseMessages.GetOneRolResponse "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /rols/{uuid_rol} [get]
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
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetOneRolOutput(container))
}

/*
	*
	*  FUNCIÓN AddNewRol
	*
    *
	*
	*
    *
*/

// @Summary Agrega un nuevo rol
// @Description Genera un nuevo rol con los datos entregados
// @Accept  json
// @Produce  json
// @Param   input_rol     body    RequestMessages.AddNewRolPayload     true        "Rol a agregar"
// @Success 200 {object} ResponseMessages.AddNewRolResponse "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /rols [post]
func AddNewRol(c *gin.Context) {
	// input container
	var container RequestMessages.AddNewRolPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.AddNewRolInput(&container)

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
	ApiHelpers.RespondJSON(c, 200, OutputFormats.AddNewRolOutput(model_container))
}

/*
	*
	*  FUNCIÓN PutOneRol
	*
    *
	*
	*
    *
*/

// @Summary Modifica un rol
// @Description Modifica un rol con los datos entregados
// @Accept  json
// @Produce  json
// @Param   uuid_rol     path    string     true        "UUID del rol a modificar"
// @Param   input_actualiza_rol     body    RequestMessages.PutOneRolPayload     true        "Rol a modificar"
// @Success 200 {object} ResponseMessages.PutOneRolResponse "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /rols/{uuid_rol} [put]
func PutOneRol(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// input container
	var container RequestMessages.PutOneRolPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.PutOneRolInput(&container)

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
		ID:         model_container.ID,
		Nombre_rol: Utils.CheckUpdatedString(container.Nombre_rol, model_container.Nombre_rol),
	}

	// put query
	err = Repositories.PutOneRol(&model_container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.PutOneRolOutput(model_container))
}

/*
	*
	*  FUNCIÓN DeleteRol
	*
    *
	*
	*
    *
*/

// @Summary Elimina un rol
// @Description Elimina un rol con los datos entregados
// @Accept  json
// @Produce  json
// @Param   uuid_rol     path    string     true        "UUID del rol a eliminar"
// @Success 200 {object} ResponseMessages.DeleteRolResponse "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /rols/{uuid_rol} [delete]
func DeleteRol(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.Rol

	// query
	err := Repositories.DeleteRol(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.DeleteRolOutput(container))
}
