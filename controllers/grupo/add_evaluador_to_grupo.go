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
// @Tags 04 - Administración Académica
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "ID del curso a modificar"
// @Param   id_grupo     path    string     true        "ID del grupo a modificar"
// @Param   uuid_evaluador     path    string     true        "UUID del evaluador a asociar"
// @Success 200 {object} Swagger.AddEvaluadorToGrupoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-academica/cursos/{id_curso}/grupos/{id_grupo}/evaluadores/{uuid_evaluador} [put]
func AddEvaluadorToGrupo(c *gin.Context) {
	id_periodo := c.Params.ByName("id_periodo")
	sigla_curso := c.Params.ByName("id")
	sigla_grupo := c.Params.ByName("id_grupo")
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

	config.DB.Model(&grupo).Association("Evaluadores_grupo").Append([]models.Evaluador{evaluador})

	found, index := utils.SearchIndexGrupoBySigla(curso.Grupos_curso, "SG")
	if found {
		repositories.DeleteEvaluadorGrupo(curso.Grupos_curso[index].Sigla_grupo, curso.Grupos_curso[index].Sigla_curso, curso.Grupos_curso[index].Id_periodo_curso, id_evaluador)
	}

	api_helpers.RespondJSON(c, 200, grupo)
}
