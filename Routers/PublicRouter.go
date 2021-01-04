package Routers

import (
	"medroom-backend/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupPublicRouter(r *gin.Engine) *gin.Engine {

	router := r.Group("api/v1")
	{
		router.POST("estudiantes/login", Controllers.AutenticarEstudiante)
		router.POST("evaluadores/login", Controllers.AutenticarEvaluador)
		router.POST("administracion-academica/login", Controllers.AutenticarAdministradorAcademico)
		router.POST("administracion-ti/login", Controllers.AutenticarAdministradorTi)

		periodos := router.Group("/periodos")
		{
			periodos.GET("periodos", Controllers.ListPeriodos)
			periodos.GET("periodos/:id", Controllers.GetOnePeriodo)
		}

		roles := router.Group("/roles")
		{
			roles.GET("roles", Controllers.ListRoles)
			roles.GET("roles/:id", Controllers.GetOneRol)
		}
	}

	return r
}
