package curso

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Obtiene los cursos de un administrador academico
// @Description Obtiene los cursos de un administrador academico según su token
// @Tags Administración Académica
// @Accept  json
// @Produce  json
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-academica/me/cursos [get]
func GetCursosAdministradorAcademico(c *gin.Context) {
	id_evaluador := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_ACADEMICO")

	var cursos []models.Curso
	if err := repositories.GetCursosAdministradorAcademico(&cursos, id_evaluador); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, cursos)
}
