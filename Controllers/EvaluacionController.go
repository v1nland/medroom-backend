package Controllers

import (
	"medroom-backend/ApiHelpers"
	"medroom-backend/Formats/Input"
	"medroom-backend/Formats/Output"
	"medroom-backend/Messages/Request"
	"medroom-backend/Models"
	"medroom-backend/Repositories"
	"medroom-backend/Utils"

	"github.com/gin-gonic/gin"
)

/*
// @Summary Lista de evaluaciones
// @Description Lista todos los evaluaciones
// @Tags Evaluaciones
// @Accept  json
// @Produce  json
// @Success 200 {array} SwaggerMessages.ListEvaluacionesSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluaciones [get]

func ListEvaluaciones(c *gin.Context) {
	// model container
	var container []Models.Evaluacion

	// query
	err := Repositories.GetAllEvaluaciones(&container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.GetEvaluacionesOutput(container))
}

// @Summary Obtiene un evaluacion
// @Description Obtiene un evaluacion según su Id
// @Tags Evaluaciones
// @Accept  json
// @Produce  json
// @Param   id_evaluacion     path    string     true        "Id del evaluacion a buscar"
// @Success 200 {object} SwaggerMessages.GetOneEvaluacionSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluaciones/{id_evaluacion} [get]

func GetOneEvaluacion(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.Evaluacion

	// query
	err := Repositories.GetOneEvaluacion(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.GetOneEvaluacionOutput(container))
}

// @Summary Agrega un nuevo evaluacion
// @Description Genera un nuevo evaluacion con los datos entregados
// @Tags Evaluaciones
// @Accept  json
// @Produce  json
// @Param   input_evaluacion     body    Request.AddNewEvaluacionPayload     true        "Evaluacion a agregar"
// @Success 200 {object} SwaggerMessages.AddNewEvaluacionSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluaciones [post]

func AddNewEvaluacion(c *gin.Context) {
	// input container
	var container Request.AddNewEvaluacionPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	Input.AddNewEvaluacionInput(&container)

	// generate model entity
	model_container := Models.Evaluacion{
		Id_estudiante:                           container.Id_estudiante,
		Id_evaluador:                            container.Id_evaluador,
		Id_periodo:                              container.Id_periodo,
		Nombre_evaluacion:                       container.Nombre_evaluacion,
		Entorno_clinico_evaluacion:              container.Entorno_clinico_evaluacion,
		Paciente_evaluacion:                     container.Paciente_evaluacion,
		Asunto_principal_consulta_evaluacion:    container.Asunto_principal_consulta_evaluacion,
		Complejidad_caso_evaluacion:             container.Complejidad_caso_evaluacion,
		Numero_observaciones_previas_evaluacion: container.Numero_observaciones_previas_evaluacion,
		Categoria_observador_evaluacion:         container.Categoria_observador_evaluacion,
		Observacion_calificacion_evaluacion:     container.Observacion_calificacion_evaluacion,
		Tiempo_utilizado_evaluacion:             container.Tiempo_utilizado_evaluacion,
	}

	// query
	err := Repositories.AddNewEvaluacion(&model_container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.AddNewEvaluacionOutput(model_container))
}

// @Summary Modifica un evaluacion
// @Description Modifica un evaluacion con los datos entregados
// @Tags Evaluaciones
// @Accept  json
// @Produce  json
// @Param   id_evaluacion     path    string     true        "Id del evaluacion a modificar"
// @Param   input_actualiza_evaluacion     body    Request.PutOneEvaluacionPayload     true        "Evaluacion a modificar"
// @Success 200 {object} SwaggerMessages.PutOneEvaluacionSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluaciones/{id_evaluacion} [put]

func PutOneEvaluacion(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// input container
	var container Request.PutOneEvaluacionPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	Input.PutOneEvaluacionInput(&container)

	// generate model entity
	var model_container Models.Evaluacion

	// get query
	err := Repositories.GetOneEvaluacion(&model_container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// replace data in model entity
	model_container = Models.Evaluacion{
		Id:                                      model_container.Id,
		Id_estudiante:                           Utils.CheckUpdatedString(container.Id_estudiante, model_container.Id_estudiante),
		Id_evaluador:                            Utils.CheckUpdatedString(container.Id_evaluador, model_container.Id_evaluador),
		Id_periodo:                              Utils.CheckUpdatedInt(container.Id_periodo, model_container.Id_periodo),
		Nombre_evaluacion:                       Utils.CheckUpdatedString(container.Nombre_evaluacion, model_container.Nombre_evaluacion),
		Entorno_clinico_evaluacion:              Utils.CheckUpdatedString(container.Entorno_clinico_evaluacion, model_container.Entorno_clinico_evaluacion),
		Paciente_evaluacion:                     Utils.CheckUpdatedString(container.Paciente_evaluacion, model_container.Paciente_evaluacion),
		Asunto_principal_consulta_evaluacion:    Utils.CheckUpdatedString(container.Asunto_principal_consulta_evaluacion, model_container.Asunto_principal_consulta_evaluacion),
		Complejidad_caso_evaluacion:             Utils.CheckUpdatedString(container.Complejidad_caso_evaluacion, model_container.Complejidad_caso_evaluacion),
		Numero_observaciones_previas_evaluacion: Utils.CheckUpdatedString(container.Numero_observaciones_previas_evaluacion, model_container.Numero_observaciones_previas_evaluacion),
		Categoria_observador_evaluacion:         Utils.CheckUpdatedString(container.Categoria_observador_evaluacion, model_container.Categoria_observador_evaluacion),
		Observacion_calificacion_evaluacion:     Utils.CheckUpdatedString(container.Observacion_calificacion_evaluacion, model_container.Observacion_calificacion_evaluacion),
		Tiempo_utilizado_evaluacion:             Utils.CheckUpdatedInt(container.Tiempo_utilizado_evaluacion, model_container.Tiempo_utilizado_evaluacion),
	}

	// update foreign entity
	err = Repositories.GetOneEvaluador(&model_container.Evaluador_evaluacion, model_container.Id_evaluador)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// update foreign entity
	err = Repositories.GetOnePeriodo(&model_container.Periodo_evaluacion, Utils.ConvertIntToString(model_container.Id_periodo))
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// put query
	err = Repositories.PutOneEvaluacion(&model_container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.PutOneEvaluacionOutput(model_container))
}

// @Summary Elimina un evaluacion
// @Description Elimina un evaluacion con los datos entregados
// @Tags Evaluaciones
// @Accept  json
// @Produce  json
// @Param   id_evaluacion     path    string     true        "Id del evaluacion a eliminar"
// @Success 200 {object} SwaggerMessages.DeleteEvaluacionSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluaciones/{id_evaluacion} [delete]

func DeleteEvaluacion(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.Evaluacion

	// get query
	err := Repositories.GetOneEvaluacion(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// query
	err = Repositories.DeleteEvaluacion(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.DeleteEvaluacionOutput(container))
}*/

// @Summary Lista de evaluaciones de un estudiante
// @Description Lista todos los evaluaciones de un estudiante
// @Tags 02 - Estudiantes
// @Accept  json
// @Produce  json
// @Success 200 {array} SwaggerMessages.ListEvaluacionesSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes/me/evaluaciones [get]
func ListEvaluacionesEstudiante(c *gin.Context) {
	// params
	id_estudiante := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	// model container
	var estudiante Models.Estudiante
	if err := Repositories.GetOneEstudiante(&estudiante, id_estudiante); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.GetEvaluacionesEstudianteOutput(estudiante.Evaluaciones_estudiante))
}

// @Summary Genera una evaluación para un estudiante
// @Description Genera una nueva evaluación de un estudiante con los datos entregados
// @Tags 03 - Evaluadores
// @Accept  json
// @Produce  json
// @Param   input_evaluacion     body    Request.GenerarEvaluacionPayload     true        "Evaluacion a generar"
// @Success 200 {object} SwaggerMessages.GenerarEvaluacionSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluadores/me/evaluaciones [post]
func GenerarEvaluacion(c *gin.Context) {
	// params
	id_evaluador := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")

	// input container
	var container Request.GenerarEvaluacionPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	Input.GenerarEvaluacionInput(&container)

	// validate puntajes_evaluacion enum

	// generate model entity
	model_container := Models.Evaluacion{
		Id_estudiante:                           container.Id_estudiante,
		Id_evaluador:                            id_evaluador,
		Id_periodo:                              container.Id_periodo,
		Nombre_evaluacion:                       container.Nombre_evaluacion,
		Entorno_clinico_evaluacion:              container.Entorno_clinico_evaluacion,
		Paciente_evaluacion:                     container.Paciente_evaluacion,
		Asunto_principal_consulta_evaluacion:    container.Asunto_principal_consulta_evaluacion,
		Complejidad_caso_evaluacion:             container.Complejidad_caso_evaluacion,
		Numero_observaciones_previas_evaluacion: container.Numero_observaciones_previas_evaluacion,
		Categoria_observador_evaluacion:         container.Categoria_observador_evaluacion,
		Observacion_calificacion_evaluacion:     container.Observacion_calificacion_evaluacion,
		Tiempo_utilizado_evaluacion:             container.Tiempo_utilizado_evaluacion,
	}

	// query
	if err := Repositories.AddNewEvaluacion(&model_container); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	for i := 0; i < len(container.Puntajes_evaluacion); i++ {
		puntaje := Models.Puntaje{
			Id_evaluacion:              model_container.Id,
			Nombre_competencia_puntaje: container.Puntajes_evaluacion[i].Nombre_competencia,
			Codigo_competencia_puntaje: container.Puntajes_evaluacion[i].Codigo_competencia,
			Calificacion_puntaje:       container.Puntajes_evaluacion[i].Puntaje_competencia,
			Feedback_puntaje:           container.Puntajes_evaluacion[i].Feedback_competencia,
		}

		if err := Repositories.AddNewPuntaje(&puntaje); err != nil {
			ApiHelpers.RespondError(c, 500, "default")
			return
		}
	}

	// output
	ApiHelpers.RespondJSON(c, 200, Output.GenerarEvaluacionOutput(model_container))
}
