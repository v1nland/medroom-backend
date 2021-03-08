package curso

import (
	"errors"
	"medroom-backend/app/api_helpers"
	"medroom-backend/app/formats/f_input"
	"medroom-backend/app/messages/Request"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"
	"medroom-backend/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Modifica un curso
// @Description Modifica un curso con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso a modificar"
// @Param   input_actualiza_curso     body    Request.PutOneCurso     true        "Curso a modificar"
// @Success 200 {object} Swagger.PutOneCursoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/cursos/{id_curso} [put]
func PutOneCurso(c *gin.Context) {
	id := c.Params.ByName("id")

	var input Request.PutOneCurso
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	f_input.PutOneCurso(&input)

	var curso models.Curso
	if err := repositories.GetOneCurso(&curso, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Curso not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	// replace data in model entity
	curso = models.Curso{
		Id:           curso.Id,
		Id_periodo:   utils.CheckNullInt(input.Id_periodo, curso.Id_periodo),
		Nombre_curso: utils.CheckNullString(input.Nombre_curso, curso.Nombre_curso),
		Sigla_curso:  utils.CheckNullString(input.Sigla_curso, curso.Sigla_curso),
	}

	if err := repositories.PutOneCurso(&curso, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Curso not updated")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.PutOneCurso(model))
	api_helpers.RespondJSON(c, 200, curso)
}
