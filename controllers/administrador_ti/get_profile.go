package administrador_ti

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Obtiene el perfil del administrador ti
// @Description Obtiene el perfil del administrador ti según su token
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Success 200 {object} Swagger.GetMyAdministradorTiSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/me [get]
func Profile(c *gin.Context) {
	id := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_TI")

	var admin models.AdministradorTi
	if err := repositories.GetOneAdministradorTi(&admin, id); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJSON(c, 200, admin)
}
