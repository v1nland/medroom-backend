package curso

import (
	"errors"
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Obtiene los cursos de un administrador academico
// @Description Obtiene los cursos de un administrador academico según su token
// @Tags 05 - Administración Académica
// @Accept  json
// @Produce  json
// @Success 200 {object} Swagger.GetCursosAdministradorAcademicoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-academica/me/cursos [get]
func GetCursosAdministradorAcademico(c *gin.Context) {
	// params
	id_evaluador := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_ACADEMICO")

	var cursos []models.Curso
	if err := repositories.GetCursosAdministradorAcademico(&cursos, id_evaluador); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Curso not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.GetCursoEvaluador(cursos))
	api_helpers.RespondJSON(c, 200, cursos)
}
