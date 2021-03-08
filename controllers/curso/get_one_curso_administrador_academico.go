package curso

import (
	"errors"
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Obtiene un curso de un administrador academico
// @Description Obtiene un curso de un administrador academico según su token
// @Tags 05 - Administración Académica
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso a buscar"
// @Success 200 {object} Swagger.GetOneCursoAdministradorAcademicoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-academica/me/cursos/{id_curso} [get]
func GetOneCursoAdministradorAcademico(c *gin.Context) {
	id_curso := c.Params.ByName("id")
	id_evaluador := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_ACADEMICO")

	var curso models.Curso
	if err := repositories.GetOneCursoAdministradorAcademico(&curso, id_curso, id_evaluador); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Curso not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.GetCursoEstudiante(cursos))
	api_helpers.RespondJSON(c, 200, curso)
}
