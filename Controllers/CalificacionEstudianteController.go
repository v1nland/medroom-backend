package Controllers

import (
	"errors"
	"medroom-backend/ApiHelpers"
	"medroom-backend/Formats/Input"
	"medroom-backend/Messages/Request"
	"medroom-backend/Models"
	"medroom-backend/Repositories"
	"medroom-backend/Utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes/me/cursos/{id_curso}/grupos/{id_grupo}/evaluaciones/{id_evaluacion}/calificacion [get]
func GetOneCalificacionEstudiante(c *gin.Context) {
	id_evaluacion := c.Params.ByName("id_evaluacion")
	id_estudiante := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	var calificacion_estudiante Models.CalificacionEstudiante
	if err := Repositories.GetOneCalificacionEstudianteByIdEvaluacion(&calificacion_estudiante, id_evaluacion, id_estudiante); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Calificación estudiante not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.ListEvaluacionesEstudianteOutput(calificacion_estudiante))
	ApiHelpers.RespondJSON(c, 200, calificacion_estudiante)
}

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
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluadores/me/cursos/{id_curso}/grupos/{id_grupo}/estudiantes/{id_estudiante}/evaluaciones/{id_evaluacion}/calificacion [post]
func AddNewCalificacionEstudiante(c *gin.Context) {
	id_evaluador := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")
	// id_curso := c.Params.ByName("id_curso")
	// id_grupo := c.Params.ByName("id_grupo")
	id_estudiante := c.Params.ByName("id_estudiante")
	id_evaluacion := c.Params.ByName("id_evaluacion")

	var input Request.AddNewCalificacionEstudiante
	if err := c.ShouldBind(&input); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// TODO: validar que vengan las 7 competencias necesarias
	if len(input.Puntajes_calificacion_estudiante) != 7 {
		ApiHelpers.RespondError(c, 400, "deben venir las 7 competencias")
		return
	}

	Input.AddNewCalificacionEstudiante(&input)

	// parameter for calificacion estudiante constructor
	var puntajes_calificacion_estudiante []Models.Puntaje
	for i := 0; i < len(input.Puntajes_calificacion_estudiante); i++ {
		puntajes_calificacion_estudiante = append(puntajes_calificacion_estudiante, Models.Puntaje{
			Id_competencia:       *input.Puntajes_calificacion_estudiante[i].Id_competencia,
			Calificacion_puntaje: *input.Puntajes_calificacion_estudiante[i].Calificacion_puntaje,
			Feedback_puntaje:     *input.Puntajes_calificacion_estudiante[i].Feedback_puntaje,
		})
	}

	model := Models.CalificacionEstudiante{
		Id_evaluador:                                         uuid.MustParse(id_evaluador),
		Id_periodo:                                           *input.Id_periodo,
		Id_evaluacion:                                        Utils.ConvertStringToInt(id_evaluacion),
		Puntajes_calificacion_estudiante:                     puntajes_calificacion_estudiante,
		Id_estudiante:                                        uuid.MustParse(id_estudiante),
		Nombre_calificacion_estudiante:                       *input.Nombre_calificacion_estudiante,
		Entorno_clinico_calificacion_estudiante:              *input.Entorno_clinico_calificacion_estudiante,
		Paciente_calificacion_estudiante:                     *input.Paciente_calificacion_estudiante,
		Asunto_principal_consulta_calificacion_estudiante:    *input.Asunto_principal_consulta_calificacion_estudiante,
		Complejidad_caso_calificacion_estudiante:             *input.Complejidad_caso_calificacion_estudiante,
		Numero_observaciones_previas_calificacion_estudiante: *input.Numero_observaciones_previas_calificacion_estudiante,
		Categoria_observador_calificacion_estudiante:         *input.Categoria_observador_calificacion_estudiante,
		Observacion_calificacion_calificacion_estudiante:     *input.Observacion_calificacion_calificacion_estudiante,
		Tiempo_utilizado_calificacion_estudiante:             *input.Tiempo_utilizado_calificacion_estudiante,
	}

	if err := Repositories.AddNewCalificacionEstudiante(&model); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.ListEvaluacionesEstudianteOutput(model))
	ApiHelpers.RespondJSON(c, 200, model)
}
