package grupo

import (
	"medroom-backend/api_helpers"
	"medroom-backend/config"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Modifica los grupos de un evaluador
// @Description Modifica los grupos de un evaluador con los datos entregados
// @Tags Administración Académica
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo"
// @Param   sigla_curso     path    string     true        "Sigla del curso"
// @Param   sigla_grupo     path    string     true        "Sigla del grupo"
// @Param   uuid_evaluador     path    string     true        "UUID del evaluador a asociar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-academica/cursos/{id_periodo}/{sigla_curso}/grupos/{sigla_grupo}/evaluadors/{uuid_evaluador} [delete]
func RemoveEvaluadorFromGrupo(c *gin.Context) {
	id_periodo := c.Params.ByName("id_periodo")
	sigla_curso := c.Params.ByName("sigla_curso")
	sigla_grupo := c.Params.ByName("sigla_grupo")
	id_evaluador := c.Params.ByName("id_evaluador")

	var curso models.Curso
	if err := repositories.GetOneCurso(&curso, sigla_curso, id_periodo); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	var grupo models.Grupo
	if err := repositories.GetOneGrupo(&grupo, sigla_curso, id_periodo, sigla_grupo); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	var evaluador models.Evaluador
	if err := repositories.GetOneEvaluador(&evaluador, id_evaluador); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	repositories.DeleteEvaluadorGrupo(sigla_grupo, sigla_curso, id_periodo, id_evaluador)

	found, index := utils.SearchIndexGrupoBySigla(curso.Grupos_curso, "SG")
	if found {
		var grupo_sg models.Grupo
		if err := repositories.GetOneGrupo(&grupo_sg, curso.Grupos_curso[index].Sigla_curso, curso.Grupos_curso[index].Id_periodo_curso, curso.Grupos_curso[index].Sigla_grupo); err != nil {
			api_helpers.RespondError(c, 500, "default")
			return
		}

		config.DB.Model(&grupo_sg).Association("Evaluadores_grupo").Append([]models.Evaluador{evaluador})
	}

	api_helpers.RespondJson(c, 200, grupo)
}
