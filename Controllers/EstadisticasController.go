package Controllers

import (
	"errors"
	"medroom-backend/ApiHelpers"
	"medroom-backend/Messages/Query"
	"medroom-backend/Messages/Response"
	"medroom-backend/Repositories"
	"medroom-backend/Utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Evolución por competencia
// @Description Obtiene la evolución de un estudiante según competencia
// @Tags 02 - Estudiantes
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso"
// @Param   id_grupo     path    string     true        "Id del grupo"
// @Success 200 {array} Swagger.EvolucionEstudiantePorCompetenciaSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes/me/cursos/{id_curso}/grupos/{id_grupo}/estadisticas/evolucion-por-competencia [get]
func EvolucionEstudiantePorCompetencia(c *gin.Context) {
	// params
	id_estudiante := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")
	// id_curso := c.Params.ByName("id_curso")
	id_grupo := c.Params.ByName("id_grupo")

	var calificaciones_estudiante []Query.CalificacionesEstudiantePorCompetencia
	if err := Repositories.CalificacionesEstudiantePorCompetencia(&calificaciones_estudiante, id_grupo, id_estudiante); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondError(c, 500, "default")
			return
		}
	}

	response_container := &Response.EvolucionEstudiantePorCompetenciaResponse{
		Eje_x: Utils.BuildUniqueEvaluacionesCompetencia(calificaciones_estudiante),
		Eje_y: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		Valores: map[string][]struct {
			Puntaje_estudiante int     `json:"puntaje_estudiante"`
			Promedio_grupo     float64 `json:"promedio_grupo"`
		}{},
	}

	for i := 0; i < len(calificaciones_estudiante); i++ {
		response_container.Valores[calificaciones_estudiante[i].Id_competencia] = append(response_container.Valores[calificaciones_estudiante[i].Id_competencia], struct {
			Puntaje_estudiante int     `json:"puntaje_estudiante"`
			Promedio_grupo     float64 `json:"promedio_grupo"`
		}{
			calificaciones_estudiante[i].Calificacion_puntaje_estudiante,
			calificaciones_estudiante[i].Promedio_calificacion_puntaje_grupo,
		})
	}

	ApiHelpers.RespondJSON(c, 200, response_container)
}

// @Summary Evolución por evaluacion
// @Description Obtiene la evolución de un estudiante según evaluacion
// @Tags 02 - Estudiantes
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso"
// @Param   id_grupo     path    string     true        "Id del grupo"
// @Success 200 {array} Swagger.EvolucionEstudiantePorEvaluacionSwagger "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes/me/cursos/{id_curso}/grupos/{id_grupo}/estadisticas/evolucion-por-evaluacion [get]
func EvolucionEstudiantePorEvaluacion(c *gin.Context) {
	// params
	id_estudiante := Utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")
	// id_curso := c.Params.ByName("id_curso")
	id_grupo := c.Params.ByName("id_grupo")

	var calificaciones_estudiante []Query.CalificacionesEstudiantePorEvaluacion
	if err := Repositories.CalificacionesEstudiantePorEvaluacion(&calificaciones_estudiante, id_grupo, id_estudiante); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			ApiHelpers.RespondError(c, 500, "default")
			return
		}
	}

	response_container := &Response.EvolucionEstudiantePorEvaluacionResponse{
		Eje_x: Utils.BuildUniqueEvaluacionesEvaluacion(calificaciones_estudiante),
		Eje_y: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		Valores: map[string][]struct {
			Promedio_estudiante float64 `json:"promedio_estudiante"`
			Promedio_grupo      float64 `json:"promedio_grupo"`
		}{},
	}

	for i := 0; i < len(calificaciones_estudiante); i++ {
		response_container.Valores[calificaciones_estudiante[i].Nombre_evaluacion] = append(response_container.Valores[calificaciones_estudiante[i].Nombre_evaluacion], struct {
			Promedio_estudiante float64 `json:"promedio_estudiante"`
			Promedio_grupo      float64 `json:"promedio_grupo"`
		}{
			calificaciones_estudiante[i].Promedio_calificacion_puntaje_estudiante,
			calificaciones_estudiante[i].Promedio_calificacion_puntaje_grupo,
		})
	}

	ApiHelpers.RespondJSON(c, 200, response_container)
}
