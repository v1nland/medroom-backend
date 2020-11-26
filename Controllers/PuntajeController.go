package Controllers

import (
	"github.com/gin-gonic/gin"
	"medroom-backend/ApiHelpers"
	"medroom-backend/InputFormats"
	"medroom-backend/Models"
	"medroom-backend/OutputFormats"
	"medroom-backend/Repositories"
	"medroom-backend/RequestMessages"
	"medroom-backend/Utils"
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
// @Success 200 {array} ResponseMessages.ListPuntajesResponse "OK"
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
// @Description Obtiene un puntaje según su UUID
// @Tags Puntajes
// @Accept  json
// @Produce  json
// @Param   uuid_puntaje     path    string     true        "UUID del puntaje a buscar"
// @Success 200 {object} ResponseMessages.GetOnePuntajeResponse "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /puntajes/{uuid_puntaje} [get]
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
// @Success 200 {object} ResponseMessages.AddNewPuntajeResponse "OK"
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
		Id_evaluacion:        container.Id_evaluacion,
		Id_competencia:       container.Id_competencia,
		Calificacion_puntaje: container.Calificacion_puntaje,
		Nivel_logro_puntaje:  container.Nivel_logro_puntaje,
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
// @Param   uuid_puntaje     path    string     true        "UUID del puntaje a modificar"
// @Param   input_actualiza_puntaje     body    RequestMessages.PutOnePuntajePayload     true        "Puntaje a modificar"
// @Success 200 {object} ResponseMessages.PutOnePuntajeResponse "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /puntajes/{uuid_puntaje} [put]
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
		ID:                   model_container.ID,
		Id_evaluacion:        Utils.CheckUpdatedInt(container.Id_evaluacion, model_container.Id_evaluacion),
		Id_competencia:       Utils.CheckUpdatedInt(container.Id_competencia, model_container.Id_competencia),
		Calificacion_puntaje: Utils.CheckUpdatedInt(container.Calificacion_puntaje, model_container.Calificacion_puntaje),
		Nivel_logro_puntaje:  Utils.CheckUpdatedString(container.Nivel_logro_puntaje, model_container.Nivel_logro_puntaje),
	}

	// update foreign entity
	// err = Repositories.GetOneEvaluacion(&model_container.Evaluacion_puntaje, Utils.ConvertIntToString(model_container.Id_evaluacion))
	// if err != nil {
	// 	ApiHelpers.RespondError(c, 500, "default")
	// 	return
	// }

	// update foreign entity
	err = Repositories.GetOneCompetencia(&model_container.Competencia_puntaje, Utils.ConvertIntToString(model_container.Id_competencia))
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

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
// @Param   uuid_puntaje     path    string     true        "UUID del puntaje a eliminar"
// @Success 200 {object} ResponseMessages.DeletePuntajeResponse "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /puntajes/{uuid_puntaje} [delete]
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
