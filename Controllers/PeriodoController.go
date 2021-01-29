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
	var periodos []Models.Periodo

	if err := Repositories.GetAllPeriodos(&periodos); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	ApiHelpers.RespondJSON(c, 200, Output.ListPeriodosOutput(periodos))
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
	id := c.Params.ByName("id")

	var periodo Models.Periodo
	if err := Repositories.GetOnePeriodo(&periodo, id); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	ApiHelpers.RespondJSON(c, 200, Output.GetOnePeriodoOutput(periodo))
}

// @Summary Agrega un nuevo periodo
// @Description Genera un nuevo periodo con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_periodo     body    Request.AddNewPeriodo     true        "Periodo a agregar"
// @Success 200 {object} Swagger.AddNewPeriodoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/periodos [post]
func AddNewPeriodo(c *gin.Context) {
	var input Request.AddNewPeriodo
	if err := c.ShouldBind(&input); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	Input.AddNewPeriodo(&input)

	periodo := Models.Periodo{
		Nombre_periodo: *input.Nombre_periodo,
	}

	if err := Repositories.AddNewPeriodo(&periodo); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	ApiHelpers.RespondJSON(c, 200, Output.AddNewPeriodoOutput(periodo))
}

// @Summary Modifica un periodo
// @Description Modifica un periodo con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo a modificar"
// @Param   input_actualiza_periodo     body    Request.PutOnePeriodo     true        "Periodo a modificar"
// @Success 200 {object} Swagger.PutOnePeriodoSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /administracion-ti/periodos/{id_periodo} [put]
func PutOnePeriodo(c *gin.Context) {
	id := c.Params.ByName("id")

	var input Request.PutOnePeriodo
	if err := c.ShouldBind(&input); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	Input.PutOnePeriodo(&input)

	var periodo Models.Periodo
	if err := Repositories.GetOnePeriodo(&periodo, id); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	periodo = Models.Periodo{
		Id:             periodo.Id,
		Nombre_periodo: Utils.CheckNullString(input.Nombre_periodo, periodo.Nombre_periodo),
	}

	if err := Repositories.PutOnePeriodo(&periodo, id); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	ApiHelpers.RespondJSON(c, 200, Output.PutOnePeriodoOutput(periodo))
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
	id := c.Params.ByName("id")

	var periodo Models.Periodo
	if err := Repositories.GetOnePeriodo(&periodo, id); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	if err := Repositories.DeletePeriodo(&periodo, id); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	ApiHelpers.RespondJSON(c, 200, Output.DeletePeriodoOutput(periodo))
}
