package grupo

import (
	"medroom-backend/api_helpers"
	"medroom-backend/formats/f_input"
	"medroom-backend/messages/Request"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Agrega un nuevo grupo
// @Description Genera un nuevo grupo con los datos entregados
// @Tags 04 - Administración Académica
// @Accept  json
// @Produce  json
// @Param   input_grupo     body    Request.AddNewGrupo     true        "Grupo a agregar"
// @Success 200 {object} Swagger.AddNewGrupoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-academica/grupos [post]
func AddNewGrupo(c *gin.Context) {
	id_curso := c.Params.ByName("id")

	var input Request.AddNewGrupo
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	f_input.AddNewGrupo(&input)

	grupo := models.Grupo{

		Sigla_curso: id_curso,
		// add idPeriodoCurso
		Nombre_grupo: *input.Nombre_grupo,
		Sigla_grupo:  *input.Sigla_grupo,
	}

	if err := repositories.AddNewGrupo(&grupo); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	api_helpers.RespondJSON(c, 200, grupo)
}
