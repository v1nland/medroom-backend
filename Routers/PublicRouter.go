package Routers

import (
	"medroom-backend/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupPublicRouter(r *gin.Engine) *gin.Engine {

	router := r.Group("api/v1")
	{
		// login routes
		router.POST("estudiantes/login", Controllers.AutenticarEstudiante)
		router.POST("evaluadores/login", Controllers.AutenticarEvaluador)
		router.POST("administracion-academica/login", Controllers.AutenticarAdministradorAcademico)
		router.POST("administracion-ti/login", Controllers.AutenticarAdministradorTi)

		// global routes
		router.GET("periodos", Controllers.ListPeriodos)
		router.GET("periodos/:id", Controllers.GetOnePeriodo)

		// puntajes := router.Group("/puntajes")
		// {
		// 	puntajes.GET("", Controllers.ListPuntajes)
		// 	puntajes.GET(":id", Controllers.GetOnePuntaje)
		// 	puntajes.POST("", Controllers.AddNewPuntaje)
		// 	puntajes.PUT(":id", Controllers.PutOnePuntaje)
		// 	puntajes.DELETE(":id", Controllers.DeletePuntaje)
		// }

		// evaluaciones := router.Group("/evaluaciones")
		// {
		// 	evaluaciones.GET("", Controllers.ListEvaluaciones)
		// 	evaluaciones.GET(":id", Controllers.GetOneEvaluacion)
		// 	evaluaciones.POST("", Controllers.AddNewEvaluacion)
		// 	evaluaciones.PUT(":id", Controllers.PutOneEvaluacion)
		// 	evaluaciones.DELETE(":id", Controllers.DeleteEvaluacion)
		// }
	}

	return r
}
