package curso

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Modifica los cursos de un administrador_academico
// @Description Modifica los cursos de un administrador_academico con los datos entregados
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo"
// @Param   sigla_curso     path    string     true        "Sigla del curso a modificar"
// @Param   uuid_administrador_academico     path    string     true        "UUID del administrador_academico a asociar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/cursos/{id_periodo}/{sigla_curso}/administradores-academicos/{uuid_administrador_academico} [put]
func AddAdministradorAcademicoToCurso(c *gin.Context) {
	id_periodo := c.Params.ByName("id_periodo")
	sigla_curso := c.Params.ByName("sigla_curso")
	id_administrador_academico := c.Params.ByName("id_administrador_academico")

	var administrador_academico models.AdministradorAcademico
	if err := repositories.GetAdministradorAcademico(&administrador_academico, id_administrador_academico); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	var curso models.Curso
	if err := repositories.GetOneCurso(&curso, sigla_curso, id_periodo); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	curso.Administradores_academicos_curso = append(curso.Administradores_academicos_curso, administrador_academico)

	if err := repositories.PutOneCurso(&curso, sigla_curso, id_periodo); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, curso)
}
