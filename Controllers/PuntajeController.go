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
	*  FUNCIÓN ListPuntaje
	*
    *
	*
	*
    *
*/

// @Summary Lista de puntajes
// @Description Lista todos los puntajes
// @Tags Puntajes
// @Accept  json
// @Produce  json
// @Success 200 {array} SwaggerMessages.ListPuntajesSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /puntajes [get]
func ListPuntajes(c *gin.Context) {
	// model container
	var container []Models.Puntaje

	// query
	err := Repositories.GetAllPuntajes(&container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetPuntajesOutput(container))
}

/*
	*
	*  FUNCIÓN GetOnePuntaje
	*
    *
	*
	*
    *
*/

// @Summary Obtiene un puntaje
// @Description Obtiene un puntaje según su ID
// @Tags Puntajes
// @Accept  json
// @Produce  json
// @Param   id_puntaje     path    string     true        "ID del puntaje a buscar"
// @Success 200 {object} SwaggerMessages.GetOnePuntajeSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /puntajes/{id_puntaje} [get]
func GetOnePuntaje(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.Puntaje

	// query
	err := Repositories.GetOnePuntaje(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetOnePuntajeOutput(container))
}

/*
	*
	*  FUNCIÓN AddNewPuntaje
	*
    *
	*
	*
    *
*/

// @Summary Agrega un nuevo puntaje
// @Description Genera un nuevo puntaje con los datos entregados
// @Tags Puntajes
// @Accept  json
// @Produce  json
// @Param   input_puntaje     body    RequestMessages.AddNewPuntajePayload     true        "Puntaje a agregar"
// @Success 200 {object} SwaggerMessages.AddNewPuntajeSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /puntajes [post]
func AddNewPuntaje(c *gin.Context) {
	// input container
	var container RequestMessages.AddNewPuntajePayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.AddNewPuntajeInput(&container)

	// generate model entity
	model_container := Models.Puntaje{
		Id_evaluacion:              container.Id_evaluacion,
		Nombre_competencia_puntaje: container.Nombre_competencia_puntaje,
		Calificacion_puntaje:       container.Calificacion_puntaje,
		Feedback_puntaje:           container.Feedback_puntaje,
	}

	// query
	err := Repositories.AddNewPuntaje(&model_container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.AddNewPuntajeOutput(model_container))
}

/*
	*
	*  FUNCIÓN PutOnePuntaje
	*
    *
	*
	*
    *
*/

// @Summary Modifica un puntaje
// @Description Modifica un puntaje con los datos entregados
// @Tags Puntajes
// @Accept  json
// @Produce  json
// @Param   id_puntaje     path    string     true        "ID del puntaje a modificar"
// @Param   input_actualiza_puntaje     body    RequestMessages.PutOnePuntajePayload     true        "Puntaje a modificar"
// @Success 200 {object} SwaggerMessages.PutOnePuntajeSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /puntajes/{id_puntaje} [put]
func PutOnePuntaje(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// input container
	var container RequestMessages.PutOnePuntajePayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.PutOnePuntajeInput(&container)

	// generate model entity
	var model_container Models.Puntaje

	// get query
	err := Repositories.GetOnePuntaje(&model_container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// replace data in model entity
	model_container = Models.Puntaje{
		Id_evaluacion:              Utils.CheckUpdatedInt(container.Id_evaluacion, model_container.Id_evaluacion),
		Nombre_competencia_puntaje: Utils.CheckUpdatedString(container.Nombre_competencia_puntaje, model_container.Nombre_competencia_puntaje),
		Calificacion_puntaje:       Utils.CheckUpdatedInt(container.Calificacion_puntaje, model_container.Calificacion_puntaje),
		Feedback_puntaje:           Utils.CheckUpdatedString(container.Feedback_puntaje, model_container.Feedback_puntaje),
	}

	// update foreign entity
	// err = Repositories.GetOneEvaluacion(&model_container.Evaluacion_puntaje, Utils.ConvertIntToString(model_container.Id_evaluacion))
	// if err != nil {
	// 	ApiHelpers.RespondError(c, 500, "default")
	// 	return
	// }

	// put query
	err = Repositories.PutOnePuntaje(&model_container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.PutOnePuntajeOutput(model_container))
}

/*
	*
	*  FUNCIÓN DeletePuntaje
	*
    *
	*
	*
    *
*/

// @Summary Elimina un puntaje
// @Description Elimina un puntaje con los datos entregados
// @Tags Puntajes
// @Accept  json
// @Produce  json
// @Param   id_puntaje     path    string     true        "ID del puntaje a eliminar"
// @Success 200 {object} SwaggerMessages.DeletePuntajeSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /puntajes/{id_puntaje} [delete]
func DeletePuntaje(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.Puntaje

	// get query
	err := Repositories.GetOnePuntaje(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// query
	err = Repositories.DeletePuntaje(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.DeletePuntajeOutput(container))
}
