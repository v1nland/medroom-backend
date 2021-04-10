package administrador_academico

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Obtiene un administrador_academico
// @Description Obtiene un administrador_academico según su UUID
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   uuid_administrador_academico     path    string     true        "UUID del administrador_academico a buscar"
// @Success 200 {object} Swagger.GetOneAdministradorAcademicoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-academicos/{uuid_administrador_academico} [get]
func Get(c *gin.Context) {
	id := c.Params.ByName("id")

	var admin models.AdministradorAcademico
	if err := repositories.GetAdministradorAcademico(&admin, id); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJSON(c, 200, admin)
}
