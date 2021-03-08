package routers

import (
	"medroom-backend/Utils"
	"medroom-backend/api_helpers"
	"medroom-backend/controllers/calificacion_estudiante"
	"medroom-backend/controllers/curso"
	"medroom-backend/controllers/evaluacion"
	"medroom-backend/controllers/evaluador"
	"medroom-backend/controllers/grupo"

	"github.com/gin-gonic/gin"
)

func evaluadorAuthMiddleware(c *gin.Context) {
	authorization := c.GetHeader("authorization")

	if Utils.ValidarToken(authorization, "SECRET_KEY_EVALUADOR") {
		c.Next()
	} else {
		api_helpers.RespondError(c, 401, "default")
		c.Abort()
		return
	}
}

func SetupEvaluadorRouter(r *gin.Engine) *gin.Engine {

	router := r.Group("api/v1/evaluadores")
	router.Use(evaluadorAuthMiddleware)
	{
		// profile routes
		router.GET("me", evaluador.GetMyEvaluador)
		router.PUT("me", evaluador.PutMyEvaluador)

		// my course routes
		router.GET("me/cursos", curso.GetCursosEvaluador)
		router.GET("me/cursos/:id_curso", curso.GetOneCursoEvaluador)

		// my group routes
		router.GET("me/cursos/:id_curso/grupos", grupo.GetGruposEvaluador)
		router.GET("me/cursos/:id_curso/grupos/:id_grupo", grupo.GetOneGrupoEvaluador)

		// make evaluation routes
		router.GET("me/cursos/:id_curso/grupos/:id_grupo/evaluaciones", evaluacion.ListEvaluacionesGrupoEvaluador)
		router.POST("me/cursos/:id_curso/grupos/:id_grupo/evaluaciones", evaluacion.AddNewEvaluacion)
		router.POST("me/cursos/:id_curso/grupos/:id_grupo/estudiantes/:id_estudiante/evaluaciones/:id_evaluacion/calificacion", calificacion_estudiante.AddNewCalificacionEstudiante)

		// reports
		// router.GET("me/reports", controllers.GetReportesEvaluador)
	}

	return r
}
