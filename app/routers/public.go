package routers

import (
	"medroom-backend/app/controllers/auth"
	"medroom-backend/app/controllers/periodo"
	"medroom-backend/app/controllers/rol"

	"github.com/gin-gonic/gin"
)

func SetupPublicRouter(r *gin.Engine) *gin.Engine {

	router := r.Group("api/v1")
	{
		router.POST("estudiantes/login", auth.AutenticarEstudiante)
		router.POST("evaluadores/login", auth.AutenticarEvaluador)
		router.POST("administracion-academica/login", auth.AutenticarAdministradorAcademico)
		router.POST("administracion-ti/login", auth.AutenticarAdministradorTi)

		periodos := router.Group("/periodos")
		{
			periodos.GET("periodos", periodo.ListPeriodos)
			periodos.GET("periodos/:id", periodo.GetOnePeriodo)
		}

		roles := router.Group("/roles")
		{
			roles.GET("roles", rol.ListRoles)
			roles.GET("roles/:id", rol.GetOneRol)
		}
	}

	return r
}
