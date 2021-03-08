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

// @Summary Obtiene los cursos de un evaluador
// @Description Obtiene los cursos de un evaluador según su token
// @Tags 03 - Evaluadores
// @Accept  json
// @Produce  json
// @Success 200 {object} Swagger.GetCursosEvaluadorSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /evaluadores/me/cursos [get]
func GetCursosEvaluador(c *gin.Context) {
	// params
	id_evaluador := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")

	var cursos []models.Curso
	if err := repositories.GetCursosEvaluador(&cursos, id_evaluador); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Curso not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	// api_helpers.RespondJSON(c, 200, Output.GetCursoEvaluadorOutput(cursos))
	api_helpers.RespondJSON(c, 200, cursos)
}
