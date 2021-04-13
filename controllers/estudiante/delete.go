package estudiante

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Elimina un estudiante
// @Description Elimina un estudiante con los datos entregados
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_estudiante     path    string     true        "UUID del estudiante a eliminar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/estudiantes/{uuid_estudiante} [delete]
func Delete(c *gin.Context) {
	id_estudiante := c.Params.ByName("id_estudiante")

	var estudiante models.Estudiante
	if err := repositories.GetOneEstudiante(&estudiante, id_estudiante); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	if err := repositories.DeleteEstudiante(&estudiante, id_estudiante); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.DeleteEstudiante(container))
	api_helpers.RespondJson(c, 200, estudiante)
}
