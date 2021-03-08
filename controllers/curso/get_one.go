package curso

import (
	"errors"
	"medroom-backend/api_helpers"
	"medroom-backend/formats/Output"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Obtiene un curso
// @Description Obtiene un curso según su Id
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso a buscar"
// @Success 200 {object} Swagger.GetOneCursoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/cursos/{id_curso} [get]
func GetOneCurso(c *gin.Context) {
	id := c.Params.ByName("id")

	var curso models.Curso
	if err := repositories.GetOneCurso(&curso, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Curso not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	api_helpers.RespondJSON(c, 200, Output.GetOneCursoOutput(curso))
}
