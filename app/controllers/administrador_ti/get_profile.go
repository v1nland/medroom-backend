package administrador_ti

import (
	"medroom-backend/app/api_helpers"
	"medroom-backend/app/formats/f_output"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"
	"medroom-backend/app/utils"

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
func GetMyAdministradorTi(c *gin.Context) {
	id := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_TI")

	var administrador_ti models.AdministradorTi
	if err := repositories.GetOneAdministradorTi(&administrador_ti, id); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	api_helpers.RespondJSON(c, 200, f_output.GetMyAdministradorTi(administrador_ti))
}
