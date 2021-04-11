package periodo

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Lista de periodos
// @Description Lista todos los periodos
// @Tags 00 - Rutas públicas
// @Accept  json
// @Produce  json
// @Success 200 {array} Swagger.ListPeriodosSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /periodos [get]
func List(c *gin.Context) {
	var periodos []models.Periodo

	if err := repositories.GetAllPeriodos(&periodos); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJSON(c, 200, periodos)
}
