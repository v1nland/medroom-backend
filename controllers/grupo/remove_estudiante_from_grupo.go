package grupo

import (
	"errors"
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Modifica los grupos de un estudiante
// @Description Modifica los grupos de un estudiante con los datos entregados
// @Tags 04 - Administración Académica
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "ID del curso a modificar"
// @Param   id_grupo     path    string     true        "ID del grupo a modificar"
// @Param   uuid_estudiante     path    string     true        "UUID del estudiante a asociar"
// @Success 200 {object} Swagger.RemoveEstudianteFromGrupoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-academica/cursos/{id_curso}/grupos/{id_grupo}/estudiantes/{uuid_estudiante} [delete]
func RemoveEstudianteFromGrupo(c *gin.Context) {
	id_curso := c.Params.ByName("id")
	id_grupo := c.Params.ByName("id_grupo")
	id_estudiante := c.Params.ByName("id_estudiante")

	var curso models.Curso
	if err := repositories.GetOneCurso(&curso, id_curso); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Curso not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	var grupo models.Grupo
	if err := repositories.GetOneGrupo(&grupo, id_grupo); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Grupo not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	var estudiante models.Estudiante
	if err := repositories.GetOneEstudiante(&estudiante, id_estudiante); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Estudiante not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	repositories.DeleteEstudianteGrupo(utils.ConvertIntToString(utils.ConvertStringToInt(id_grupo)), id_estudiante)

	// found, id_grupo_sg := utils.SearchIdGrupoBySigla(curso.Grupos_curso, "SG")
	// if found {
	// 	var grupo_sg models.Grupo
	// 	if err := repositories.GetOneGrupo(&grupo_sg, utils.ConvertIntToString(id_grupo_sg)); err != nil {
	// 		if errors.Is(err, gorm.ErrRecordNotFound) {
	// 			api_helpers.RespondJSON(c, 200, "Grupo SG not found")
	// 		} else {
	// 			api_helpers.RespondError(c, 500, "default")
	// 		}

	// 		return
	// 	}

	// 	config.DB.Model(&grupo_sg).Association("Estudiantes_grupo").Append([]models.Estudiante{estudiante})
	// }

	api_helpers.RespondJSON(c, 200, grupo)
}
