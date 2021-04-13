package grupo

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Obtiene un grupo de un evaluador
// @Description Obtiene un grupo de un evaluador según su token
// @Tags Evaluadores
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo"
// @Param   sigla_curso     path    string     true        "Sigla del curso"
// @Param   sigla_grupo     path    string     true        "Sigla del grupo"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /evaluadores/me/cursos/{id_periodo}/{sigla_curso}/grupos/{sigla_grupo} [get]
func GetOneGrupoEvaluador(c *gin.Context) {
	id_periodo := c.Params.ByName("id_periodo")
	sigla_curso := c.Params.ByName("sigla_curso")
	sigla_grupo := c.Params.ByName("sigla_grupo")
	id_evaluador := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")

	var grupo models.Grupo
	if err := repositories.GetOneGrupoEvaluador(&grupo, sigla_grupo, sigla_curso, id_periodo, id_evaluador); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, grupo)
}
