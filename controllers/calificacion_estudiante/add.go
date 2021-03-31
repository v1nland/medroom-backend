package calificacion_estudiante

import (
	"medroom-backend/api_helpers"
	"medroom-backend/formats/f_input"
	"medroom-backend/messages/Request"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Califica un estudiante
// @Description Genera una calificación para un estudiante
// @Tags 03 - Evaluadores
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso"
// @Param   id_grupo     path    string     true        "Id del grupo"
// @Param   id_estudiante     path    string     true        "Id del estudiante"
// @Param   id_evaluacion     path    string     true        "Id de la evaluacion"
// @Param   input_calificacion_estudiante     body    Request.AddNewCalificacionEstudiante     true        "CalificacionEstudiante a agregar"
// @Success 200 {array} Swagger.AddNewCalificacionEstudianteSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /evaluadores/me/cursos/{id_curso}/grupos/{id_grupo}/estudiantes/{id_estudiante}/evaluaciones/{id_evaluacion}/calificacion [post]
func AddNewCalificacionEstudiante(c *gin.Context) {
	id_evaluador := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")
	// id_curso := c.Params.ByName("id_curso")
	// id_grupo := c.Params.ByName("id_grupo")
	id_estudiante := c.Params.ByName("id_estudiante")
	id_evaluacion := c.Params.ByName("id_evaluacion")

	var input Request.AddNewCalificacionEstudiante
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	// TODO: validar que vengan las 7 competencias necesarias
	if len(input.Puntajes_calificacion_estudiante) != 6 {
		api_helpers.RespondError(c, 400, "deben venir las 6 competencias")
		return
	}

	f_input.AddNewCalificacionEstudiante(&input)

	// parameter for calificacion estudiante constructor
	var puntajes_calificacion_estudiante []models.Puntaje
	for i := 0; i < len(input.Puntajes_calificacion_estudiante); i++ {
		puntajes_calificacion_estudiante = append(puntajes_calificacion_estudiante, models.Puntaje{
			Id_competencia:       *input.Puntajes_calificacion_estudiante[i].Id_competencia,
			Calificacion_puntaje: *input.Puntajes_calificacion_estudiante[i].Calificacion_puntaje,
			Feedback_puntaje:     *input.Puntajes_calificacion_estudiante[i].Feedback_puntaje,
		})
	}

	model := models.CalificacionEstudiante{
		Id_evaluador:                     uuid.MustParse(id_evaluador),
		Id_periodo:                       *input.Id_periodo,
		Id_evaluacion:                    utils.ConvertStringToInt(id_evaluacion),
		Puntajes_calificacion_estudiante: puntajes_calificacion_estudiante,
		Id_estudiante:                    uuid.MustParse(id_estudiante),
		Nombre_calificacion_estudiante:   *input.Nombre_calificacion_estudiante,
		Observacion_calificacion_calificacion_estudiante: *input.Observacion_calificacion_calificacion_estudiante,
		Valoracion_general_calificacion_estudiante:       *input.Valoracion_general_calificacion_estudiante,
	}

	if err := repositories.AddNewCalificacionEstudiante(&model); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.ListEvaluacionesEstudiante(model))
	api_helpers.RespondJSON(c, 200, model)
}
