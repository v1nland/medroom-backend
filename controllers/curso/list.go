package curso

import (
	"errors"
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Lista de cursos
// @Description Lista todos los cursos
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Success 200 {array} Swagger.ListCursosSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/cursos [get]
func ListCursos(c *gin.Context) {
	var cursos []models.Curso
	if err := repositories.GetAllCursos(&cursos); err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			api_helpers.RespondJSON(c, 200, cursos)
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	// api_helpers.RespondJSON(c, 200, Output.ListCursosOutput(container))
	api_helpers.RespondJSON(c, 200, cursos)
}
