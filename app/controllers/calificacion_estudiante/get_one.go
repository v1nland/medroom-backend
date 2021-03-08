package calificacion_estudiante

import (
	"errors"
	"medroom-backend/app/api_helpers"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"
	"medroom-backend/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Obtiene una calificación
// @Description Obtiene una calificación de un estudiante
// @Tags 02 - Estudiantes
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso"
// @Param   id_grupo     path    string     true        "Id del grupo"
// @Param   id_evaluacion     path    string     true        "Id de la evaluacion"
// @Success 200 {array} Swagger.ListEvaluacionesSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /estudiantes/me/cursos/{id_curso}/grupos/{id_grupo}/evaluaciones/{id_evaluacion}/calificacion [get]
func GetOneCalificacionEstudiante(c *gin.Context) {
	id_evaluacion := c.Params.ByName("id_evaluacion")
	id_estudiante := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	var calificacion_estudiante models.CalificacionEstudiante
	if err := repositories.GetOneCalificacionEstudianteByIdEvaluacion(&calificacion_estudiante, id_evaluacion, id_estudiante); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Calificación estudiante not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.ListEvaluacionesEstudiante(calificacion_estudiante))
	api_helpers.RespondJSON(c, 200, calificacion_estudiante)
}
