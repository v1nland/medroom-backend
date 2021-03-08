package administrador_academico

import (
	"medroom-backend/api_helpers"
	"medroom-backend/formats/f_output"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Elimina un administrador_academico
// @Description Elimina un administrador_academico con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_administrador_academico     path    string     true        "UUID del administrador_academico a eliminar"
// @Success 200 {object} Swagger.DeleteAdministradorAcademicoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-academicos/{uuid_administrador_academico} [delete]
func DeleteAdministradorAcademico(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container models.AdministradorAcademico

	// get query
	err := repositories.GetOneAdministradorAcademico(&container, id)
	if err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	// query
	err = repositories.DeleteAdministradorAcademico(&container, id)
	if err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	// output
	api_helpers.RespondJSON(c, 200, f_output.DeleteAdministradorAcademico(container))
}
