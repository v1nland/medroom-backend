package curso

import (
	"errors"
	"medroom-backend/app/api_helpers"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Modifica los cursos de un administrador_academico
// @Description Modifica los cursos de un administrador_academico con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "ID del curso a modificar"
// @Param   uuid_administrador_academico     path    string     true        "UUID del administrador_academico a asociar"
// @Success 200 {object} Swagger.AddAdministradorAcademicoToCursoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/cursos/{id_curso}/administradores-academicos/{uuid_administrador_academico} [put]
func AddAdministradorAcademicoToCurso(c *gin.Context) {
	id_curso := c.Params.ByName("id")
	id_administrador_academico := c.Params.ByName("id_administrador_academico")

	var administrador_academico models.AdministradorAcademico
	if err := repositories.GetOneAdministradorAcademico(&administrador_academico, id_administrador_academico); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Administrador académico not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	var curso models.Curso
	if err := repositories.GetOneCurso(&curso, id_curso); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Curso not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	curso.Administradores_academicos_curso = append(curso.Administradores_academicos_curso, administrador_academico)

	if err := repositories.PutOneCurso(&curso, id_curso); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Curso not updated")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	api_helpers.RespondJSON(c, 200, curso)
}
