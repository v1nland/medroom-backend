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

		// my course routes
		router.GET("me/cursos", Controllers.GetCursosEstudiante)
		router.GET("me/cursos/:id_curso", Controllers.GetOneCursoEstudiante)

		// my group routes
		router.GET("me/cursos/:id_curso/grupos", Controllers.GetGruposEstudiante)
		router.GET("me/cursos/:id_curso/grupos/:id_grupo", Controllers.GetOneGrupoEstudiante)

		// my evaluations
		router.GET("me/cursos/:id_curso/grupos/:id_grupo/evaluaciones", Controllers.ListEvaluacionesGrupoEstudiante)
		router.GET("me/cursos/:id_curso/grupos/:id_grupo/evaluaciones/:id_evaluacion/calificacion", Controllers.GetOneCalificacionEstudiante)

		// reports
		router.GET("me/cursos/:id_curso/grupos/:id_grupo/estadisticas/evolucion-por-competencia", Controllers.EvolucionEstudiantePorCompetencia)
		// router.GET("me/cursos/:id_curso/grupos/:id_grupo/estadisticas/evolucion/competencia/:id_competencia", Controllers.EvolucionEstudiantePorCompetencia)
	}

	return r
}
