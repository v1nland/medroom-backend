package administrador_ti

import (
	"medroom-backend/api_helpers"
	"medroom-backend/formats/Output"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Elimina un administrador_ti
// @Description Elimina un administrador_ti con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_administrador_ti     path    string     true        "UUID del administrador_ti a eliminar"
// @Success 200 {object} Swagger.DeleteAdministradorTiSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-ti/{uuid_administrador_ti} [delete]
func DeleteAdministradorTi(c *gin.Context) {
	id := c.Params.ByName("id")

	var container models.AdministradorTi

	if err := repositories.GetOneAdministradorTi(&container, id); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	if err := repositories.DeleteAdministradorTi(&container, id); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	api_helpers.RespondJSON(c, 200, Output.DeleteAdministradorTiOutput(container))
}
