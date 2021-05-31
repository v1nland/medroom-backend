package grupo

import (
	"medroom-backend/api_helpers"
	"medroom-backend/config"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Modifica los grupos de un estudiante
// @Description Modifica los grupos de un estudiante con los datos entregados
// @Tags Administración Académica
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo"
// @Param   sigla_curso     path    string     true        "Sigla del curso"
// @Param   sigla_grupo     path    string     true        "Sigla del grupo"
// @Param   uuid_estudiante     path    string     true        "UUID del estudiante a asociar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-academica/cursos/{id_periodo}/{sigla_curso}/grupos/{sigla_grupo}/estudiantes/{uuid_estudiante} [delete]
func RemoveEstudianteFromGrupo(c *gin.Context) {
	id_periodo := c.Params.ByName("id_periodo")
	sigla_curso := c.Params.ByName("sigla_curso")
	sigla_grupo := c.Params.ByName("sigla_grupo")
	id_estudiante := c.Params.ByName("id_estudiante")

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

	var estudiante models.Estudiante
	if err := repositories.GetOneEstudiante(&estudiante, id_estudiante); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	found, index := utils.SearchIndexGrupoBySigla(curso.Grupos_curso, "SG")
	if found {
		var grupo_sg models.Grupo
		if err := repositories.GetOneGrupo(&grupo_sg, curso.Grupos_curso[index].Sigla_curso, curso.Grupos_curso[index].Id_periodo_curso, curso.Grupos_curso[index].Sigla_grupo); err != nil {
			api_helpers.RespondError(c, 500, err.Error())
			return
		}

		config.DB.Model(&grupo_sg).Association("Estudiantes_grupo").Append([]models.Estudiante{estudiante})
	}

	repositories.DeleteEstudianteGrupo(sigla_grupo, sigla_curso, id_periodo, id_estudiante)

	api_helpers.RespondJson(c, 200, grupo)
}
