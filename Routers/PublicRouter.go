package Routers

import (
	"medroom-backend/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupPublicRouter(r *gin.Engine) *gin.Engine {

	router := r.Group("api/v1")
	{
		router.POST("login", Controllers.AutenticarEstudiante)
		router.POST("login-evaluador", Controllers.AutenticarEvaluador)
		router.POST("login-administrador-academico", Controllers.AutenticarAdministradorAcademico)
		router.POST("login-administrador-ti", Controllers.AutenticarAdministradorTi)

		router.GET("gruposs", Controllers.ListGrupos)
		router.GET("gruposs/:id", Controllers.GetOneGrupo)
		router.GET("estudiantess", Controllers.ListEstudiantes)
	}

	return r
}
