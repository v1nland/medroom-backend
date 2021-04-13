package administrador_ti

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Obtiene un administrador_ti
// @Description Obtiene un administrador_ti según su UUID
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_administrador_ti     path    string     true        "UUID del administrador_ti a buscar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/administradores-ti/{uuid_administrador_ti} [get]
func Get(c *gin.Context) {
	id_administrador_ti := c.Params.ByName("id_administrador_ti")

	var admin models.AdministradorTi
	if err := repositories.GetOneAdministradorTi(&admin, id_administrador_ti); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, admin)
}
