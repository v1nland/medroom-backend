package administrador_academico

import (
	"errors"
	"medroom-backend/app/api_helpers"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
func GetOneAdministradorAcademico(c *gin.Context) {
	id := c.Params.ByName("id")

	var container models.AdministradorAcademico
	if err := repositories.GetOneAdministradorAcademico(&container, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}
	}

	// api_helpers.RespondJSON(c, 200, f_output.GetOneAdministradorAcademico(container))
	api_helpers.RespondJSON(c, 200, container)
}
