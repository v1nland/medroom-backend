package administrador_academico

import (
	"errors"
	"medroom-backend/app/api_helpers"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Lista de administradores-academicos
// @Description Lista todos los administradores-academicos
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Success 200 {array} Swagger.ListAdministradoresAcademicosSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/administradores-academicos [get]
func ListAdministradoresAcademicos(c *gin.Context) {
	var container []models.AdministradorAcademico

	if err := repositories.GetAllAdministradoresAcademicos(&container); err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			api_helpers.RespondJSON(c, 200, container)
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.ListAdministradoresAcademicos(container))
	api_helpers.RespondJSON(c, 200, container)
}
