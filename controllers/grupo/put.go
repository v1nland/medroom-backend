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
// @Tags Administración Académica
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo"
// @Param   sigla_curso     path    string     true        "Sigla del curso"
// @Param   sigla_grupo     path    string     true        "Sigla del grupo"
// @Param   input_actualiza_grupo     body    putRequest     true        "Grupo a modificar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-academica/cursos/{id_periodo}/{sigla_curso}/grupos/{sigla_grupo} [put]
func Put(c *gin.Context) {
	id_periodo := c.Params.ByName("id_periodo")
	sigla_curso := c.Params.ByName("sigla_curso")
	sigla_grupo := c.Params.ByName("sigla_grupo")

	var input putRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	putRequestParse(&input)

	var grupo models.Grupo
	if err := repositories.GetOneGrupo(&grupo, sigla_curso, id_periodo, sigla_grupo); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	grupo = models.Grupo{
		Id_periodo_curso:   grupo.Id_periodo_curso,
		Evaluaciones_grupo: grupo.Evaluaciones_grupo,
		Evaluadores_grupo:  grupo.Evaluadores_grupo,
		Estudiantes_grupo:  grupo.Estudiantes_grupo,
		Sigla_curso:        utils.CheckNullString(input.Id_curso, grupo.Sigla_curso),
		Nombre_grupo:       utils.CheckNullString(input.Nombre_grupo, grupo.Nombre_grupo),
		Sigla_grupo:        utils.CheckNullString(input.Sigla_grupo, grupo.Sigla_grupo),
	}

	if err := repositories.PutOneGrupo(&grupo, sigla_grupo, sigla_curso, id_periodo); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, grupo)
}
