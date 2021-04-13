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
// @Tags Administración Académica
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo"
// @Param   sigla_curso     path    string     true        "Sigla del curso"
// @Param   input_grupo     body    addRequest     true        "Grupo a agregar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-academica/cursos/{id_periodo}/{sigla_curso}/grupos [post]
func Add(c *gin.Context) {
	id_periodo := c.Params.ByName("id_periodo")
	sigla_curso := c.Params.ByName("sigla_curso")

	var input addRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	addRequestParse(&input)

	grupo := models.Grupo{
		Id_periodo_curso: id_periodo,
		Sigla_curso:      sigla_curso,
		Nombre_grupo:     *input.Nombre_grupo,
		Sigla_grupo:      *input.Sigla_grupo,
	}

	if err := repositories.AddNewGrupo(&grupo); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, grupo)
}
