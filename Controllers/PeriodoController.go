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
	*  FUNCIÓN ListPeriodo
	*
    *
	*
	*
    *
*/

// @Summary Lista de periodos
// @Description Lista todos los periodos
// @Tags Periodos
// @Accept  json
// @Produce  json
// @Success 200 {array} SwaggerMessages.ListPeriodosSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /periodos [get]
func ListPeriodos(c *gin.Context) {
	// model container
	var container []Models.Periodo

	// query
	err := Repositories.GetAllPeriodos(&container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetPeriodosOutput(container))
}

/*
	*
	*  FUNCIÓN GetOnePeriodo
	*
    *
	*
	*
    *
*/

// @Summary Obtiene un periodo
// @Description Obtiene un periodo según su UUID
// @Tags Periodos
// @Accept  json
// @Produce  json
// @Param   uuid_periodo     path    string     true        "UUID del periodo a buscar"
// @Success 200 {object} SwaggerMessages.GetOnePeriodoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /periodos/{uuid_periodo} [get]
func GetOnePeriodo(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.Periodo

	// query
	err := Repositories.GetOnePeriodo(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetOnePeriodoOutput(container))
}

/*
	*
	*  FUNCIÓN AddNewPeriodo
	*
    *
	*
	*
    *
*/

// @Summary Agrega un nuevo periodo
// @Description Genera un nuevo periodo con los datos entregados
// @Tags Periodos
// @Accept  json
// @Produce  json
// @Param   input_periodo     body    RequestMessages.AddNewPeriodoPayload     true        "Periodo a agregar"
// @Success 200 {object} SwaggerMessages.AddNewPeriodoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /periodos [post]
func AddNewPeriodo(c *gin.Context) {
	// input container
	var container RequestMessages.AddNewPeriodoPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.AddNewPeriodoInput(&container)

	// generate model entity
	model_container := Models.Periodo{
		Nombre_periodo: container.Nombre_periodo,
	}

	// query
	err := Repositories.AddNewPeriodo(&model_container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.AddNewPeriodoOutput(model_container))
}

/*
	*
	*  FUNCIÓN PutOnePeriodo
	*
    *
	*
	*
    *
*/

// @Summary Modifica un periodo
// @Description Modifica un periodo con los datos entregados
// @Tags Periodos
// @Accept  json
// @Produce  json
// @Param   uuid_periodo     path    string     true        "UUID del periodo a modificar"
// @Param   input_actualiza_periodo     body    RequestMessages.PutOnePeriodoPayload     true        "Periodo a modificar"
// @Success 200 {object} SwaggerMessages.PutOnePeriodoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /periodos/{uuid_periodo} [put]
func PutOnePeriodo(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// input container
	var container RequestMessages.PutOnePeriodoPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.PutOnePeriodoInput(&container)

	// generate model entity
	var model_container Models.Periodo

	// get query
	err := Repositories.GetOnePeriodo(&model_container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// replace data in model entity
	model_container = Models.Periodo{
		ID:             model_container.ID,
		Nombre_periodo: Utils.CheckUpdatedString(container.Nombre_periodo, model_container.Nombre_periodo),
	}

	// update foreign entity
	// err = Repositories.GetOneEvaluador(&model_container.Evaluador_periodo, model_container.Id_evaluador)
	// if err != nil {
	// 	ApiHelpers.RespondError(c, 500, "default")
	// 	return
	// }

	// put query
	err = Repositories.PutOnePeriodo(&model_container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.PutOnePeriodoOutput(model_container))
}

/*
	*
	*  FUNCIÓN DeletePeriodo
	*
    *
	*
	*
    *
*/

// @Summary Elimina un periodo
// @Description Elimina un periodo con los datos entregados
// @Tags Periodos
// @Accept  json
// @Produce  json
// @Param   uuid_periodo     path    string     true        "UUID del periodo a eliminar"
// @Success 200 {object} SwaggerMessages.DeletePeriodoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /periodos/{uuid_periodo} [delete]
func DeletePeriodo(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.Periodo

	// get query
	err := Repositories.GetOnePeriodo(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// query
	err = Repositories.DeletePeriodo(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.DeletePeriodoOutput(container))
}
