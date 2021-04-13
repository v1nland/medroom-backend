package administrador_ti

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Lista de administradores-ti
// @Description Lista todos los administradores-ti
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/administradores-ti [get]
func List(c *gin.Context) {
	var admins []models.AdministradorTi
	if err := repositories.GetAllAdministradoresTi(&admins); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, admins)
}
