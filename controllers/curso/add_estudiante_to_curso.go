package curso

import (
	"errors"
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Modifica los cursos de un estudiante
// @Description Modifica los cursos de un estudiante con los datos entregados
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo"
// @Param   sigla_curso     path    string     true        "Sigla del curso a modificar"
// @Param   uuid_estudiante     path    string     true        "UUID del estudiante a asociar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/cursos/{id_periodo}/{sigla_curso}/estudiantes/{uuid_estudiante} [put]
func AddEstudianteToCurso(c *gin.Context) {
	id_periodo := c.Params.ByName("id_periodo")
	sigla_curso := c.Params.ByName("sigla_curso")
	id_estudiante := c.Params.ByName("id_estudiante")

	var estudiante models.Estudiante
	if err := repositories.GetOneEstudiante(&estudiante, id_estudiante); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	var curso models.Curso
	if err := repositories.GetOneCurso(&curso, sigla_curso, id_periodo); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	// validar que el alumno pertenece al curso
	var grupos_este_curso []models.Grupo
	if err := repositories.GetGruposEstudiante(&grupos_este_curso, sigla_curso, id_periodo, id_estudiante); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJson(c, 400, "Alumno no pertenece a este curso")
			return
		}
	}

	// validar que el estudiante no tenga un grupo ya en ese curso
	for i := 0; i < len(grupos_este_curso); i++ {
		if grupos_este_curso[i].Sigla_grupo != "SG" {
			api_helpers.RespondJson(c, 400, "Alumno ya pertenece a un grupo de este curso")
			return
		}
	}

	// search grupo "sin grupo" in curso
	found, index := utils.SearchIndexGrupoBySigla(curso.Grupos_curso, "SG")
	if found {
		curso.Grupos_curso[index].Estudiantes_grupo = append(curso.Grupos_curso[index].Estudiantes_grupo, estudiante)
	} else {
		api_helpers.RespondJson(c, 200, "Grupo 'SIN GRUPO' not found")
		return
	}

	if err := repositories.PutOneCurso(&curso, sigla_curso, id_periodo); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, curso)
}
