package Routers

import (
	"medroom-backend/ApiHelpers"
	"medroom-backend/Controllers"
	"medroom-backend/Utils"

	"github.com/gin-gonic/gin"
)

func evaluadorAuthMiddleware(c *gin.Context) {
	authorization := c.GetHeader("authorization")

	if Utils.ValidarToken(authorization, "SECRET_KEY_EVALUADOR") {
		c.Next()
	} else {
		ApiHelpers.RespondError(c, 401, "default")
		c.Abort()
		return
	}
}

func SetupEvaluadorRouter(r *gin.Engine) *gin.Engine {

	router := r.Group("api/v1/evaluadores")
	router.Use(evaluadorAuthMiddleware)
	{
		// profile routes
		router.GET("me", Controllers.GetMyEvaluador)
		router.PUT("me", Controllers.PutMyEvaluador)

		// my course routes
		router.GET("me/cursos", Controllers.GetCursosEvaluador)
		router.GET("me/cursos/:id_curso", Controllers.GetOneCursoEvaluador)

		// my group routes
		router.GET("me/cursos/:id_curso/grupos", Controllers.GetGruposEvaluador)
		router.GET("me/cursos/:id_curso/grupos/:id_grupo", Controllers.GetOneGrupoEvaluador)

		// make evaluation routes
		router.GET("me/cursos/:id_curso/grupos/:id_grupo/evaluaciones", Controllers.ListEvaluacionesGrupoEvaluador)
		router.POST("me/cursos/:id_curso/grupos/:id_grupo/evaluaciones", Controllers.AddNewEvaluacion)
		router.POST("me/cursos/:id_curso/grupos/:id_grupo/estudiantes/:id_estudiante/evaluaciones/:id_evaluacion/calificacion", Controllers.AddNewCalificacionEstudiante)

		// reports
		// router.GET("me/reports", Controllers.GetReportesEvaluador)
	}

	return r
}
