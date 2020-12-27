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
	"gorm.io/gorm"
)

/*
// @Summary Lista de evaluaciones
// @Description Lista todos los evaluaciones
// @Tags Evaluaciones
// @Accept  json
// @Produce  json
// @Success 200 {array} Swagger.ListEvaluacionesSwagger "OK"
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
// @Success 200 {object} Swagger.GetOneEvaluacionSwagger "OK"
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
// @Success 200 {object} Swagger.AddNewEvaluacionSwagger "OK"
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
// @Success 200 {object} Swagger.PutOneEvaluacionSwagger "OK"
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
// @Success 200 {object} Swagger.DeleteEvaluacionSwagger "OK"
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

// @Summary Lista de evaluaciones de un grupo
// @Description Lista todas los evaluaciones disponibles de un estudiante de un grupo
// @Tags 02 - Estudiantes
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso"
// @Param   id_grupo     path    string     true        "Id del grupo"
// @Success 200 {array} Swagger.ListEvaluacionesGrupoEstudianteSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes/me/cursos/{id_curso}/grupos/{id_grupo}/evaluaciones [get]
func ListEvaluacionesGrupoEstudiante(c *gin.Context) {
	id_grupo := c.Params.ByName("id_grupo")
	id_curso := c.Params.ByName("id_curso")
	id_estudiante := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	var grupo Models.Grupo
	if err := Repositories.GetOneGrupoEstudiante(&grupo, id_grupo, id_curso, id_estudiante); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Grupo not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.ListEvaluacionesEstudianteOutput(grupo.Evaluaciones_grupo))
	ApiHelpers.RespondJSON(c, 200, grupo.Evaluaciones_grupo)
}

// @Summary Lista de evaluaciones de un grupo
// @Description Lista todas los evaluaciones disponibles de un evaluador de un grupo
// @Tags 03 - Evaluadores
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso"
// @Param   id_grupo     path    string     true        "Id del grupo"
// @Success 200 {array} Swagger.ListEvaluacionesGrupoEvaluadorSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluadores/me/cursos/{id_curso}/grupos/{id_grupo}/evaluaciones [get]
func ListEvaluacionesGrupoEvaluador(c *gin.Context) {
	id_grupo := c.Params.ByName("id_grupo")
	id_curso := c.Params.ByName("id_curso")
	id_evaluador := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")

	var grupo Models.Grupo
	if err := Repositories.GetOneGrupoEvaluador(&grupo, id_grupo, id_curso, id_evaluador); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondJSON(c, 200, "Grupo not found")
		} else {
			ApiHelpers.RespondError(c, 500, "default")
		}

		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.ListEvaluacionesEstudianteOutput(grupo.Evaluaciones_grupo))
	ApiHelpers.RespondJSON(c, 200, grupo.Evaluaciones_grupo)
}

// @Summary Agrega una evaluación
// @Description Genera una evaluación para un grupo
// @Tags 03 - Evaluadores
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso"
// @Param   id_grupo     path    string     true        "Id del grupo"
// @Param   input_evaluacion     body    Request.AddNewEvaluacionPayload     true        "Evaluacion a agregar"
// @Success 200 {array} Swagger.AddNewEvaluacionSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /evaluadores/me/cursos/{id_curso}/grupos/{id_grupo}/evaluaciones [post]
func AddNewEvaluacion(c *gin.Context) {
	// id_evaluador := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")
	// id_curso := c.Params.ByName("id_curso")
	id_grupo := c.Params.ByName("id_grupo")

	var input Request.AddNewEvaluacionPayload
	if err := c.ShouldBind(&input); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	Input.AddNewEvaluacionInput(&input)

	model := Models.Evaluacion{
		Id_grupo:          Utils.ConvertStringToInt(id_grupo),
		Nombre_evaluacion: *input.Nombre_evaluacion,
	}

	if err := Repositories.AddNewEvaluacion(&model); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// ApiHelpers.RespondJSON(c, 200, Output.ListEvaluacionesEstudianteOutput(model))
	ApiHelpers.RespondJSON(c, 200, model)
}
