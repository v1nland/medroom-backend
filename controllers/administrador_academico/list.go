package administrador_academico

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Lista de administradores-academicos
// @Description Lista todos los administradores-academicos
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/administradores-academicos [get]
func List(c *gin.Context) {
	var admins []models.AdministradorAcademico
	if err := repositories.ListAdministradoresAcademicos(&admins); err != nil {
		api_helpers.RespondJson(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, admins)
}
