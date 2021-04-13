package rol

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Elimina un rol
// @Description Elimina un rol con los datos entregados
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_rol     path    string     true        "Id del rol a eliminar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/roles/{id_rol} [delete]
func Delete(c *gin.Context) {
	id_rol := c.Params.ByName("id_rol")

	var rol models.Rol
	if err := repositories.GetOneRol(&rol, id_rol); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	if err := repositories.DeleteRol(&rol, id_rol); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, rol)
}
