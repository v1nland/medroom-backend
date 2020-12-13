package Controllers

import (
	"medroom-backend/ApiHelpers"
	"medroom-backend/InputFormats"
	"medroom-backend/Models"
	"medroom-backend/OutputFormats"
	"medroom-backend/Repositories"
	"medroom-backend/RequestMessages"
	"medroom-backend/Utils"

	"github.com/gin-gonic/gin"
)

/*
	*
	*  FUNCIÓN ListEvaluacion
	*
    *
	*
	*
    *
*/

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
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetEvaluacionesOutput(container))
}

/*
	*
	*  FUNCIÓN GetOneEvaluacion
	*
    *
	*
	*
    *
*/

// @Summary Obtiene un evaluacion
// @Description Obtiene un evaluacion según su ID
// @Tags Evaluaciones
// @Accept  json
// @Produce  json
// @Param   id_evaluacion     path    string     true        "ID del evaluacion a buscar"
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
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetOneEvaluacionOutput(container))
}

/*
	*
	*  FUNCIÓN AddNewEvaluacion
	*
    *
	*
	*
    *
*/

// @Summary Agrega un nuevo evaluacion
// @Description Genera un nuevo evaluacion con los datos entregados
// @Tags Evaluaciones
// @Accept  json
// @Produce  json
// @Param   input_evaluacion     body    RequestMessages.AddNewEvaluacionPayload     true        "Evaluacion a agregar"
// @Success 200 {object} SwaggerMessages.AddNewEvaluacionSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluaciones [post]
func AddNewEvaluacion(c *gin.Context) {
	// input container
	var container RequestMessages.AddNewEvaluacionPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.AddNewEvaluacionInput(&container)

	// generate model entity
	model_container := Models.Evaluacion{
		Id_estudiante:                           container.Id_estudiante,
		Id_evaluador:                            container.Id_evaluador,
		Id_competencia:                          container.Id_competencia,
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
	ApiHelpers.RespondJSON(c, 200, OutputFormats.AddNewEvaluacionOutput(model_container))
}

/*
	*
	*  FUNCIÓN PutOneEvaluacion
	*
    *
	*
	*
    *
*/

// @Summary Modifica un evaluacion
// @Description Modifica un evaluacion con los datos entregados
// @Tags Evaluaciones
// @Accept  json
// @Produce  json
// @Param   id_evaluacion     path    string     true        "ID del evaluacion a modificar"
// @Param   input_actualiza_evaluacion     body    RequestMessages.PutOneEvaluacionPayload     true        "Evaluacion a modificar"
// @Success 200 {object} SwaggerMessages.PutOneEvaluacionSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluaciones/{id_evaluacion} [put]
func PutOneEvaluacion(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// input container
	var container RequestMessages.PutOneEvaluacionPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.PutOneEvaluacionInput(&container)

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
		ID:                                      model_container.ID,
		Id_estudiante:                           Utils.CheckUpdatedString(container.Id_estudiante, model_container.Id_estudiante),
		Id_evaluador:                            Utils.CheckUpdatedString(container.Id_evaluador, model_container.Id_evaluador),
		Id_competencia:                          Utils.CheckUpdatedInt(container.Id_competencia, model_container.Id_competencia),
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
	err = Repositories.GetOneEstudiante(&model_container.Estudiante_evaluacion, model_container.Id_estudiante)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// update foreign entity
	err = Repositories.GetOneEvaluador(&model_container.Evaluador_evaluacion, model_container.Id_evaluador)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// update foreign entity
	err = Repositories.GetOneCompetencia(&model_container.Competencia_evaluacion, Utils.ConvertIntToString(model_container.Id_competencia))
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
	ApiHelpers.RespondJSON(c, 200, OutputFormats.PutOneEvaluacionOutput(model_container))
}

/*
	*
	*  FUNCIÓN DeleteEvaluacion
	*
    *
	*
	*
    *
*/

// @Summary Elimina un evaluacion
// @Description Elimina un evaluacion con los datos entregados
// @Tags Evaluaciones
// @Accept  json
// @Produce  json
// @Param   id_evaluacion     path    string     true        "ID del evaluacion a eliminar"
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
	ApiHelpers.RespondJSON(c, 200, OutputFormats.DeleteEvaluacionOutput(container))
}
