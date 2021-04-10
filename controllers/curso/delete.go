package curso

import (
	"errors"
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

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
func Delete(c *gin.Context) {
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

	// clear grupo associations
	// for _, gp := range curso.Grupos_curso {
	// 	var grupo models.Grupo
	// 	if err := repositories.GetOneGrupo(&grupo, utils.ConvertIntToString(gp.Id)); err != nil {
	// 		if errors.Is(err, gorm.ErrRecordNotFound) {
	// 			api_helpers.RespondJSON(c, 200, "Curso not found")
	// 		} else {
	// 			api_helpers.RespondError(c, 500, "default")
	// 		}

	// 		return
	// 	}

	// 	if err := repositories.ClearGrupo(utils.ConvertIntToString(grupo.Id)); err != nil {
	// 		api_helpers.RespondError(c, 500, "default")
	// 		return
	// 	}

	// 	for _, evaluacion := range grupo.Evaluaciones_grupo {
	// 		if err := repositories.DeleteEvaluacion(utils.ConvertIntToString(evaluacion.Id)); err != nil {
	// 			api_helpers.RespondError(c, 500, "default")
	// 			return
	// 		}
	// 	}
	// }

	// eliminar grupos del curso
	for _, grupo := range curso.Grupos_curso {
		repositories.DeleteGrupo(&grupo)
	}

	// eliminar curso
	if err := repositories.ClearCurso(&curso, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Curso not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	api_helpers.RespondJSON(c, 200, curso)
}
