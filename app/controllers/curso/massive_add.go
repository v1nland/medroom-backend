package curso

import (
	"medroom-backend/app/api_helpers"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"
	"medroom-backend/app/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type massive_add_input struct {
	Cursos []struct {
		Id           *int    `json:"id"`
		Id_periodo   *int    `json:"id_periodo"`
		Nombre_curso *string `json:"nombre_curso"`
		Sigla_curso  *string `json:"sigla_curso"`
	} `json:"cursos"`
}

func massive_add_format(u *massive_add_input) {
	for i := 0; i < len(u.Cursos); i++ {
		if u.Cursos[i].Nombre_curso != nil {
			*u.Cursos[i].Nombre_curso = strings.TrimSpace(*u.Cursos[i].Nombre_curso)
			*u.Cursos[i].Nombre_curso = strings.ToUpper(*u.Cursos[i].Nombre_curso)
			*u.Cursos[i].Nombre_curso = utils.RemoveAccents(*u.Cursos[i].Nombre_curso)
		}

		if u.Cursos[i].Sigla_curso != nil {
			*u.Cursos[i].Sigla_curso = strings.TrimSpace(*u.Cursos[i].Sigla_curso)
			*u.Cursos[i].Sigla_curso = strings.ToUpper(*u.Cursos[i].Sigla_curso)
			*u.Cursos[i].Sigla_curso = utils.RemoveAccents(*u.Cursos[i].Sigla_curso)
		}
	}
}

// @Summary Agrega nuevos cursos de forma masiva
// @Description Genera nuevos cursos con los datos entregados
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_curso     body    massive_add_input     true        "Curso a agregar"
// @Success 200 {object} Swagger.AddNewCursoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/cursos/carga-masiva [post]
func AddNewCursos(c *gin.Context) {
	var payload massive_add_input
	if err := c.ShouldBind(&payload); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	massive_add_format(&payload)

	var cursos_error []int
	for i := 0; i < len(payload.Cursos); i++ {
		curso := models.Curso{
			Id:           *payload.Cursos[i].Id,
			Id_periodo:   *payload.Cursos[i].Id_periodo,
			Nombre_curso: *payload.Cursos[i].Nombre_curso,
			Sigla_curso:  *payload.Cursos[i].Sigla_curso,
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
			cursos_error = append(cursos_error, *payload.Cursos[i].Id)
		}
	}

	if len(cursos_error) > 0 {
		api_helpers.RespondJSON(c, 201, cursos_error)
	} else {
		api_helpers.RespondJSON(c, 200, "ok")
	}
}
