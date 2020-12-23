package Routers

import (
	"medroom-backend/ApiHelpers"
	"medroom-backend/Controllers"
	"medroom-backend/Utils"

	"github.com/gin-gonic/gin"
)

func estudianteAuthMiddleware(c *gin.Context) {
	authorization := c.GetHeader("authorization")

	if Utils.ValidarToken(authorization, "SECRET_KEY_ESTUDIANTE") {
		c.Next()
	} else {
		ApiHelpers.RespondError(c, 401, "default")
		c.Abort()
		return
	}
}

func SetupEstudianteRouter(r *gin.Engine) *gin.Engine {

	router := r.Group("api/v1/estudiantes")
	router.Use(estudianteAuthMiddleware)
	{
		// profile routes
		router.GET("me", Controllers.GetMyEstudiante)
		router.PUT("me", Controllers.PutMyEstudiante)

		// my group routes
		router.GET("me/grupo", Controllers.GetGrupoEstudiante)

		// my course routes
		router.GET("me/cursos", Controllers.GetCursosEstudiante)

		// my evaluations
		router.GET("me/evaluaciones", Controllers.ListEvaluacionesEstudiante)
		// router.GET("me/estadisticas/evolucion/evaluacion/:id_evaluacion", Controllers.EvolucionEstudiantePorEvaluacion)
		// router.GET("me/estadisticas/evolucion/competencia/:codigo_competencia", Controllers.EvolucionEstudiantePorCompetencia)
	}

	return r
}
