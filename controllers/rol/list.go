package rol

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Lista de roles
// @Description Lista todos los roles
// @Tags 00 - Rutas públicas
// @Accept  json
// @Produce  json
// @Success 200 {array} Swagger.ListRolesSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /roles [get]
func List(c *gin.Context) {
	var roles []models.Rol

	if err := repositories.GetAllRoles(&roles); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJSON(c, 200, roles)
}
