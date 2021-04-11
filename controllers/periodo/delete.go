package periodo

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Elimina un periodo
// @Description Elimina un periodo con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo a eliminar"
// @Success 200 {object} Swagger.DeletePeriodoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/periodos/{id_periodo} [delete]
func Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	var periodo models.Periodo
	if err := repositories.GetOnePeriodo(&periodo, id); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	if err := repositories.DeletePeriodo(&periodo, id); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJSON(c, 200, periodo)
}
