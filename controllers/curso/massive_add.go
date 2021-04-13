package curso

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type massiveAddRequest struct {
	Cursos []struct {
		Id           *int    `json:"id"`
		Id_periodo   *int    `json:"id_periodo"`
		Nombre_curso *string `json:"nombre_curso"`
		Sigla_curso  *string `json:"sigla_curso"`
	} `json:"cursos"`
}

func massiveAddRequestParse(u *massiveAddRequest) {
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
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Param   input_curso     body    massiveAddRequest     true        "Curso a agregar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/cursos/carga-masiva [post]
func MassiveAdd(c *gin.Context) {
	var payload massiveAddRequest
	if err := c.ShouldBind(&payload); err != nil {
		api_helpers.RespondError(c, 400, err.Error())
		return
	}

	massiveAddRequestParse(&payload)

	var periodo models.Periodo
	if err := repositories.GetUltimoPeriodo(&periodo); err != nil {
		api_helpers.RespondError(c, 500, "ultimo periodo no encontrado")
		return
	}

	var cursos_error []int
	for i := 0; i < len(payload.Cursos); i++ {
		curso := models.Curso{
			Sigla_curso:  *payload.Cursos[i].Sigla_curso,
			Id_periodo:   periodo.Id,
			Nombre_curso: *payload.Cursos[i].Nombre_curso,
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
		api_helpers.RespondJson(c, 201, cursos_error)
	} else {
		api_helpers.RespondJson(c, 200, "ok")
	}
}
