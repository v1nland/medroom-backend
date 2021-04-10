package periodo

import (
	"medroom-backend/api_helpers"
	"medroom-backend/formats/f_input"
	"medroom-backend/messages/Request"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Modifica un periodo
// @Description Modifica un periodo con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo a modificar"
// @Param   input_actualiza_periodo     body    Request.PutOnePeriodo     true        "Periodo a modificar"
// @Success 200 {object} Swagger.PutOnePeriodoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/periodos/{id_periodo} [put]
func PutOnePeriodo(c *gin.Context) {
	id := c.Params.ByName("id")

	var input Request.PutOnePeriodo
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	f_input.PutOnePeriodo(&input)

	var periodo models.Periodo
	if err := repositories.GetOnePeriodo(&periodo, id); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	periodo = models.Periodo{
		Id: periodo.Id,
		// Id: utils.CheckNullString(input.Id, periodo.Id),
	}

	if err := repositories.PutOnePeriodo(&periodo, id); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	api_helpers.RespondJSON(c, 200, periodo)
}
