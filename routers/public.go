package routers

import (
	"medroom-backend/controllers/auth"
	"medroom-backend/controllers/periodo"
	"medroom-backend/controllers/rol"

	"github.com/gin-gonic/gin"
)

func SetupPublicRouter(r *gin.Engine) *gin.Engine {

	router := r.Group("api/v1")
	{
		router.POST("estudiantes/login", auth.AuthenticateEstudiante)
		router.POST("evaluadores/login", auth.AuthenticateEvaluador)
		router.POST("administracion-academica/login", auth.AuthenticateAdministradorAcademico)
		router.POST("administracion-ti/login", auth.AuthenticateAdministradorTi)

		periodos := router.Group("/periodos")
		{
			periodos.GET("periodos", periodo.List)
			periodos.GET("periodos/:id", periodo.Get)
		}

		roles := router.Group("/roles")
		{
			roles.GET("roles", rol.List)
			roles.GET("roles/:id", rol.Get)
		}
	}

	return r
}
