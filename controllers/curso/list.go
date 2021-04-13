package curso

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Lista de cursos
// @Description Lista todos los cursos
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/cursos [get]
func ListCursos(c *gin.Context) {
	var cursos []models.Curso
	if err := repositories.GetAllCursos(&cursos); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, cursos)
}
