package administrador_ti

import (
	"medroom-backend/Utils"
	"medroom-backend/api_helpers"
	"medroom-backend/formats/Output"
	"medroom-backend/models"
	"medroom-backend/repositories"

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
	id := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_TI")

	var administrador_ti models.AdministradorTi
	if err := repositories.GetOneAdministradorTi(&administrador_ti, id); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	api_helpers.RespondJSON(c, 200, Output.GetMyAdministradorTiOutput(administrador_ti))
}
