package grupo

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type putRequest struct {
	Id_curso     *string `json:"id_curso"`
	Nombre_grupo *string `json:"nombre_grupo"`
	Sigla_grupo  *string `json:"sigla_grupo"`
}

func putRequestParse(u *putRequest) {
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

// @Summary Modifica un grupo
// @Description Modifica un grupo con los datos entregados
// @Tags 04 - Administración Académica
// @Accept  json
// @Produce  json
// @Param   id_grupo     path    string     true        "Id del grupo a modificar"
// @Param   input_actualiza_grupo     body    Request.Put     true        "Grupo a modificar"
// @Success 200 {object} Swagger.PutOneGrupoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-academica/grupos/{id_grupo} [put]
func Put(c *gin.Context) {
	id := c.Params.ByName("id_grupo")

	var input putRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	putRequestParse(&input)

	var grupo models.Grupo
	if err := repositories.GetOneGrupo(&grupo, id); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	grupo = models.Grupo{
		// Id:           grupo.Id,
		Sigla_curso:  utils.CheckNullString(input.Id_curso, grupo.Sigla_curso),
		Nombre_grupo: utils.CheckNullString(input.Nombre_grupo, grupo.Nombre_grupo),
		Sigla_grupo:  utils.CheckNullString(input.Sigla_grupo, grupo.Sigla_grupo),
	}

	if err := repositories.PutOneGrupo(&grupo, id); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJSON(c, 200, grupo)
}
