package Routers

import (
	"medroom-backend/ApiHelpers"
	"medroom-backend/Controllers"
	"medroom-backend/Utils"

	"github.com/gin-gonic/gin"
)

func administradorTiAuthMiddleware(c *gin.Context) {
	authorization := c.GetHeader("authorization")

	if Utils.ValidarToken(authorization, "SECRET_KEY_ADMINISTRADOR_TI") {
		c.Next()
	} else {
		ApiHelpers.RespondError(c, 401, "default")
		c.Abort()
		return
	}
}

func SetupAdministradorTiRouter(r *gin.Engine) *gin.Engine {

	router := r.Group("api/v1/administracion-ti")
	router.Use(administradorTiAuthMiddleware)
	{
		// profile routes
		router.GET("me", Controllers.GetMyAdministradorTi)
		router.PUT("me", Controllers.PutMyAdministradorTi)

		estudiantes := router.Group("/estudiantes")
		{
			estudiantes.GET("", Controllers.ListEstudiantes)
			estudiantes.GET(":id", Controllers.GetOneEstudiante)
			estudiantes.POST("", Controllers.AddNewEstudiante)
			estudiantes.PUT(":id", Controllers.PutOneEstudiante)
			// estudiantes.PUT(":id/grupos", Controllers.AddEstudianteToGrupo)
			estudiantes.DELETE(":id", Controllers.DeleteEstudiante)
		}

		evaluadores := router.Group("/evaluadores")
		{
			evaluadores.GET("", Controllers.ListEvaluadores)
			evaluadores.GET(":id", Controllers.GetOneEvaluador)
			evaluadores.POST("", Controllers.AddNewEvaluador)
			evaluadores.PUT(":id", Controllers.PutOneEvaluador)
			// evaluadores.PUT(":id/grupos", Controllers.AddEvaluadorToGrupo)
			evaluadores.DELETE(":id", Controllers.DeleteEvaluador)
		}

		administradores_academicos := router.Group("/administradores-academicos")
		{
			administradores_academicos.GET("", Controllers.ListAdministradoresAcademicos)
			administradores_academicos.GET(":id", Controllers.GetOneAdministradorAcademico)
			administradores_academicos.POST("", Controllers.AddNewAdministradorAcademico)
			administradores_academicos.PUT(":id", Controllers.PutOneAdministradorAcademico)
			administradores_academicos.DELETE(":id", Controllers.DeleteAdministradorAcademico)
		}

		administradores_ti := router.Group("/administradores-ti")
		{
			administradores_ti.GET("", Controllers.ListAdministradoresTi)
			administradores_ti.GET(":id", Controllers.GetOneAdministradorTi)
			administradores_ti.POST("", Controllers.AddNewAdministradorTi)
			administradores_ti.PUT(":id", Controllers.PutOneAdministradorTi)
			administradores_ti.DELETE(":id", Controllers.DeleteAdministradorTi)
		}

		cursos := router.Group("/cursos")
		{
			cursos.GET("", Controllers.ListCursos)
			cursos.GET(":id", Controllers.GetOneCurso)
			cursos.POST("", Controllers.AddNewCurso)
			cursos.PUT(":id", Controllers.PutOneCurso)
			cursos.PUT(":id/estudiantes/:id_estudiante", Controllers.AddEstudianteToCurso)
			cursos.DELETE(":id", Controllers.DeleteCurso)
		}

		grupos := router.Group("/grupos")
		{
			grupos.GET("", Controllers.ListGrupos)
			grupos.GET(":id", Controllers.GetOneGrupo)
			grupos.POST("", Controllers.AddNewGrupo)
			grupos.PUT(":id", Controllers.PutOneGrupo)
			cursos.PUT(":id/grupos/:id_grupo/estudiantes/:id_estudiante", Controllers.AddEstudianteToGrupo)
			grupos.DELETE(":id", Controllers.DeleteGrupo)
		}

		periodos := router.Group("/periodos")
		{
			periodos.GET("", Controllers.ListPeriodos)
			periodos.GET(":id", Controllers.GetOnePeriodo)
			periodos.POST("", Controllers.AddNewPeriodo)
			periodos.PUT(":id", Controllers.PutOnePeriodo)
			periodos.DELETE(":id", Controllers.DeletePeriodo)
		}

		roles := router.Group("/roles")
		{
			roles.GET("", Controllers.ListRoles)
			roles.GET(":id", Controllers.GetOneRol)
			roles.POST("", Controllers.AddNewRol)
			roles.PUT(":id", Controllers.PutOneRol)
			roles.DELETE(":id", Controllers.DeleteRol)
		}
	}

	return r
}
