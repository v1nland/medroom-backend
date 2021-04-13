package administrador_ti

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Elimina un administrador_ti
// @Description Elimina un administrador_ti con los datos entregados
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_administrador_ti     path    string     true        "UUID del administrador_ti a eliminar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/administradores-ti/{uuid_administrador_ti} [delete]
func Delete(c *gin.Context) {
	id_administrador_ti := c.Params.ByName("id_administrador_ti")

	var admin models.AdministradorTi
	if err := repositories.GetOneAdministradorTi(&admin, id_administrador_ti); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	if err := repositories.DeleteAdministradorTi(&admin, id_administrador_ti); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, admin)
}
