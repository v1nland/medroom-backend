package curso

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type putRequest struct {
	Nombre_curso *string `json:"nombre_curso"`
}

func putRequestParse(u *putRequest) {
	if u.Nombre_curso != nil {
		*u.Nombre_curso = strings.TrimSpace(*u.Nombre_curso)
		*u.Nombre_curso = strings.ToUpper(*u.Nombre_curso)
		*u.Nombre_curso = utils.RemoveAccents(*u.Nombre_curso)
	}
}

// @Summary Modifica un curso
// @Description Modifica un curso con los datos entregados
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo"
// @Param   sigla_curso     path    string     true        "Sigla del curso a modificar"
// @Param   input_actualiza_curso     body    putRequest     true        "Curso a modificar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/cursos/{id_periodo}/{sigla_curso} [put]
func Put(c *gin.Context) {
	id_periodo := c.Params.ByName("id_periodo")
	sigla_curso := c.Params.ByName("sigla_curso")

	var input putRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	putRequestParse(&input)

	var curso models.Curso
	if err := repositories.GetOneCurso(&curso, sigla_curso, id_periodo); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	// replace data in model entity
	curso = models.Curso{
		Sigla_curso:  curso.Sigla_curso,
		Id_periodo:   curso.Id_periodo,
		Nombre_curso: utils.CheckNullString(input.Nombre_curso, curso.Nombre_curso),
	}

	if err := repositories.PutOneCurso(&curso, sigla_curso, id_periodo); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, curso)
}
