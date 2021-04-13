package calificacion_estudiante

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Obtiene una calificación
// @Description Obtiene una calificación de un estudiante
// @Tags Estudiantes
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo"
// @Param   sigla_curso     path    string     true        "Sigla del curso"
// @Param   sigla_grupo     path    string     true        "Sigla del grupo"
// @Param   id_evaluacion     path    string     true        "Id de la evaluacion"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /estudiantes/me/cursos/{id_periodo}/{sigla_curso}/grupos/{sigla_grupo}/evaluaciones/{id_evaluacion}/calificacion [get]
func Get(c *gin.Context) {
	id_evaluacion := c.Params.ByName("id_evaluacion")
	id_estudiante := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	var calificacion_estudiante models.CalificacionEstudiante
	if err := repositories.GetOneCalificacionEstudianteByIdEvaluacion(&calificacion_estudiante, id_evaluacion, id_estudiante); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, calificacion_estudiante)
}
