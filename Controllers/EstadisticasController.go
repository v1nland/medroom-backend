package Controllers

import (
	"medroom-backend/ApiHelpers"
	"medroom-backend/Messages/Response"
	"medroom-backend/Models"
	"medroom-backend/Repositories"
	"medroom-backend/Utils"

	"github.com/gin-gonic/gin"
)

// @Summary Evolución por evaluación
// @Description Obtiene la evolución de un estudiante según una evaluación
// @Tags 02 - Estudiantes
// @Accept  json
// @Produce  json
// @Param   id_evaluacion     path    string     true        "Id de la evaluación"
// @Success 200 {array} SwaggerMessages.EvolucionEstudiantePorEvaluacionSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes/me/estadisticas/evolucion/evaluacion/{id_evaluacion} [get]
func EvolucionEstudiantePorEvaluacion(c *gin.Context) {
	// params
	id_evaluacion := c.Params.ByName("id_evaluacion")
	id_estudiante := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	// search evaluacion
	var evaluacion Models.Evaluacion
	if err := Repositories.GetOneEvaluacion(&evaluacion, id_evaluacion); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// model estudiante
	var estudiante Models.Estudiante
	if err := Repositories.GetOneEstudiante(&estudiante, id_estudiante); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// model grupo estudiante
	// var grupo_estudiante Models.Grupo
	// if err := Repositories.GetOneGrupo(&grupo_estudiante, Utils.ConvertIntToString(estudiante.Id_grupo)); err != nil {
	// 	ApiHelpers.RespondError(c, 500, "default")
	// 	return
	// }

	// model curso estudiante
	// var curso_estudiante Models.Curso
	// if err := Repositories.GetOneCurso(&curso_estudiante, Utils.ConvertIntToString(grupo_estudiante.Id_curso)); err != nil {
	// 	ApiHelpers.RespondError(c, 500, "default")
	// 	return
	// }

	// calculate promedio estudiante
	suma_notas_estudiante := 0.0
	contador_estudiante := 0.0
	for k := 0; k < len(estudiante.Evaluaciones_estudiante); k++ {
		for m := 0; m < len(estudiante.Evaluaciones_estudiante[k].Puntajes_evaluacion); m++ {
			if estudiante.Evaluaciones_estudiante[k].Nombre_evaluacion == evaluacion.Nombre_evaluacion {
				suma_notas_estudiante += float64(estudiante.Evaluaciones_estudiante[k].Puntajes_evaluacion[m].Calificacion_puntaje)
				contador_estudiante++
			}
		}
	}

	// // calculate promedio grupo
	// suma_notas_grupo := 0.0
	// contador_grupo := 0.0
	// for j := 0; j < len(grupo_estudiante.Estudiantes_grupo); j++ {
	// 	for k := 0; k < len(grupo_estudiante.Estudiantes_grupo[j].Evaluaciones_estudiante); k++ {
	// 		for m := 0; m < len(grupo_estudiante.Estudiantes_grupo[j].Evaluaciones_estudiante[k].Puntajes_evaluacion); m++ {
	// 			if grupo_estudiante.Estudiantes_grupo[j].Evaluaciones_estudiante[k].Nombre_evaluacion == evaluacion.Nombre_evaluacion {
	// 				suma_notas_grupo += float64(grupo_estudiante.Estudiantes_grupo[j].Evaluaciones_estudiante[k].Puntajes_evaluacion[m].Calificacion_puntaje)
	// 				contador_grupo++
	// 			}
	// 		}
	// 	}
	// }

	// // calculate promedio curso
	// suma_notas_curso := 0.0
	// contador_curso := 0.0
	// for i := 0; i < len(curso_estudiante.Grupos_curso); i++ {
	// 	for j := 0; j < len(curso_estudiante.Grupos_curso[i].Estudiantes_grupo); j++ {
	// 		for k := 0; k < len(curso_estudiante.Grupos_curso[i].Estudiantes_grupo[j].Evaluaciones_estudiante); k++ {
	// 			for m := 0; m < len(curso_estudiante.Grupos_curso[i].Estudiantes_grupo[j].Evaluaciones_estudiante[k].Puntajes_evaluacion); m++ {
	// 				if curso_estudiante.Grupos_curso[i].Estudiantes_grupo[j].Evaluaciones_estudiante[k].Nombre_evaluacion == evaluacion.Nombre_evaluacion {
	// 					suma_notas_curso += float64(curso_estudiante.Grupos_curso[i].Estudiantes_grupo[j].Evaluaciones_estudiante[k].Puntajes_evaluacion[m].Calificacion_puntaje)
	// 					contador_curso++
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	response := &Response.EvolucionEstudiantePorEvaluacionResponse{
		// Nombre_evaluacion:   evaluacion.Nombre_evaluacion,
		// Promedio_estudiante: float64(suma_notas_estudiante / contador_estudiante),
		// Promedio_grupo:      float64(suma_notas_grupo / contador_grupo),
		// Promedio_curso:      float64(suma_notas_curso / contador_curso),
	}

	// output
	ApiHelpers.RespondJSON(c, 200, response)
}

// @Summary Evolución por competencia
// @Description Obtiene la evolución de un estudiante según una competencia
// @Tags 02 - Estudiantes
// @Accept  json
// @Produce  json
// @Param   codigo_competencia     path    string     true        "Código de la competencia"
// @Success 200 {array} SwaggerMessages.EvolucionEstudiantePorCompetenciaSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes/me/estadisticas/evolucion/competencia/{codigo_competencia} [get]
func EvolucionEstudiantePorCompetencia(c *gin.Context) {
	// params
	codigo_competencia := c.Params.ByName("codigo_competencia")
	id_estudiante := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")

	// model estudiante
	var estudiante Models.Estudiante
	if err := Repositories.GetOneEstudiante(&estudiante, id_estudiante); err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// model grupo estudiante
	// var grupo_estudiante Models.Grupo
	// if err := Repositories.GetOneGrupo(&grupo_estudiante, Utils.ConvertIntToString(estudiante.Id_grupo)); err != nil {
	// 	ApiHelpers.RespondError(c, 500, "default")
	// 	return
	// }

	// // model curso estudiante
	// var curso_estudiante Models.Curso
	// if err := Repositories.GetOneCurso(&curso_estudiante, Utils.ConvertIntToString(grupo_estudiante.Id_curso)); err != nil {
	// 	ApiHelpers.RespondError(c, 500, "default")
	// 	return
	// }

	// calculate promedio estudiante
	suma_notas_estudiante := 0.0
	contador_estudiante := 0.0
	for k := 0; k < len(estudiante.Evaluaciones_estudiante); k++ {
		for m := 0; m < len(estudiante.Evaluaciones_estudiante[k].Puntajes_evaluacion); m++ {
			if estudiante.Evaluaciones_estudiante[k].Puntajes_evaluacion[m].Codigo_competencia_puntaje == codigo_competencia {
				suma_notas_estudiante += float64(estudiante.Evaluaciones_estudiante[k].Puntajes_evaluacion[m].Calificacion_puntaje)
				contador_estudiante++
			}
		}
	}

	// // calculate promedio grupo
	// suma_notas_grupo := 0.0
	// contador_grupo := 0.0
	// for j := 0; j < len(grupo_estudiante.Estudiantes_grupo); j++ {
	// 	for k := 0; k < len(grupo_estudiante.Estudiantes_grupo[j].Evaluaciones_estudiante); k++ {
	// 		for m := 0; m < len(grupo_estudiante.Estudiantes_grupo[j].Evaluaciones_estudiante[k].Puntajes_evaluacion); m++ {
	// 			if grupo_estudiante.Estudiantes_grupo[j].Evaluaciones_estudiante[k].Puntajes_evaluacion[m].Codigo_competencia_puntaje == codigo_competencia {
	// 				suma_notas_grupo += float64(grupo_estudiante.Estudiantes_grupo[j].Evaluaciones_estudiante[k].Puntajes_evaluacion[m].Calificacion_puntaje)
	// 				contador_grupo++
	// 			}
	// 		}
	// 	}
	// }

	// // calculate promedio curso
	// suma_notas_curso := 0.0
	// contador_curso := 0.0
	// for i := 0; i < len(curso_estudiante.Grupos_curso); i++ {
	// 	for j := 0; j < len(curso_estudiante.Grupos_curso[i].Estudiantes_grupo); j++ {
	// 		for k := 0; k < len(curso_estudiante.Grupos_curso[i].Estudiantes_grupo[j].Evaluaciones_estudiante); k++ {
	// 			for m := 0; m < len(curso_estudiante.Grupos_curso[i].Estudiantes_grupo[j].Evaluaciones_estudiante[k].Puntajes_evaluacion); m++ {
	// 				if curso_estudiante.Grupos_curso[i].Estudiantes_grupo[j].Evaluaciones_estudiante[k].Puntajes_evaluacion[m].Codigo_competencia_puntaje == codigo_competencia {
	// 					suma_notas_curso += float64(curso_estudiante.Grupos_curso[i].Estudiantes_grupo[j].Evaluaciones_estudiante[k].Puntajes_evaluacion[m].Calificacion_puntaje)
	// 					contador_curso++
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	response := &Response.EvolucionEstudiantePorCompetenciaResponse{
		// Nombre_competencia:  codigo_competencia,
		// Promedio_estudiante: float64(suma_notas_estudiante / contador_estudiante),
		// Promedio_grupo:      float64(suma_notas_grupo / contador_grupo),
		// Promedio_curso:      float64(suma_notas_curso / contador_curso),
	}

	// output
	ApiHelpers.RespondJSON(c, 200, response)
}
