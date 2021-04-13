package administrador_academico

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Elimina un administrador_academico
// @Description Elimina un administrador_academico con los datos entregados
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_administrador_academico     path    string     true        "UUID del administrador_academico a eliminar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/administradores-academicos/{uuid_administrador_academico} [delete]
func Delete(c *gin.Context) {
	id_administrador_academico := c.Params.ByName("id_administrador_academico")

	var admin models.AdministradorAcademico
	if err := repositories.GetAdministradorAcademico(&admin, id_administrador_academico); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	if err := repositories.DeleteAdministradorAcademico(&admin, id_administrador_academico); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, "OK")
}
