package grupo

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Obtiene los grupos de un administrador academico
// @Description Obtiene los grupos de un administrador academico según su token
// @Tags 04 - Administración Académica
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso a buscar"
// @Success 200 {object} Swagger.GetGruposAdministradorAcademicoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-academica/me/cursos/{id_curso}/grupos [get]
func GetGruposAdministradorAcademico(c *gin.Context) {
	id_curso := c.Params.ByName("id_curso")
	id_administrador_academico := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_ACADEMICO")

	var grupos []models.Grupo
	if err := repositories.GetGruposAdministradorAcademico(&grupos, id_curso, id_administrador_academico); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJSON(c, 200, grupos)
}
