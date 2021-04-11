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

type evolucionGrupoPorCompetenciaResponse struct {
	Eje_x   []string `json:"eje_x"`
	Eje_y   []int    `json:"eje_y"`
	Valores map[string][]struct {
		Promedio_grupo float64 `json:"promedio_grupo"`
	} `json:"valores"`
}

// @Summary Evolución por competencia
// @Description Obtiene la evolución de un grupo según competencia
// @Tags 03 - Evaluadores
// @Accept  json
// @Produce  json
// @Param   id_curso     path    string     true        "Id del curso"
// @Param   id_grupo     path    string     true        "Id del grupo"
// @Success 200 {array} Swagger.EvolucionGrupoPorCompetenciaSwagger "OK"
// @Failure 400 {object} api_helpers.ResponseError "Bad request"
// @Router /evaluadores/me/cursos/{id_curso}/grupos/{id_grupo}/estadisticas/evolucion-por-competencia [get]
func EvolucionGrupoPorCompetencia(c *gin.Context) {
	// params
	// id_evaluador := utils.DecodificarToken(c.GetHeader("authorization"), "SECRET_KEY_EVALUADOR")
	// id_curso := c.Params.ByName("id_curso")
	id_grupo := c.Params.ByName("id_grupo")

	var calificaciones_grupo []Query.CalificacionesGrupoPorCompetencia
	if err := repositories.CalificacionesGrupoPorCompetencia(&calificaciones_grupo, id_grupo); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			api_helpers.RespondError(c, 500, "default")
			return
		}
	}

	response_container := &evolucionGrupoPorCompetenciaResponse{
		Eje_x: utils.BuildUniqueEvaluacionesCompetenciaGrupo(calificaciones_grupo),
		Eje_y: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		Valores: map[string][]struct {
			Promedio_grupo float64 `json:"promedio_grupo"`
		}{},
	}

	for i := 0; i < len(calificaciones_grupo); i++ {
		response_container.Valores[calificaciones_grupo[i].Id_competencia] = append(response_container.Valores[calificaciones_grupo[i].Id_competencia], struct {
			Promedio_grupo float64 `json:"promedio_grupo"`
		}{
			calificaciones_grupo[i].Promedio_calificacion_puntaje_grupo,
		})
	}

	api_helpers.RespondJSON(c, 200, response_container)
}
