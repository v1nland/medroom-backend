package curso

import (
	"errors"
	"medroom-backend/app/api_helpers"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"
	"medroom-backend/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Elimina un curso
// @Description Elimina un curso con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso a eliminar"
// @Success 200 {object} Swagger.DeleteCursoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/cursos/{id_curso} [delete]
func DeleteCurso(c *gin.Context) {
	id := c.Params.ByName("id")

	var curso models.Curso
	if err := repositories.GetOneCurso(&curso, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Curso not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	// disociar estudiantes
	for i := 0; i < len(curso.Grupos_curso); i++ {
		for j := 0; j < len(curso.Grupos_curso[i].Estudiantes_grupo); j++ {
			repositories.DeleteEstudianteGrupo(utils.ConvertIntToString(curso.Grupos_curso[i].Id), curso.Grupos_curso[i].Estudiantes_grupo[j].Id.String())
		}
	}

	// disociar evaluadores
	for i := 0; i < len(curso.Grupos_curso); i++ {
		for j := 0; j < len(curso.Grupos_curso[i].Evaluadores_grupo); j++ {
			repositories.DeleteEvaluadorGrupo(utils.ConvertIntToString(curso.Grupos_curso[i].Id), curso.Grupos_curso[i].Evaluadores_grupo[j].Id.String())
		}
	}

	if err := repositories.DeleteCurso(&curso, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Curso not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.DeleteCurso(container))
	api_helpers.RespondJSON(c, 200, curso)
}
