package administrador_academico

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Lista de administradores-academicos
// @Description Lista todos los administradores-academicos
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Success 200 {array} Swagger.ListAdministradoresAcademicosSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-academicos [get]
func List(c *gin.Context) {
	var admins []models.AdministradorAcademico
	if err := repositories.ListAdministradoresAcademicos(&admins); err != nil {
		api_helpers.RespondJSON(c, 500, err.Error())
		return
	}

	api_helpers.RespondJSON(c, 200, admins)
}
