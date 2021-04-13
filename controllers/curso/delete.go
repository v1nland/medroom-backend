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

// @Summary Elimina un curso
// @Description Elimina un curso con los datos entregados
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo"
// @Param   sigla_curso     path    string     true        "Sigla del curso a eliminar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/cursos/{id_periodo}/{sigla_curso} [delete]
func Delete(c *gin.Context) {
	id_periodo := c.Params.ByName("id_periodo")
	sigla_curso := c.Params.ByName("sigla_curso")

	var curso models.Curso
	if err := repositories.GetOneCurso(&curso, sigla_curso, id_periodo); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	// clear grupo associations
	for _, gp := range curso.Grupos_curso {
		var grupo models.Grupo
		if err := repositories.GetOneGrupo(&grupo, gp.Sigla_curso, gp.Id_periodo_curso, gp.Sigla_grupo); err != nil {
			api_helpers.RespondError(c, 500, err.Error())
			return
		}

		if err := repositories.ClearGrupo(grupo.Sigla_grupo, grupo.Sigla_curso, grupo.Id_periodo_curso); err != nil {
			api_helpers.RespondError(c, 500, "default")
			return
		}

		for _, evaluacion := range grupo.Evaluaciones_grupo {
			if err := repositories.DeleteEvaluacion(utils.IntToString(evaluacion.Id)); err != nil {
				api_helpers.RespondError(c, 500, "default")
				return
			}
		}
	}

	// eliminar grupos del curso
	for _, grupo := range curso.Grupos_curso {
		repositories.DeleteGrupo(&grupo)
	}

	// eliminar curso
	if err := repositories.ClearCursoAssociations(&curso, sigla_curso, id_periodo); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJson(c, 200, "Curso not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	api_helpers.RespondJson(c, 200, curso)
}
