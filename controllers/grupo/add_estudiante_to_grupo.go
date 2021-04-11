package grupo

import (
	"errors"
	"medroom-backend/api_helpers"
	"medroom-backend/config"
	"medroom-backend/models"
	"medroom-backend/repositories"

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
// @Success 200 {object} Swagger.AddEstudianteToGrupoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-academica/cursos/{id_curso}/grupos/{id_grupo}/estudiantes/{uuid_estudiante} [put]
func AddEstudianteToGrupo(c *gin.Context) {
	id_curso := c.Params.ByName("id")
	id_grupo := c.Params.ByName("id_grupo")
	id_estudiante := c.Params.ByName("id_estudiante")

	var curso models.Curso
	if err := repositories.GetOneCurso(&curso, id_curso); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	var grupo models.Grupo
	if err := repositories.GetOneGrupo(&grupo, id_grupo); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	var estudiante models.Estudiante
	if err := repositories.GetOneEstudiante(&estudiante, id_estudiante); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	// validar que el estudiante no tenga un grupo ya en ese curso
	var grupos_este_curso []models.Grupo
	if err := repositories.GetGruposEstudiante(&grupos_este_curso, id_curso, id_estudiante); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 400, "Alumno no pertenece a este curso")
			return
		}
	}

	for i := 0; i < len(grupos_este_curso); i++ {
		if grupos_este_curso[i].Sigla_grupo != "SG" {
			api_helpers.RespondJSON(c, 400, "Alumno ya pertenece a un grupo de este curso")
			return
		}
	}
	// end

	config.DB.Model(&grupo).Association("Estudiantes_grupo").Append([]models.Estudiante{estudiante})

	// found, id_grupo_sg := utils.SearchIdGrupoBySigla(curso.Grupos_curso, "SG")
	// if found {
	// 	repositories.DeleteEstudianteGrupo(utils.ConvertIntToString(id_grupo_sg), id_estudiante)
	// }

	api_helpers.RespondJSON(c, 200, grupo)
}
