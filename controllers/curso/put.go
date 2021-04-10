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
	Id_periodo   *string `json:"id_periodo"`
	Nombre_curso *string `json:"nombre_curso"`
	Sigla_curso  *string `json:"sigla_curso"`
}

func putRequestParse(u *putRequest) {
	if u.Nombre_curso != nil {
		*u.Nombre_curso = strings.TrimSpace(*u.Nombre_curso)
		*u.Nombre_curso = strings.ToUpper(*u.Nombre_curso)
		*u.Nombre_curso = utils.RemoveAccents(*u.Nombre_curso)
	}
	if u.Sigla_curso != nil {
		*u.Sigla_curso = strings.TrimSpace(*u.Sigla_curso)
		*u.Sigla_curso = strings.ToUpper(*u.Sigla_curso)
		*u.Sigla_curso = utils.RemoveAccents(*u.Sigla_curso)
	}
}

// @Summary Modifica un curso
// @Description Modifica un curso con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso a modificar"
// @Param   input_actualiza_curso     body    Request.Put     true        "Curso a modificar"
// @Success 200 {object} Swagger.PutOneCursoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/cursos/{id_curso} [put]
func Put(c *gin.Context) {
	id := c.Params.ByName("id")

	var input putRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	putRequestParse(&input)

	var curso models.Curso
	if err := repositories.GetOneCurso(&curso, id); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	// replace data in model entity
	curso = models.Curso{
		Sigla_curso:  utils.CheckNullString(input.Sigla_curso, curso.Sigla_curso),
		Id_periodo:   utils.CheckNullString(input.Id_periodo, curso.Id_periodo),
		Nombre_curso: utils.CheckNullString(input.Nombre_curso, curso.Nombre_curso),
	}

	if err := repositories.PutOneCurso(&curso, id); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJSON(c, 200, curso)
}
