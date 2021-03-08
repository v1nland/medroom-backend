package grupo

import (
	"errors"
	"medroom-backend/app/api_helpers"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"
	"medroom-backend/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Obtiene los grupos de un estudiante
// @Description Obtiene los grupos de un estudiante según su token
// @Tags 02 - Estudiantes
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso a buscar"
// @Success 200 {object} Swagger.GetGruposEstudianteSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /estudiantes/me/cursos/{id_curso}/grupos [get]
func GetGruposEstudiante(c *gin.Context) {
	id_curso := c.Params.ByName("id_curso")
	id_estudiante := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	var grupos []models.Grupo
	if err := repositories.GetGruposEstudiante(&grupos, id_curso, id_estudiante); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Grupo not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	api_helpers.RespondJSON(c, 200, grupos)
}
