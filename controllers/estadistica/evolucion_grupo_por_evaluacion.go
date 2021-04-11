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

type evolucionGrupoPorEvaluacionResponse struct {
	Eje_x   []string `json:"eje_x"`
	Eje_y   []int    `json:"eje_y"`
	Valores map[string][]struct {
		Promedio_grupo float64 `json:"promedio_grupo"`
	} `json:"valores"`
}

// @Summary Evolución por evaluacion
// @Description Obtiene la evolución de un grupo según evaluacion
// @Tags 03 - Evaluadores
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso"
// @Param   id_grupo     path    string     true        "Id del grupo"
// @Success 200 {array} Swagger.EvolucionGrupoPorEvaluacionSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /evaluadores/me/cursos/{id_curso}/grupos/{id_grupo}/estadisticas/evolucion-por-evaluacion [get]
func EvolucionGrupoPorEvaluacion(c *gin.Context) {
	// params
	// id_evaluador := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")
	// id_curso := c.Params.ByName("id_curso")
	id_grupo := c.Params.ByName("id_grupo")

	var calificaciones_grupo []Query.CalificacionesGrupoPorEvaluacion
	if err := repositories.CalificacionesGrupoPorEvaluacion(&calificaciones_grupo, id_grupo); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondError(c, 500, "default")
			return
		}
	}

	response_container := &evolucionGrupoPorEvaluacionResponse{
		Eje_x: utils.BuildUniqueEvaluacionesEvaluacionGrupo(calificaciones_grupo),
		Eje_y: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		Valores: map[string][]struct {
			Promedio_grupo float64 `json:"promedio_grupo"`
		}{},
	}

	for i := 0; i < len(calificaciones_grupo); i++ {
		response_container.Valores[calificaciones_grupo[i].Nombre_evaluacion] = append(response_container.Valores[calificaciones_grupo[i].Nombre_evaluacion], struct {
			Promedio_grupo float64 `json:"promedio_grupo"`
		}{
			calificaciones_grupo[i].Promedio_calificacion_puntaje_grupo,
		})
	}

	api_helpers.RespondJSON(c, 200, response_container)
}
