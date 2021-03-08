package grupo

import (
	"medroom-backend/app/api_helpers"
	"medroom-backend/app/formats/f_input"
	"medroom-backend/app/messages/Request"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Agrega un nuevo grupo
// @Description Genera un nuevo grupo con los datos entregados
// @Tags 05 - Administración Académica
// @Accept  json
// @Produce  json
// @Param   input_grupo     body    Request.AddNewGrupo     true        "Grupo a agregar"
// @Success 200 {object} Swagger.AddNewGrupoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-academica/grupos [post]
func AddNewGrupo(c *gin.Context) {
	var input Request.AddNewGrupo
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	f_input.AddNewGrupo(&input)

	grupo := models.Grupo{
		Id_curso:     *input.Id_curso,
		Nombre_grupo: *input.Nombre_grupo,
		Sigla_grupo:  *input.Sigla_grupo,
	}

	if err := repositories.AddNewGrupo(&grupo); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	api_helpers.RespondJSON(c, 200, grupo)
}
