package estudiante

import (
	"errors"
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Elimina un estudiante
// @Description Elimina un estudiante con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_estudiante     path    string     true        "UUID del estudiante a eliminar"
// @Success 200 {object} Swagger.DeleteEstudianteSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/estudiantes/{uuid_estudiante} [delete]
func DeleteEstudiante(c *gin.Context) {
	id := c.Params.ByName("id")

	var estudiante models.Estudiante
	if err := repositories.GetOneEstudiante(&estudiante, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	if err := repositories.DeleteEstudiante(&estudiante, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	// api_helpers.RespondJSON(c, 200, Output.DeleteEstudianteOutput(container))
	api_helpers.RespondJSON(c, 200, estudiante)
}
