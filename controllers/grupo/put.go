package grupo

import (
	"medroom-backend/api_helpers"
	"medroom-backend/formats/f_input"
	"medroom-backend/messages/Request"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Modifica un grupo
// @Description Modifica un grupo con los datos entregados
// @Tags 05 - Administración Académica
// @Accept  json
// @Produce  json
// @Param   id_grupo     path    string     true        "Id del grupo a modificar"
// @Param   input_actualiza_grupo     body    Request.PutOneGrupo     true        "Grupo a modificar"
// @Success 200 {object} Swagger.PutOneGrupoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-academica/grupos/{id_grupo} [put]
func PutOneGrupo(c *gin.Context) {
	id := c.Params.ByName("id")

	var input Request.PutOneGrupo
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	f_input.PutOneGrupo(&input)

	var grupo models.Grupo
	if err := repositories.GetOneGrupo(&grupo, id); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	grupo = models.Grupo{
		Id:           grupo.Id,
		Id_curso:     utils.CheckNullInt(input.Id_curso, grupo.Id_curso),
		Nombre_grupo: utils.CheckNullString(input.Nombre_grupo, grupo.Nombre_grupo),
		Sigla_grupo:  utils.CheckNullString(input.Sigla_grupo, grupo.Sigla_grupo),
	}

	if err := repositories.PutOneGrupo(&grupo, id); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	api_helpers.RespondJSON(c, 200, grupo)
}
