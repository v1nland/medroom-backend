package grupo

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Lista de grupos
// @Description Lista todos los grupos
// @Tags 04 - Administración Académica
// @Accept  json
// @Produce  json
// @Success 200 {array} Swagger.ListGruposSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-academica/grupos [get]
func ListGrupos(c *gin.Context) {
	var grupos []models.Grupo

	if err := repositories.GetAllGrupos(&grupos); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	api_helpers.RespondJSON(c, 200, grupos)
}
