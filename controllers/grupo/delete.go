package grupo

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

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
func DeleteGrupo(c *gin.Context) {
	id := c.Params.ByName("id_grupo")

	var grupo models.Grupo
	if err := repositories.GetOneGrupo(&grupo, id); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	if err := repositories.DeleteGrupo(&grupo, id); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	api_helpers.RespondJSON(c, 200, grupo)
}
