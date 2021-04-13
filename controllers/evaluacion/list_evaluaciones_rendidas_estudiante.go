package evaluacion

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Lista de evaluaciones rendidas por un estudiante
// @Description Lista todas los evaluaciones rendidas de un estudiante
// @Tags Estudiantes
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo"
// @Param   sigla_curso     path    string     true        "Sigla del curso"
// @Param   sigla_grupo     path    string     true        "Sigla del grupo"
// @Param   uuid_estudiante     path    string     true        "UUID del estudiante a eliminar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /evaluadores/me/cursos/{id_periodo}/{sigla_curso}/grupos/{sigla_grupo}/estudiantes/{uuid_estudiante}/evaluaciones-rendidas [get]
func ListEvaluacionesRendidasEstudiante(c *gin.Context) {
	// id_evaluador := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")
	// id_grupo := c.Params.ByName("id_grupo")
	// id_curso := c.Params.ByName("id_curso")
	id_estudiante := c.Params.ByName("id_estudiante")

	var evaluaciones []models.Evaluacion
	if err := repositories.GetEvaluacionesRendidasEstudiante(&evaluaciones, id_estudiante); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.ListEvaluacionesEstudiante(grupo.Evaluaciones_grupo))
	api_helpers.RespondJson(c, 200, evaluaciones)
}
