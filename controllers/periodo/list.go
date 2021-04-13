package periodo

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Lista de periodos
// @Description Lista todos los periodos
// @Tags Rutas públicas
// @Accept  json
// @Produce  json
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /periodos [get]
func List(c *gin.Context) {
	var periodos []models.Periodo
	if err := repositories.GetAllPeriodos(&periodos); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, periodos)
}
