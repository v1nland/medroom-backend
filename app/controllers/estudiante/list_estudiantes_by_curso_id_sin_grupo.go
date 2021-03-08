package estudiante

import (
	"errors"
	"medroom-backend/app/api_helpers"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Lista de estudiantes de un curso sin grupo
// @Description Lista todos los estudiantes existentes en un curso sin grupo
// @Tags 05 - Administración Académica
// @Accept  json
// @Produce  json
// @Success 200 {array} Swagger.ListEstudiantesCursoSinGrupoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-academica/cursos/:id_curso/estudiantes/sin-grupo [get]
func ListEstudiantesCursoSinGrupo(c *gin.Context) {
	// id := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_ACADEMICO")
	id_curso := c.Params.ByName("id")

	var estudiantes []models.Estudiante
	if err := repositories.GetAllEstudiantesCursoSinGrupo(&estudiantes, id_curso); err != nil {
		if errors.Is(err, gorm.ErrEmptySlice) {
			api_helpers.RespondJSON(c, 200, estudiantes)
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	api_helpers.RespondJSON(c, 200, estudiantes)
}
