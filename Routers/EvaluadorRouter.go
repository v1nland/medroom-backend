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

		// my group routes
		router.GET("me/grupo", Controllers.GetGrupoEvaluador)

		// my course routes
		router.GET("me/cursos", Controllers.GetCursosEvaluador)

		// make evaluation routes
		// call 'me/group' to list students and it's ramos
		router.POST("me/evaluaciones", Controllers.GenerarEvaluacion)

		// view reports routes
		// call 'me/group' to list students and it's ramos
		// router.GET("me/reports", Controllers.GetReportesEvaluador)
	}

	return r
}
