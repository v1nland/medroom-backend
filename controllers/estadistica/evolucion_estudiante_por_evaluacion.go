package estadistica

import (
	"errors"
	"medroom-backend/api_helpers"
	"medroom-backend/messages/Query"
	"medroom-backend/repositories"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type evolucionEstudiantePorEvaluacionResponse struct {
	Eje_x   []string `json:"eje_x"`
	Eje_y   []int    `json:"eje_y"`
	Valores map[string][]struct {
		Promedio_estudiante float64 `json:"promedio_estudiante"`
		Promedio_grupo      float64 `json:"promedio_grupo"`
	} `json:"valores"`
}

// @Summary Evolución por evaluacion
// @Description Obtiene la evolución de un estudiante según evaluacion
// @Tags Estudiantes
// @Accept  json
// @Produce  json
// @Param   id_periodo     path    string     true        "Id del periodo"
// @Param   sigla_curso     path    string     true        "Sigla del curso"
// @Param   sigla_grupo     path    string     true        "Sigla del grupo"
// @Success 200 {object} api_helpers.Json "OK"
// @Failure 400 {object} api_helpers.Error "Bad request"
// @Router /estudiantes/me/cursos/{id_periodo}/{sigla_curso}/grupos/{sigla_grupo}/estadisticas/evolucion-por-evaluacion [get]
func EvolucionEstudiantePorEvaluacion(c *gin.Context) {
	// params
	id_estudiante := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_ESTUDIANTE")
	// id_curso := c.Params.ByName("id_curso")
	sigla_grupo := c.Params.ByName("sigla_grupo")

	var calificaciones_estudiante []Query.CalificacionesEstudiantePorEvaluacion
	if err := repositories.CalificacionesEstudiantePorEvaluacion(&calificaciones_estudiante, sigla_grupo, id_estudiante); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondError(c, 500, "default")
			return
		}
	}

	response_container := &evolucionEstudiantePorEvaluacionResponse{
		Eje_x: utils.BuildUniqueEvaluacionesEvaluacion(calificaciones_estudiante),
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

	api_helpers.RespondJson(c, 200, response_container)
}
