package routers

import (
	"medroom-backend/api_helpers"
	"medroom-backend/controllers/calificacion_estudiante"
	"medroom-backend/controllers/curso"
	"medroom-backend/controllers/estadistica"
	"medroom-backend/controllers/estudiante"
	"medroom-backend/controllers/evaluacion"
	"medroom-backend/controllers/grupo"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
)

func estudianteAuthMiddleware(c *gin.Context) {
	authorization := c.GetHeader("authorization")

	if utils.ValidarToken(authorization, "SECRET_KEY_ESTUDIANTE") {
		c.Next()
	} else {
		api_helpers.RespondError(c, 401, "default")
		c.Abort()
		return
	}
}

func SetupEstudianteRouter(r *gin.Engine) *gin.Engine {

	router := r.Group("api/v1/estudiantes")
	router.Use(estudianteAuthMiddleware)
	{
		// profile routes
		router.GET("me", estudiante.Profile)
		router.PUT("me", estudiante.PutProfile)

		// my course routes
		router.GET("me/cursos", curso.GetCursosEstudiante)
		router.GET("me/cursos/:id_periodo/:sigla_curso", curso.GetOneCursoEstudiante)

		// my group routes
		router.GET("me/cursos/:id_periodo/:sigla_curso/grupos", grupo.GetGruposEstudiante)
		router.GET("me/cursos/:id_periodo/:sigla_curso/grupos/:sigla_grupo", grupo.GetOneGrupoEstudiante)

		// my evaluations
		router.GET("me/cursos/:id_periodo/:sigla_curso/grupos/:sigla_grupo/evaluaciones", evaluacion.ListEvaluacionesGrupoEstudiante)
		router.GET("me/cursos/:id_periodo/:sigla_curso/grupos/:sigla_grupo/evaluaciones/:id_evaluacion/calificacion", calificacion_estudiante.Get)

		// reports
		router.GET("me/cursos/:id_periodo/:sigla_curso/grupos/:sigla_grupo/estadisticas/evolucion-por-competencia", estadistica.EvolucionEstudiantePorCompetencia)
		router.GET("me/cursos/:id_periodo/:sigla_curso/grupos/:sigla_grupo/estadisticas/evolucion-por-evaluacion", estadistica.EvolucionEstudiantePorEvaluacion)
	}

	return r
}
