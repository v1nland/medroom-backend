package administrador_academico

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Obtiene el perfil del administrador academico
// @Description Obtiene el perfil del administrador academico según su token
// @Tags Administración Académica
// @Accept  json
// @Produce  json
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-academica/me [get]
func Profile(c *gin.Context) {
	id_administrador_academico := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_ACADEMICO")

	var container models.AdministradorAcademico
	if err := repositories.GetAdministradorAcademico(&container, id_administrador_academico); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, container)
}
