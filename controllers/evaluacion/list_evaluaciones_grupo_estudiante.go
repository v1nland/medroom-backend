package evaluacion

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Lista de evaluaciones de un grupo
// @Description Lista todas los evaluaciones disponibles de un estudiante de un grupo
// @Tags Estudiantes
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo"
// @Param   sigla_curso     path    string     true        "Sigla del curso"
// @Param   sigla_grupo     path    string     true        "Sigla del grupo"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /estudiantes/me/cursos/{id_periodo}/{sigla_curso}/grupos/{sigla_grupo}/evaluaciones [get]
func ListEvaluacionesGrupoEstudiante(c *gin.Context) {
	id_periodo := c.Params.ByName("id_periodo")
	sigla_grupo := c.Params.ByName("sigla_grupo")
	sigla_curso := c.Params.ByName("sigla_curso")
	id_estudiante := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	var grupo models.Grupo
	if err := repositories.GetOneGrupoEstudiante(&grupo, sigla_grupo, sigla_curso, id_periodo, id_estudiante); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.ListEvaluacionesEstudiante(grupo.Evaluaciones_grupo))
	api_helpers.RespondJson(c, 200, grupo.Evaluaciones_grupo)
}
