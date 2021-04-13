package periodo

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Elimina un periodo
// @Description Elimina un periodo con los datos entregados
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo a eliminar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/periodos/{id_periodo} [delete]
func Delete(c *gin.Context) {
	id_periodo := c.Params.ByName("id_periodo")

	var periodo models.Periodo
	if err := repositories.GetOnePeriodo(&periodo, id_periodo); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	if err := repositories.DeletePeriodo(&periodo, id_periodo); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, periodo)
}
