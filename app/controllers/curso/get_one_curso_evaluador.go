package curso

import (
	"errors"
	"medroom-backend/app/api_helpers"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"
	"medroom-backend/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Obtiene un curso de un evaluador
// @Description Obtiene un curso de un evaluador según su token
// @Tags 03 - Evaluadores
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso a buscar"
// @Success 200 {object} Swagger.GetOneCursoEvaluadorSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /evaluadores/me/cursos/{id_curso} [get]
func GetOneCursoEvaluador(c *gin.Context) {
	id_curso := c.Params.ByName("id_curso")
	id_evaluador := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")

	var curso models.Curso
	if err := repositories.GetOneCursoEvaluador(&curso, id_curso, id_evaluador); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Curso not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.GetCursoEstudiante(cursos))
	api_helpers.RespondJSON(c, 200, curso)
}
