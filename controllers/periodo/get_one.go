package periodo

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Obtiene un periodo
// @Description Obtiene un periodo según su Id
// @Tags Rutas públicas
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo a buscar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /periodos/{id_periodo} [get]
func Get(c *gin.Context) {
	id_periodo := c.Params.ByName("id_periodo")

	var periodo models.Periodo
	if err := repositories.GetOnePeriodo(&periodo, id_periodo); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, periodo)
}
