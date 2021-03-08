package estudiante

import (
	"errors"
	"medroom-backend/app/api_helpers"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Lista de estudiantes
// @Description Lista todos los estudiantes existentes
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Success 200 {array} Swagger.ListEstudiantesSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/estudiantes [get]
func ListEstudiantes(c *gin.Context) {
	var estudiantes []models.Estudiante
	if err := repositories.GetAllEstudiantes(&estudiantes); err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			api_helpers.RespondJSON(c, 200, estudiantes)
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.ListEstudiantes(container))
	api_helpers.RespondJSON(c, 200, estudiantes)
}
