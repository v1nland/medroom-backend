package grupo

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type addRequest struct {
	Sigla_curso  *string `json:"sigla_curso"`
	Sigla_grupo  *string `json:"sigla_grupo"`
	Nombre_grupo *string `json:"nombre_grupo"`
}

func addRequestParse(u *addRequest) {
	if u.Nombre_grupo != nil {
		*u.Nombre_grupo = strings.TrimSpace(*u.Nombre_grupo)
		*u.Nombre_grupo = strings.ToUpper(*u.Nombre_grupo)
		*u.Nombre_grupo = utils.RemoveAccents(*u.Nombre_grupo)
	}
	if u.Sigla_grupo != nil {
		*u.Sigla_grupo = strings.TrimSpace(*u.Sigla_grupo)
		*u.Sigla_grupo = strings.ToUpper(*u.Sigla_grupo)
		*u.Sigla_grupo = utils.RemoveAccents(*u.Sigla_grupo)
	}
}

// @Summary Agrega un nuevo grupo
// @Description Genera un nuevo grupo con los datos entregados
// @Tags 04 - Administración Académica
// @Accept  json
// @Produce  json
// @Param   input_grupo     body    Request.Add     true        "Grupo a agregar"
// @Success 200 {object} Swagger.AddNewGrupoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-academica/grupos [post]
func Add(c *gin.Context) {
	id_curso := c.Params.ByName("id")

	var input addRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	addRequestParse(&input)

	grupo := models.Grupo{
		Sigla_curso: id_curso,
		// add idPeriodoCurso
		Nombre_grupo: *input.Nombre_grupo,
		Sigla_grupo:  *input.Sigla_grupo,
	}

	if err := repositories.AddNewGrupo(&grupo); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJSON(c, 200, grupo)
}
