package rol

import (
	"medroom-backend/Messages/Request"
	"medroom-backend/Utils"
	"medroom-backend/api_helpers"
	"medroom-backend/formats/Output"
	"medroom-backend/formats/f_input"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Modifica un rol
// @Description Modifica un rol con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_rol     path    string     true        "Id del rol a modificar"
// @Param   input_actualiza_rol     body    Request.PutOneRol     true        "Rol a modificar"
// @Success 200 {object} Swagger.PutOneRolSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/roles/{id_rol} [put]
func PutOneRol(c *gin.Context) {
	id := c.Params.ByName("id")

	var input Request.PutOneRol
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	f_input.PutOneRol(&input)

	var rol models.Rol
	if err := repositories.GetOneRol(&rol, id); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	rol = models.Rol{
		Id:         rol.Id,
		Nombre_rol: Utils.CheckNullString(input.Nombre_rol, rol.Nombre_rol),
	}

	if err := repositories.PutOneRol(&rol, id); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	api_helpers.RespondJSON(c, 200, Output.PutOneRolOutput(rol))
}
