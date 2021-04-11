package evaluacion

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Lista de evaluaciones rendidas por un estudiante
// @Description Lista todas los evaluaciones rendidas de un estudiante
// @Tags 02 - Estudiantes
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso"
// @Param   id_grupo     path    string     true        "Id del grupo"
// @Success 200 {array} Swagger.ListEvaluaciones "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /evaluadores/me/cursos/{id_curso}/grupos/{id_grupo}/estudiantes/{id_estudiante}/evaluaciones-rendidas [get]
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
	api_helpers.RespondJSON(c, 200, evaluaciones)
}
