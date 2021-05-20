package curso

import (
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
)

type PuntajesEvaluacionEstudianteReport struct {
	Competencia  string `json:"competencia"`
	Calificacion int    `json:"calificacion"`
}

type EvaluacionEstudianteReport struct {
	Nombres_evaluador   string                               `json:"nombres_evaluador"`
	Apellidos_evaluador string                               `json:"apellidos_evaluador"`
	Nombre_evaluacion   string                               `json:"nombre_evaluacion"`
	Valoracion_general  int                                  `json:"valoracion_general"`
	Puntajes            []PuntajesEvaluacionEstudianteReport `json:"puntajes"`
}

type EstudianteReport struct {
	Curso        string                       `json:"curso"`
	Grupo        string                       `json:"grupo"`
	Rut          string                       `json:"rut"`
	Nombres      string                       `json:"nombres"`
	Apellidos    string                       `json:"apellidos"`
	Evaluaciones []EvaluacionEstudianteReport `json:"evaluaciones"`
}

// @Summary Obtiene un reporte de notas
// @Description Obtiene un reporte de notas según el id de curso
// @Tags Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo"
// @Param   sigla_curso     path    string     true        "Sigla del curso a buscar"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /administracion-ti/cursos/{id_periodo}/{sigla_curso}/reporte-notas [get]
func GetMarksReport(c *gin.Context) {
	id_periodo := c.Params.ByName("id_periodo")
	sigla_curso := c.Params.ByName("sigla_curso")

	var curso models.Curso
	if err := repositories.GetOneCurso(&curso, sigla_curso, id_periodo); err != nil {
		api_helpers.RespondError(c, 500, "default")
		return
	}

	var reporte []EstudianteReport

	for _, grupo := range curso.Grupos_curso {
		var gp models.Grupo

		if err := repositories.GetOneGrupo(&gp, grupo.Sigla_curso, grupo.Id_periodo_curso, grupo.Sigla_grupo); err != nil {
			api_helpers.RespondError(c, 500, err.Error())
		}

		for _, est := range gp.Estudiantes_grupo {
			var estudiante models.Estudiante

			if err := repositories.GetReporteEstudiante(&estudiante, est.Id.String()); err != nil {
				api_helpers.RespondError(c, 500, err.Error())
			}

			new_reporte := EstudianteReport{
				Curso:        curso.Nombre_curso,
				Grupo:        gp.Nombre_grupo,
				Rut:          estudiante.Rut_estudiante,
				Nombres:      estudiante.Nombres_estudiante,
				Apellidos:    estudiante.Apellidos_estudiante,
				Evaluaciones: []EvaluacionEstudianteReport{},
			}

			for _, calificacion := range estudiante.Calificaciones_estudiante {
				new_evaluacion := EvaluacionEstudianteReport{
					Nombres_evaluador:   calificacion.Evaluador_calificacion_estudiante.Nombres_evaluador,
					Apellidos_evaluador: calificacion.Evaluador_calificacion_estudiante.Apellidos_evaluador,
					Nombre_evaluacion:   calificacion.Evaluacion_calificacion_estudiante.Nombre_evaluacion,
					Valoracion_general:  calificacion.Valoracion_general_calificacion_estudiante,
					Puntajes:            []PuntajesEvaluacionEstudianteReport{},
				}

				for _, puntaje := range calificacion.Puntajes_calificacion_estudiante {
					new_puntaje := PuntajesEvaluacionEstudianteReport{
						Competencia:  puntaje.Competencia_puntaje.Nombre_competencia,
						Calificacion: puntaje.Calificacion_puntaje,
					}

					new_evaluacion.Puntajes = append(new_evaluacion.Puntajes, new_puntaje)
				}

				new_reporte.Evaluaciones = append(new_reporte.Evaluaciones, new_evaluacion)
			}

			reporte = append(reporte, new_reporte)
		}
	}

	api_helpers.RespondJson(c, 200, reporte)
}
