package rol

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Obtiene un rol
// @Description Obtiene un rol según su Id
// @Tags 00 - Rutas públicas
// @Accept  json
// @Produce  json
// @Param   id_rol     path    string     true        "Id del rol a buscar"
// @Success 200 {object} Swagger.GetOneRolSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /roles/{id_rol} [get]
func Get(c *gin.Context) {
	id := c.Params.ByName("id")

	var rol models.Rol
	if err := repositories.GetOneRol(&rol, id); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJSON(c, 200, rol)
}
