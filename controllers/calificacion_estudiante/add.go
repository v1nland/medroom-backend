package calificacion_estudiante

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type addPuntaje struct {
	Id_competencia       *string `json:"id_competencia"`
	Calificacion_puntaje *int    `json:"calificacion_puntaje"`
}
type addRequest struct {
	Observacion_calificacion_calificacion_estudiante *string      `json:"observacion_calificacion_calificacion_estudiante"`
	Valoracion_general_calificacion_estudiante       *int         `json:"valoracion_general_calificacion_estudiante"`
	Puntajes_calificacion_estudiante                 []addPuntaje `json:"puntajes_calificacion_estudiante"`
}

func addRequestParse(u *addRequest) {
	if u.Observacion_calificacion_calificacion_estudiante != nil {
		*u.Observacion_calificacion_calificacion_estudiante = strings.TrimSpace(*u.Observacion_calificacion_calificacion_estudiante)
		*u.Observacion_calificacion_calificacion_estudiante = strings.ToUpper(*u.Observacion_calificacion_calificacion_estudiante)
		*u.Observacion_calificacion_calificacion_estudiante = utils.RemoveAccents(*u.Observacion_calificacion_calificacion_estudiante)
	}
}

// @Summary Califica un estudiante
// @Description Genera una calificación para un estudiante
// @Tags Evaluadores
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo"
// @Param   sigla_curso     path    string     true        "Sigla del curso"
// @Param   sigla_grupo     path    string     true        "Sigla del grupo"
// @Param   id_estudiante     path    string     true        "Id del estudiante"
// @Param   id_evaluacion     path    string     true        "Id de la evaluacion"
// @Param   input_calificacion_estudiante     body    addRequest     true        "CalificacionEstudiante a agregar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /evaluadores/me/cursos/{id_periodo}/{sigla_curso}/grupos/{sigla_grupo}/estudiantes/{id_estudiante}/evaluaciones/{id_evaluacion}/calificacion [post]
func Add(c *gin.Context) {
	id_evaluador := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")
	id_estudiante := c.Params.ByName("id_estudiante")
	id_evaluacion := c.Params.ByName("id_evaluacion")
	id_periodo := c.Params.ByName("id_periodo")
	// sigla_curso := c.Params.ByName("sigla_curso")
	// sigla_grupo := c.Params.ByName("sigla_grupo")

	var input addRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	// TODO: validar que vengan las 6 competencias necesarias
	if len(input.Puntajes_calificacion_estudiante) != 6 {
		api_helpers.RespondError(c, 400, "deben venir las 6 competencias")
		return
	}

	addRequestParse(&input)

	// parameter for calificacion estudiante constructor
	var puntajes_calificacion_estudiante []models.Puntaje
	for i := 0; i < len(input.Puntajes_calificacion_estudiante); i++ {
		puntajes_calificacion_estudiante = append(puntajes_calificacion_estudiante, models.Puntaje{
			Id_competencia:       *input.Puntajes_calificacion_estudiante[i].Id_competencia,
			Calificacion_puntaje: *input.Puntajes_calificacion_estudiante[i].Calificacion_puntaje,
		})
	}

	model := models.CalificacionEstudiante{
		Id_evaluador:                     uuid.MustParse(id_evaluador),
		Id_periodo:                       id_periodo,
		Id_evaluacion:                    utils.StringToInt(id_evaluacion),
		Puntajes_calificacion_estudiante: puntajes_calificacion_estudiante,
		Id_estudiante:                    uuid.MustParse(id_estudiante),
		Observacion_calificacion_calificacion_estudiante: *input.Observacion_calificacion_calificacion_estudiante,
		Valoracion_general_calificacion_estudiante:       *input.Valoracion_general_calificacion_estudiante,
	}

	if err := repositories.AddNewCalificacionEstudiante(&model); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.ListEvaluacionesEstudiante(model))
	api_helpers.RespondJson(c, 200, model)
}
