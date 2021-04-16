package grupo

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// to do: filtrar evaluaciones que estan rendidas unicamente

type estudianteGetOneGrupoEstudiante struct {
	Nombres_estudiante            string `json:"nombres_estudiante"`
	Apellidos_estudiante          string `json:"apellidos_estudiante"`
	Correo_electronico_estudiante string `json:"correo_electronico_estudiante" gorm:"unique;not null"`
}

type evaluacionGetOneGrupoEstudiante struct {
	Id                int       `json:"id"`
	Nombre_evaluacion string    `json:"nombre_evaluacion"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type getOneGrupoEstudianteResponse struct {
	Sigla_curso        string                            `json:"sigla_curso"`
	Id_periodo_curso   string                            `json:"id_periodo_curso"`
	Nombre_curso       string                            `json:"nombre_curso"`
	Sigla_grupo        string                            `json:"sigla_grupo"`
	Nombre_grupo       string                            `json:"nombre_grupo"`
	Evaluaciones_grupo []evaluacionGetOneGrupoEstudiante `json:"evaluaciones_grupo"` //id, nombre, created_at, updated_at
	Estudiantes_grupo  []estudianteGetOneGrupoEstudiante `json:"estudiantes_grupo"`  //id,nombre,apellido,correo
}

// @Summary Obtiene un grupo de un estudiante
// @Description Obtiene un grupo de un estudiante según su token
// @Tags Estudiantes
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo"
// @Param   sigla_curso     path    string     true        "Sigla del curso"
// @Param   sigla_grupo     path    string     true        "Sigla del grupo"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /estudiantes/me/cursos/{id_periodo}/{sigla_curso}/grupos/{sigla_grupo} [get]
func GetOneGrupoEstudiante(c *gin.Context) {
	id_periodo := c.Params.ByName("id_periodo")
	sigla_curso := c.Params.ByName("sigla_curso")
	sigla_grupo := c.Params.ByName("sigla_grupo")
	id_estudiante := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	var grupo models.Grupo
	if err := repositories.GetOneGrupoEstudiante(&grupo, sigla_grupo, sigla_curso, id_periodo, id_estudiante); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	var curso models.Curso
	if err := repositories.GetOneCurso(&curso, grupo.Sigla_curso, grupo.Id_periodo_curso); err != nil {
		api_helpers.RespondError(c, 500, err.Error())
		return
	}

	response := getOneGrupoEstudianteResponse{
		Sigla_curso:        grupo.Sigla_curso,
		Id_periodo_curso:   grupo.Id_periodo_curso,
		Nombre_curso:       curso.Nombre_curso,
		Sigla_grupo:        grupo.Sigla_grupo,
		Nombre_grupo:       grupo.Nombre_grupo,
		Evaluaciones_grupo: []evaluacionGetOneGrupoEstudiante{},
		Estudiantes_grupo:  []estudianteGetOneGrupoEstudiante{},
	}

	for _, evaluacion := range grupo.Evaluaciones_grupo {
		ev := evaluacionGetOneGrupoEstudiante{
			Id:                evaluacion.Id,
			Nombre_evaluacion: evaluacion.Nombre_evaluacion,
			CreatedAt:         evaluacion.CreatedAt,
			UpdatedAt:         evaluacion.UpdatedAt,
		}
		response.Evaluaciones_grupo = append(response.Evaluaciones_grupo, ev)
	}

	for _, estudiante := range grupo.Estudiantes_grupo {
		es := estudianteGetOneGrupoEstudiante{
			Nombres_estudiante:            estudiante.Nombres_estudiante,
			Apellidos_estudiante:          estudiante.Apellidos_estudiante,
			Correo_electronico_estudiante: estudiante.Correo_electronico_estudiante,
		}
		response.Estudiantes_grupo = append(response.Estudiantes_grupo, es)
	}

	api_helpers.RespondJson(c, 200, response)
}
