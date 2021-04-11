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
	Feedback_puntaje     *string `json:"feedback_puntaje"`
}

type putRequest struct {
	Id_periodo                                           *int         `json:"id_periodo"`
	Nombre_calificacion_estudiante                       *string      `json:"nombre_calificacion_estudiante"`
	Entorno_clinico_calificacion_estudiante              *string      `json:"entorno_clinico_calificacion_estudiante"`
	Paciente_calificacion_estudiante                     *string      `json:"paciente_calificacion_estudiante"`
	Asunto_principal_consulta_calificacion_estudiante    *string      `json:"asunto_principal_consulta_calificacion_estudiante"`
	Complejidad_caso_calificacion_estudiante             *string      `json:"complejidad_caso_calificacion_estudiante"`
	Numero_observaciones_previas_calificacion_estudiante *string      `json:"numero_observaciones_previas_calificacion_estudiante"`
	Categoria_observador_calificacion_estudiante         *string      `json:"categoria_observador_calificacion_estudiante"`
	Observacion_calificacion_calificacion_estudiante     *string      `json:"observacion_calificacion_calificacion_estudiante"`
	Tiempo_utilizado_calificacion_estudiante             *int         `json:"tiempo_utilizado_calificacion_estudiante"`
	Valoracion_general_calificacion_estudiante           *int         `json:"valoracion_general_calificacion_estudiante"`
	Puntajes_calificacion_estudiante                     []putPuntaje `json:"puntajes_calificacion_estudiante"`
}

func putRequestParse(u *putRequest) {
	if u.Nombre_calificacion_estudiante != nil {
		*u.Nombre_calificacion_estudiante = strings.TrimSpace(*u.Nombre_calificacion_estudiante)
		*u.Nombre_calificacion_estudiante = strings.ToUpper(*u.Nombre_calificacion_estudiante)
		*u.Nombre_calificacion_estudiante = utils.RemoveAccents(*u.Nombre_calificacion_estudiante)
	}

	if u.Entorno_clinico_calificacion_estudiante != nil {
		*u.Entorno_clinico_calificacion_estudiante = strings.TrimSpace(*u.Entorno_clinico_calificacion_estudiante)
		*u.Entorno_clinico_calificacion_estudiante = strings.ToUpper(*u.Entorno_clinico_calificacion_estudiante)
		*u.Entorno_clinico_calificacion_estudiante = utils.RemoveAccents(*u.Entorno_clinico_calificacion_estudiante)
	}

	if u.Paciente_calificacion_estudiante != nil {
		*u.Paciente_calificacion_estudiante = strings.TrimSpace(*u.Paciente_calificacion_estudiante)
		*u.Paciente_calificacion_estudiante = strings.ToUpper(*u.Paciente_calificacion_estudiante)
		*u.Paciente_calificacion_estudiante = utils.RemoveAccents(*u.Paciente_calificacion_estudiante)
	}

	if u.Entorno_clinico_calificacion_estudiante != nil {
		*u.Entorno_clinico_calificacion_estudiante = strings.TrimSpace(*u.Entorno_clinico_calificacion_estudiante)
		*u.Entorno_clinico_calificacion_estudiante = strings.ToUpper(*u.Entorno_clinico_calificacion_estudiante)
		*u.Entorno_clinico_calificacion_estudiante = utils.RemoveAccents(*u.Entorno_clinico_calificacion_estudiante)
	}

	if u.Asunto_principal_consulta_calificacion_estudiante != nil {
		*u.Asunto_principal_consulta_calificacion_estudiante = strings.TrimSpace(*u.Asunto_principal_consulta_calificacion_estudiante)
		*u.Asunto_principal_consulta_calificacion_estudiante = strings.ToUpper(*u.Asunto_principal_consulta_calificacion_estudiante)
		*u.Asunto_principal_consulta_calificacion_estudiante = utils.RemoveAccents(*u.Asunto_principal_consulta_calificacion_estudiante)
	}

	if u.Complejidad_caso_calificacion_estudiante != nil {
		*u.Complejidad_caso_calificacion_estudiante = strings.TrimSpace(*u.Complejidad_caso_calificacion_estudiante)
		*u.Complejidad_caso_calificacion_estudiante = strings.ToUpper(*u.Complejidad_caso_calificacion_estudiante)
		*u.Complejidad_caso_calificacion_estudiante = utils.RemoveAccents(*u.Complejidad_caso_calificacion_estudiante)
	}

	if u.Numero_observaciones_previas_calificacion_estudiante != nil {
		*u.Numero_observaciones_previas_calificacion_estudiante = strings.TrimSpace(*u.Numero_observaciones_previas_calificacion_estudiante)
		*u.Numero_observaciones_previas_calificacion_estudiante = strings.ToUpper(*u.Numero_observaciones_previas_calificacion_estudiante)
		*u.Numero_observaciones_previas_calificacion_estudiante = utils.RemoveAccents(*u.Numero_observaciones_previas_calificacion_estudiante)
	}

	if u.Categoria_observador_calificacion_estudiante != nil {
		*u.Categoria_observador_calificacion_estudiante = strings.TrimSpace(*u.Categoria_observador_calificacion_estudiante)
		*u.Categoria_observador_calificacion_estudiante = strings.ToUpper(*u.Categoria_observador_calificacion_estudiante)
		*u.Categoria_observador_calificacion_estudiante = utils.RemoveAccents(*u.Categoria_observador_calificacion_estudiante)
	}

	if u.Observacion_calificacion_calificacion_estudiante != nil {
		*u.Observacion_calificacion_calificacion_estudiante = strings.TrimSpace(*u.Observacion_calificacion_calificacion_estudiante)
		*u.Observacion_calificacion_calificacion_estudiante = strings.ToUpper(*u.Observacion_calificacion_calificacion_estudiante)
		*u.Observacion_calificacion_calificacion_estudiante = utils.RemoveAccents(*u.Observacion_calificacion_calificacion_estudiante)
	}

	if u.Puntajes_calificacion_estudiante != nil {
		for i := 0; i < len(u.Puntajes_calificacion_estudiante); i++ {
			if u.Puntajes_calificacion_estudiante[i].Feedback_puntaje != nil {
				*u.Puntajes_calificacion_estudiante[i].Feedback_puntaje = strings.TrimSpace(*u.Puntajes_calificacion_estudiante[i].Feedback_puntaje)
				*u.Puntajes_calificacion_estudiante[i].Feedback_puntaje = strings.ToUpper(*u.Puntajes_calificacion_estudiante[i].Feedback_puntaje)
				*u.Puntajes_calificacion_estudiante[i].Feedback_puntaje = utils.RemoveAccents(*u.Puntajes_calificacion_estudiante[i].Feedback_puntaje)
			}
		}
	}
}

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
func Put(c *gin.Context) {
	// id_evaluador := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")
	// id_curso := c.Params.ByName("id_curso")
	// id_grupo := c.Params.ByName("id_grupo")
	id_estudiante := c.Params.ByName("id_estudiante")
	id_evaluacion := c.Params.ByName("id_evaluacion")

	var input putRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	// TODO: validar que vengan las 7 competencias necesarias
	// if len(input.Puntajes_calificacion_estudiante) != 7 {
	// 	api_helpers.RespondError(c, 400, "Deben venir las 7 competencias")
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
