package estudiante

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Lista de estudiantes de un curso
// @Description Lista todos los estudiantes existentes en un curso
// @Tags Administración Académica
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo"
// @Param   sigla_curso     path    string     true        "Sigla del curso"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-academica/cursos/{id_periodo}/{sigla_curso}/estudiantes [get]
func ListEstudiantesCurso(c *gin.Context) {
	// id := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ADMINISTRADOR_ACADEMICO")
	sigla_curso := c.Params.ByName("sigla_curso")

	var estudiantes []models.Estudiante
	if err := repositories.GetAllEstudiantesCurso(&estudiantes, sigla_curso); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJson(c, 200, estudiantes)
}
