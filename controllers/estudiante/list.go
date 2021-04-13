package estudiante

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Lista de estudiantes
// @Description Lista todos los estudiantes existentes
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/estudiantes [get]
func List(c *gin.Context) {
	var estudiantes []models.Estudiante
	if err := repositories.GetAllEstudiantes(&estudiantes); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.ListEstudiantes(container))
	api_helpers.RespondJson(c, 200, estudiantes)
}
