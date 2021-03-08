package administrador_ti

import (
	"medroom-backend/app/api_helpers"
	"medroom-backend/app/formats/f_output"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Lista de administradores-ti
// @Description Lista todos los administradores-ti
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Success 200 {array} Swagger.ListAdministradoresTiSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-ti [get]
func ListAdministradoresTi(c *gin.Context) {
	var administradores_ti []models.AdministradorTi
	if err := repositories.GetAllAdministradoresTi(&administradores_ti); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	api_helpers.RespondJSON(c, 200, f_output.ListAdministradoresTi(administradores_ti))
}
