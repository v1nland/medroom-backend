package grupo

import (
	"errors"
	"medroom-backend/Utils"
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Obtiene un grupo de un administrador academico
// @Description Obtiene un grupo de un administrador academico según su token
// @Tags 05 - Administración Académica
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso a buscar"
// @Param   id_grupo     path    string     true        "Id del grupo a buscar"
// @Success 200 {object} Swagger.GetOneGrupoAdministradorAcademicoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-academica/me/cursos/{id_curso}/grupos/{id_grupo} [get]
func GetOneGrupoAdministradorAcademico(c *gin.Context) {
	id_grupo := c.Params.ByName("id_grupo")
	id_curso := c.Params.ByName("id_curso")
	id_administrador_academico := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_ACADEMICO")

	var grupo models.Grupo
	if err := repositories.GetOneGrupoAdministradorAcademico(&grupo, id_grupo, id_curso, id_administrador_academico); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Grupo not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	api_helpers.RespondJSON(c, 200, grupo)
}
