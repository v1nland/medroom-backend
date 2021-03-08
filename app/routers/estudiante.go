package routers

import (
	"medroom-backend/app/api_helpers"
	"medroom-backend/app/controllers/calificacion_estudiante"
	"medroom-backend/app/controllers/curso"
	"medroom-backend/app/controllers/estadistica"
	"medroom-backend/app/controllers/estudiante"
	"medroom-backend/app/controllers/evaluacion"
	"medroom-backend/app/controllers/grupo"
	"medroom-backend/app/utils"

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
		router.GET("me", estudiante.GetMyEstudiante)
		router.PUT("me", estudiante.PutMyEstudiante)

		// my course routes
		router.GET("me/cursos", curso.GetCursosEstudiante)
		router.GET("me/cursos/:id_curso", curso.GetOneCursoEstudiante)

		// my group routes
		router.GET("me/cursos/:id_curso/grupos", grupo.GetGruposEstudiante)
		router.GET("me/cursos/:id_curso/grupos/:id_grupo", grupo.GetOneGrupoEstudiante)

		// my evaluations
		router.GET("me/cursos/:id_curso/grupos/:id_grupo/evaluaciones", evaluacion.ListEvaluacionesGrupoEstudiante)
		router.GET("me/cursos/:id_curso/grupos/:id_grupo/evaluaciones/:id_evaluacion/calificacion", calificacion_estudiante.GetOneCalificacionEstudiante)

		// reports
		router.GET("me/cursos/:id_curso/grupos/:id_grupo/estadisticas/evolucion-por-competencia", estadistica.EvolucionEstudiantePorCompetencia)
		router.GET("me/cursos/:id_curso/grupos/:id_grupo/estadisticas/evolucion-por-evaluacion", estadistica.EvolucionEstudiantePorEvaluacion)
	}

	return r
}
