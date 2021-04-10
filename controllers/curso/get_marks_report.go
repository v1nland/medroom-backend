package curso

import (
	"errors"
	"fmt"
	"medroom-backend/api_helpers"
	"medroom-backend/models"
	"medroom-backend/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
// @Tags 05 - Administración Ti
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso a buscar"
// @Success 200 {object} Swagger.GetOneCursoSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /administracion-ti/cursos/{id_curso} [get]
func GetMarksReport(c *gin.Context) {
	id := c.Params.ByName("id")

	var curso models.Curso
	if err := repositories.GetOneCurso(&curso, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondJSON(c, 200, "Curso not found")
		} else {
			api_helpers.RespondError(c, 500, "default")
		}

		return
	}

	var reporte []EstudianteReport

	for _, grupo := range curso.Grupos_curso {
		var gp models.Grupo
		fmt.Println(grupo)

		// if err := repositories.GetOneGrupo(&gp, utils.ConvertIntToString(grupo.Id)); err != nil {
		// 	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 		api_helpers.RespondJSON(c, 200, "Grupo not found")
		// 	} else {
		// 		api_helpers.RespondError(c, 500, "default")
		// 	}
		// }

		for _, est := range gp.Estudiantes_grupo {
			var estudiante models.Estudiante

			if err := repositories.GetReporteEstudiante(&estudiante, est.Id.String()); err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					api_helpers.RespondJSON(c, 200, "Estudiante not found")
				} else {
					api_helpers.RespondError(c, 500, "default")
				}
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

	api_helpers.RespondJSON(c, 200, reporte)
}
