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

// @Summary Lista de periodos
// @Description Lista todos los periodos
// @Tags 00 - Rutas públicas
// @Accept  json
// @Produce  json
// @Success 200 {array} Swagger.ListPeriodosSwagger "OK"
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
	ApiHelpers.RespondJSON(c, 200, Output.ListPeriodosOutput(container))
}

// @Summary Obtiene un periodo
// @Description Obtiene un periodo según su Id
// @Tags 00 - Rutas públicas
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo a buscar"
// @Success 200 {object} Swagger.GetOnePeriodoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /periodos/{id_periodo} [get]
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
	ApiHelpers.RespondJSON(c, 200, Output.GetOnePeriodoOutput(container))
}

// @Summary Agrega un nuevo periodo
// @Description Genera un nuevo periodo con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_periodo     body    Request.AddNewPeriodoPayload     true        "Periodo a agregar"
// @Success 200 {object} Swagger.AddNewPeriodoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/periodos [post]
func AddNewPeriodo(c *gin.Context) {
	// input container
	var container Request.AddNewPeriodoPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	Input.AddNewPeriodoInput(&container)

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
	ApiHelpers.RespondJSON(c, 200, Output.AddNewPeriodoOutput(model_container))
}

// @Summary Modifica un periodo
// @Description Modifica un periodo con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo a modificar"
// @Param   input_actualiza_periodo     body    Request.PutOnePeriodoPayload     true        "Periodo a modificar"
// @Success 200 {object} Swagger.PutOnePeriodoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/periodos/{id_periodo} [put]
func PutOnePeriodo(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// input container
	var container Request.PutOnePeriodoPayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	Input.PutOnePeriodoInput(&container)

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
		Id:             model_container.Id,
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
	ApiHelpers.RespondJSON(c, 200, Output.PutOnePeriodoOutput(model_container))
}

// @Summary Elimina un periodo
// @Description Elimina un periodo con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo a eliminar"
// @Success 200 {object} Swagger.DeletePeriodoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/periodos/{id_periodo} [delete]
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
	ApiHelpers.RespondJSON(c, 200, Output.DeletePeriodoOutput(container))
}
