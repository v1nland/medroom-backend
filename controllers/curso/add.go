package curso

import (
	"medroom-backend/Messages/Request"
	"medroom-backend/api_helpers"
	"medroom-backend/formats/f_input"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

// @Summary Agrega un nuevo curso
// @Description Genera un nuevo curso con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_curso     body    Request.AddNewCurso     true        "Curso a agregar"
// @Success 200 {object} Swagger.AddNewCursoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/cursos [post]
func AddNewCurso(c *gin.Context) {
	var input Request.AddNewCurso
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, "default")
		return
	}

	f_input.AddNewCurso(&input)

	curso := models.Curso{
		Id_periodo:   *input.Id_periodo,
		Nombre_curso: *input.Nombre_curso,
		Sigla_curso:  *input.Sigla_curso,
		Estado_curso: true,
		Grupos_curso: []models.Grupo{
			{
				Evaluadores_grupo: []models.Evaluador{},
				Estudiantes_grupo: []models.Estudiante{},
				Nombre_grupo:      "SIN GRUPO",
				Sigla_grupo:       "SG",
			},
		},
	}

	if err := repositories.AddNewCurso(&curso); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	// api_helpers.RespondJSON(c, 200, f_output.AddNewCurso(model_container))
	api_helpers.RespondJSON(c, 200, curso)
}
