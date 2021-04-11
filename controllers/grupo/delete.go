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
// @Tags 04 - Administración Académica
// @Accept  json
// @Produce  json
// @Param   id_grupo     path    string     true        "Id del grupo a eliminar"
// @Success 200 {object} Swagger.DeleteGrupoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-academica/grupos/{id_grupo} [delete]
func Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	var grupo models.Grupo
	if err := repositories.GetOneGrupo(&grupo, id); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	// clear grupo associations
	// if err := repositories.ClearGrupo(utils.ConvertIntToString(grupo.Id)); err != nil {
	// 	api_helpers.RespondError(c, 500, err.Error())
	// 	return
	// }

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
	api_helpers.RespondJSON(c, 200, grupo)
}
