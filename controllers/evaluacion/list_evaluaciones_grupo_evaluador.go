package evaluacion

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Lista de evaluaciones de un grupo
// @Description Lista todas los evaluaciones disponibles de un evaluador de un grupo
// @Tags 03 - Evaluadores
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso"
// @Param   id_grupo     path    string     true        "Id del grupo"
// @Success 200 {array} Swagger.ListEvaluacionesGrupoEvaluadorSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /evaluadores/me/cursos/{id_curso}/grupos/{id_grupo}/evaluaciones [get]
func ListEvaluacionesGrupoEvaluador(c *gin.Context) {
	id_periodo := c.Params.ByName("id_periodo")
	sigla_grupo := c.Params.ByName("id_grupo")
	sigla_curso := c.Params.ByName("id_curso")
	id_evaluador := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")

	var grupo models.Grupo
	if err := repositories.GetOneGrupoEvaluador(&grupo, sigla_grupo, sigla_curso, id_periodo, id_evaluador); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.ListEvaluacionesEstudiante(grupo.Evaluaciones_grupo))
	api_helpers.RespondJSON(c, 200, grupo.Evaluaciones_grupo)
}
