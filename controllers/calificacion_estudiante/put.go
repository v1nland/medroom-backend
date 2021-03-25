package calificacion_estudiante

import (
	"medroom-backend/api_helpers"
	"medroom-backend/formats/f_input"
	"medroom-backend/messages/Request"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Cambia la calificación de un estudiante
// @Description Cambia una calificación para un estudiante
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
// @Router /evaluadores/me/cursos/{id_curso}/grupos/{id_grupo}/estudiantes/{id_estudiante}/evaluaciones/{id_evaluacion}/calificacion [put]
func PutOneCalificacionEstudiante(c *gin.Context) {
	// id_evaluador := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")
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
	// if len(input.Puntajes_calificacion_estudiante) != 7 {
	// 	api_helpers.RespondError(c, 400, "Deben venir las 7 competencias")
	// 	return
	// }

	f_input.AddNewCalificacionEstudiante(&input)

	var calificacion_actual models.CalificacionEstudiante
	if err := repositories.GetOneCalificacionEstudianteByIdEvaluacion(&calificacion_actual, id_evaluacion, id_estudiante); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	calificacion_actual.Observacion_calificacion_calificacion_estudiante = *input.Observacion_calificacion_calificacion_estudiante
	calificacion_actual.Valoracion_general_calificacion_estudiante = *input.Valoracion_general_calificacion_estudiante

	for _, val := range input.Puntajes_calificacion_estudiante {
		index := searchPuntajeByIdCompetencia(calificacion_actual.Puntajes_calificacion_estudiante, *val.Id_competencia)

		if index != -1 {
			calificacion_actual.Puntajes_calificacion_estudiante[index].Calificacion_puntaje = *val.Calificacion_puntaje
		}
	}

	if err := repositories.PutOneCalificacionEstudiante(&calificacion_actual, utils.ConvertIntToString(calificacion_actual.Id)); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.ListEvaluacionesEstudiante(model))
	api_helpers.RespondJSON(c, 200, calificacion_actual.Id)
}

func searchPuntajeByIdCompetencia(puntajes []models.Puntaje, id_competencia string) int {
	for i, v := range puntajes {
		if v.Id_competencia == id_competencia {
			return i
		}
	}

	return -1
}
