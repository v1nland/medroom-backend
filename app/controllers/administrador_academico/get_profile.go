package administrador_academico

import (
	"medroom-backend/app/api_helpers"
	"medroom-backend/app/formats/f_output"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"
	"medroom-backend/app/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Obtiene el perfil del administrador academico
// @Description Obtiene el perfil del administrador academico según su token
// @Tags 04 - Administración Academica
// @Accept  json
// @Produce  json
// @Success 200 {object} Swagger.GetMyAdministradorAcademicoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-academica/me [get]
func GetMyAdministradorAcademico(c *gin.Context) {
	// params
	id_administrador_academico := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_ACADEMICO")

	// model container
	var container models.AdministradorAcademico

	// query
	err := repositories.GetOneAdministradorAcademico(&container, id_administrador_academico)
	if err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	// output
	api_helpers.RespondJSON(c, 200, f_output.GetMyAdministradorAcademico(container))
}
