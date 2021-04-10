package curso

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type addRequest struct {
	Nombre_curso *string `json:"nombre_curso"`
	Sigla_curso  *string `json:"sigla_curso"`
}

func addRequestParse(u *addRequest) {
	if u.Nombre_curso != nil {
		*u.Nombre_curso = strings.TrimSpace(*u.Nombre_curso)
		*u.Nombre_curso = strings.ToUpper(*u.Nombre_curso)
		*u.Nombre_curso = utils.RemoveAccents(*u.Nombre_curso)
	}
	if u.Sigla_curso != nil {
		*u.Sigla_curso = strings.TrimSpace(*u.Sigla_curso)
		*u.Sigla_curso = strings.ToUpper(*u.Sigla_curso)
		*u.Sigla_curso = utils.RemoveAccents(*u.Sigla_curso)
	}
}

// @Summary Agrega un nuevo curso
// @Description Genera un nuevo curso con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_curso     body    Request.Add     true        "Curso a agregar"
// @Success 200 {object} Swagger.AddNewCursoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/cursos [post]
func Add(c *gin.Context) {
	var input addRequest
	if err := c.ShouldBind(&input); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	addRequestParse(&input)

	var periodo models.Periodo
	if err := repositories.GetUltimoPeriodo(&periodo); err != nil {
		api_helpers.RespondError(c, 500, "Last period not found")
		return
	}

	curso := models.Curso{
		Id_periodo:   periodo.Id,
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
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	api_helpers.RespondJSON(c, 200, curso)
}
