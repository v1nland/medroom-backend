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

// @Summary Modifica los cursos de un evaluador
// @Description Modifica los cursos de un evaluador con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "ID del curso a modificar"
// @Param   uuid_evaluador     path    string     true        "UUID del evaluador a asociar"
// @Success 200 {object} Swagger.AddEvaluadorToCursoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/cursos/{id_curso}/evaluadores/{uuid_evaluador} [put]
func AddEvaluadorToCurso(c *gin.Context) {
	id_curso := c.Params.ByName("id")
	id_evaluador := c.Params.ByName("id_evaluador")

	var evaluador models.Evaluador
	if err := repositories.GetOneEvaluador(&evaluador, id_evaluador); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Evaluador not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	var curso models.Curso
	if err := repositories.GetOneCurso(&curso, id_curso); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Curso not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	// search grupo "sin grupo" in curso
	found, index := utils.SearchIndexGrupoBySigla(curso.Grupos_curso, "SG")
	if found {
		curso.Grupos_curso[index].Evaluadores_grupo = append(curso.Grupos_curso[index].Evaluadores_grupo, evaluador)
	} else {
		api_helpers.RespondJSON(c, 200, "Grupo 'SIN GRUPO' not found")
		return
	}

	if err := repositories.PutOneCurso(&curso, id_curso); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Curso not updated")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	api_helpers.RespondJSON(c, 200, curso)
}
