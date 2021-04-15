package curso

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
)

type listCursosEstudianteResponse struct {
	Sigla_curso      string `json:"sigla_curso"`
	Id_periodo       string `json:"id_periodo"`
	Grupo_estudiante string `json:"grupo_estudiante"`
	Nombre_curso     string `json:"nombre_curso"`
	Estado_curso     bool   `json:"estado_curso"`
	// Administradores_academicos_curso []AdministradorAcademico `json:"administradores_academicos_curso"`
	// CreatedAt                        time.Time                `json:"created_at"`
	// UpdatedAt                        time.Time                `json:"updated_at"`
}

// @Summary Obtiene los cursos de un estudiante
// @Description Obtiene los cursos de un estudiante según su token
// @Tags Estudiantes
// @Accept  json
// @Produce  json
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /estudiantes/me/cursos [get]
func ListCursosEstudiante(c *gin.Context) {
	id_estudiante := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	var cursos []models.Curso
	if err := repositories.GetCursosEstudiante(&cursos, id_estudiante); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	var response []listCursosEstudianteResponse

	for _, curso := range cursos {
		new_curso := listCursosEstudianteResponse{
			Sigla_curso:  curso.Sigla_curso,
			Id_periodo:   curso.Id_periodo,
			Nombre_curso: curso.Nombre_curso,
			Estado_curso: curso.Estado_curso,
		}

		var grupos_estudiante []models.Grupo
		if err := repositories.GetGruposEstudiante(&grupos_estudiante, curso.Sigla_curso, curso.Id_periodo, id_estudiante); err != nil {
			api_helpers.RespondError(c, 500, err.Error())
			return
		}

		if len(grupos_estudiante) > 0 {
			new_curso.Grupo_estudiante = grupos_estudiante[0].Sigla_grupo
		} else {
			api_helpers.RespondError(c, 500, "Estudiante no pertenece a ningún grupo")
			return
		}

		response = append(response, new_curso)
	}

	api_helpers.RespondJson(c, 200, response)
}
