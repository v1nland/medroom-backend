package curso

import (
	"errors"
	"medroom-backend/Utils"
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Obtiene los cursos de un estudiante
// @Description Obtiene los cursos de un estudiante según su token
// @Tags 02 - Estudiantes
// @Accept  json
// @Produce  json
// @Success 200 {object} Swagger.GetCursosEstudianteSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /estudiantes/me/cursos [get]
func GetCursosEstudiante(c *gin.Context) {
	id_estudiante := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	var cursos []models.Curso
	if err := repositories.GetCursosEstudiante(&cursos, id_estudiante); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Curso not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	// api_helpers.RespondJSON(c, 200, Output.GetCursoEstudianteOutput(cursos))
	api_helpers.RespondJSON(c, 200, cursos)
}
