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
// @Tags 04 - Administración Académica
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "ID del curso a modificar"
// @Param   id_grupo     path    string     true        "ID del grupo a modificar"
// @Param   uuid_estudiante     path    string     true        "UUID del estudiante a asociar"
// @Success 200 {object} Swagger.RemoveEstudianteFromGrupoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-academica/cursos/{id_curso}/grupos/{id_grupo}/estudiantes/{uuid_estudiante} [delete]
func RemoveEstudianteFromGrupo(c *gin.Context) {
	id_periodo := c.Params.ByName("id_periodo")
	sigla_curso := c.Params.ByName("id")
	sigla_grupo := c.Params.ByName("id_grupo")
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

	repositories.DeleteEstudianteGrupo(sigla_grupo, sigla_curso, id_periodo, id_estudiante)

	found, index := utils.SearchIndexGrupoBySigla(curso.Grupos_curso, "SG")
	if found {
		var grupo_sg models.Grupo
		if err := repositories.GetOneGrupo(&grupo_sg, curso.Grupos_curso[index].Sigla_curso, curso.Grupos_curso[index].Id_periodo_curso, curso.Grupos_curso[index].Sigla_grupo); err != nil {
			api_helpers.RespondError(c, 500, err.Error())
			return
		}

		config.DB.Model(&grupo_sg).Association("Estudiantes_grupo").Append([]models.Estudiante{estudiante})
	}

	api_helpers.RespondJSON(c, 200, grupo)
}
