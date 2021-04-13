package grupo

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Elimina un grupo
// @Description Elimina un grupo con los datos entregados
// @Tags Administración Académica
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo"
// @Param   sigla_curso     path    string     true        "Sigla del curso"
// @Param   sigla_grupo     path    string     true        "Sigla del grupo"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-academica/cursos/{id_periodo}/{sigla_curso}/grupos/{sigla_grupo} [delete]
func Delete(c *gin.Context) {
	id_periodo := c.Params.ByName("id_periodo")
	sigla_curso := c.Params.ByName("sigla_curso")
	sigla_grupo := c.Params.ByName("sigla_grupo")

	var grupo models.Grupo
	if err := repositories.GetOneGrupo(&grupo, sigla_curso, id_periodo, sigla_grupo); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	// clear grupo associations
	if err := repositories.ClearGrupo(sigla_grupo, sigla_curso, id_periodo); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	// clear evaluaciones grupo
	for _, evaluacion := range grupo.Evaluaciones_grupo {
		if err := repositories.DeleteEvaluacion(utils.IntToString(evaluacion.Id)); err != nil {
			api_helpers.RespondError(c, 500, err.Error())
			return
		}
	}

	// delete grupo
	repositories.DeleteGrupo(&grupo)

	// response
	api_helpers.RespondJson(c, 200, grupo)
}
