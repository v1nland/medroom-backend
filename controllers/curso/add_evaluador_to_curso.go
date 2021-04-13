package curso

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Modifica los cursos de un evaluador
// @Description Modifica los cursos de un evaluador con los datos entregados
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo"
// @Param   sigla_curso     path    string     true        "Sigla del curso a modificar"
// @Param   uuid_evaluador     path    string     true        "UUID del evaluador a asociar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/cursos/{id_periodo}/{sigla_curso}/evaluadores/{uuid_evaluador} [put]
func AddEvaluadorToCurso(c *gin.Context) {
	id_periodo := c.Params.ByName("id_periodo")
	sigla_curso := c.Params.ByName("sigla_curso")
	id_evaluador := c.Params.ByName("id_evaluador")

	var evaluador models.Evaluador
	if err := repositories.GetOneEvaluador(&evaluador, id_evaluador); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	var curso models.Curso
	if err := repositories.GetOneCurso(&curso, sigla_curso, id_periodo); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	// search grupo "sin grupo" in curso
	found, index := utils.SearchIndexGrupoBySigla(curso.Grupos_curso, "SG")
	if found {
		curso.Grupos_curso[index].Evaluadores_grupo = append(curso.Grupos_curso[index].Evaluadores_grupo, evaluador)
	} else {
		api_helpers.RespondJson(c, 200, "Grupo 'SIN GRUPO' not found")
		return
	}

	if err := repositories.PutOneCurso(&curso, sigla_curso, id_periodo); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, curso)
}
