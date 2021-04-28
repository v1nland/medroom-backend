package calificacion_estudiante

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type putPuntaje struct {
	Id_competencia       *string `json:"id_competencia"`
	Calificacion_puntaje *int    `json:"calificacion_puntaje"`
}

type putRequest struct {
	Observacion_calificacion_calificacion_estudiante *string      `json:"observacion_calificacion_calificacion_estudiante"`
	Valoracion_general_calificacion_estudiante       *int         `json:"valoracion_general_calificacion_estudiante"`
	Puntajes_calificacion_estudiante                 []putPuntaje `json:"puntajes_calificacion_estudiante"`
}

func putRequestParse(u *putRequest) {
	if u.Observacion_calificacion_calificacion_estudiante != nil {
		*u.Observacion_calificacion_calificacion_estudiante = strings.TrimSpace(*u.Observacion_calificacion_calificacion_estudiante)
		*u.Observacion_calificacion_calificacion_estudiante = strings.ToUpper(*u.Observacion_calificacion_calificacion_estudiante)
		*u.Observacion_calificacion_calificacion_estudiante = utils.RemoveAccents(*u.Observacion_calificacion_calificacion_estudiante)
	}
}

// @Summary Cambia la calificación de un estudiante
// @Description Cambia una calificación para un estudiante
// @Tags Evaluadores
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo"
// @Param   sigla_curso     path    string     true        "Sigla del curso"
// @Param   sigla_grupo     path    string     true        "Sigla del grupo"
// @Param   id_estudiante     path    string     true        "Id del estudiante"
// @Param   id_evaluacion     path    string     true        "Id de la evaluacion"
// @Param   input_calificacion_estudiante     body    putRequest     true        "CalificacionEstudiante a agregar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /evaluadores/me/cursos/{id_periodo}/{sigla_curso}/grupos/{sigla_grupo}/estudiantes/{id_estudiante}/evaluaciones/{id_evaluacion}/calificacion [put]
func Put(c *gin.Context) {
	// id_evaluador := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")
	id_estudiante := c.Params.ByName("id_estudiante")
	id_evaluacion := c.Params.ByName("id_evaluacion")
	// id_periodo := c.Params.ByName("id_periodo")
	// sigla_curso := c.Params.ByName("sigla_curso")
	// sigla_grupo := c.Params.ByName("sigla_grupo")

	var input putRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	// no es necesario validar todas las competencias
	// ya que pueden actualizarse algunas y otras no
	// if len(input.Puntajes_calificacion_estudiante) != 6 {
	// 	api_helpers.RespondError(c, 400, "Deben venir las 6 competencias")
	// 	return
	// }

	putRequestParse(&input)

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

	if err := repositories.PutOneCalificacionEstudiante(&calificacion_actual, utils.IntToString(calificacion_actual.Id)); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.ListEvaluacionesEstudiante(model))
	api_helpers.RespondJson(c, 200, calificacion_actual.Id)
}

func searchPuntajeByIdCompetencia(puntajes []models.Puntaje, id_competencia string) int {
	for i, v := range puntajes {
		if v.Id_competencia == id_competencia {
			return i
		}
	}

	return -1
}
