package rol

import (
	"medroom-backend/api_helpers"
	"medroom-backend/formats/f_output"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Elimina un rol
// @Description Elimina un rol con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_rol     path    string     true        "Id del rol a eliminar"
// @Success 200 {object} Swagger.DeleteRolSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/roles/{id_rol} [delete]
func DeleteRol(c *gin.Context) {
	id := c.Params.ByName("id")

	var rol models.Rol
	if err := repositories.GetOneRol(&rol, id); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	if err := repositories.DeleteRol(&rol, id); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	api_helpers.RespondJSON(c, 200, f_output.DeleteRol(rol))
}
