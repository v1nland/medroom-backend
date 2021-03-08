package periodo

import (
	"medroom-backend/api_helpers"
	"medroom-backend/formats/f_output"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Obtiene un periodo
// @Description Obtiene un periodo según su Id
// @Tags 00 - Rutas públicas
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo a buscar"
// @Success 200 {object} Swagger.GetOnePeriodoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /periodos/{id_periodo} [get]
func GetOnePeriodo(c *gin.Context) {
	id := c.Params.ByName("id")

	var periodo models.Periodo
	if err := repositories.GetOnePeriodo(&periodo, id); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	api_helpers.RespondJSON(c, 200, f_output.GetOnePeriodo(periodo))
}
