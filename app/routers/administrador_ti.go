package routers

import (
	"medroom-backend/app/api_helpers"
	"medroom-backend/app/controllers/administrador_academico"
	"medroom-backend/app/controllers/administrador_ti"
	"medroom-backend/app/controllers/curso"
	"medroom-backend/app/controllers/estudiante"
	"medroom-backend/app/controllers/evaluador"
	"medroom-backend/app/controllers/grupo"
	"medroom-backend/app/controllers/periodo"
	"medroom-backend/app/controllers/rol"
	"medroom-backend/app/utils"

	"github.com/gin-gonic/gin"
)

func administradorTiAuthMiddleware(c *gin.Context) {
	authorization := c.GetHeader("authorization")

	if utils.ValidarToken(authorization, "SECRET_KEY_ADMINISTRADOR_TI") {
		c.Next()
	} else {
		api_helpers.RespondError(c, 401, "default")
		c.Abort()
		return
	}
}

func SetupAdministradorTiRouter(r *gin.Engine) *gin.Engine {

	router := r.Group("api/v1/administracion-ti")
	router.Use(administradorTiAuthMiddleware)
	{
		router.GET("me", administrador_ti.GetMyAdministradorTi)
		router.PUT("me", administrador_ti.PutMyAdministradorTi)

		estudiantes := router.Group("/estudiantes")
		{
			estudiantes.GET("", estudiante.ListEstudiantes)
			estudiantes.GET(":id", estudiante.GetOneEstudiante)
			estudiantes.POST("", estudiante.AddNewEstudiante)
			estudiantes.PUT(":id", estudiante.PutOneEstudiante)
			estudiantes.DELETE(":id", estudiante.DeleteEstudiante)
		}

		evaluadores := router.Group("/evaluadores")
		{
			evaluadores.GET("", evaluador.ListEvaluadores)
			evaluadores.GET(":id", evaluador.GetOneEvaluador)
			evaluadores.POST("", evaluador.AddNewEvaluador)
			evaluadores.PUT(":id", evaluador.PutOneEvaluador)
			evaluadores.DELETE(":id", evaluador.DeleteEvaluador)
		}

		administradores_academicos := router.Group("/administradores-academicos")
		{
			administradores_academicos.GET("", administrador_academico.ListAdministradoresAcademicos)
			administradores_academicos.GET(":id", administrador_academico.GetOneAdministradorAcademico)
			administradores_academicos.POST("", administrador_academico.AddNewAdministradorAcademico)
			administradores_academicos.PUT(":id", administrador_academico.PutOneAdministradorAcademico)
			administradores_academicos.DELETE(":id", administrador_academico.DeleteAdministradorAcademico)
		}

		administradores_ti := router.Group("/administradores-ti")
		{
			administradores_ti.GET("", administrador_ti.ListAdministradoresTi)
			administradores_ti.GET(":id", administrador_ti.GetOneAdministradorTi)
			administradores_ti.POST("", administrador_ti.AddNewAdministradorTi)
			administradores_ti.PUT(":id", administrador_ti.PutOneAdministradorTi)
			administradores_ti.DELETE(":id", administrador_ti.DeleteAdministradorTi)
		}

		cursos := router.Group("/cursos")
		{
			cursos.GET("", curso.ListCursos)
			cursos.GET(":id", curso.GetOneCurso)
			cursos.POST("", curso.AddNewCurso)
			cursos.POST("carga-masiva", curso.AddNewCursos)
			cursos.PUT(":id", curso.PutOneCurso)
			cursos.PUT(":id/estudiantes/:id_estudiante", curso.AddEstudianteToCurso)
			cursos.PUT(":id/evaluadores/:id_evaluador", curso.AddEvaluadorToCurso)
			cursos.PUT(":id/administradores-academicos/:id_administrador_academico", curso.AddAdministradorAcademicoToCurso)
			cursos.DELETE(":id", curso.DeleteCurso)
		}

		grupos := router.Group("/grupos")
		{
			grupos.GET("", grupo.ListGrupos)
			grupos.GET(":id", grupo.GetOneGrupo)
			grupos.POST("", grupo.AddNewGrupo)
			grupos.PUT(":id", grupo.PutOneGrupo)
			grupos.DELETE(":id", grupo.DeleteGrupo)
		}

		periodos := router.Group("/periodos")
		{
			periodos.GET("", periodo.ListPeriodos)
			periodos.GET(":id", periodo.GetOnePeriodo)
			periodos.POST("", periodo.AddNewPeriodo)
			periodos.PUT(":id", periodo.PutOnePeriodo)
			periodos.DELETE(":id", periodo.DeletePeriodo)
		}

		roles := router.Group("/roles")
		{
			roles.GET("", rol.ListRoles)
			roles.GET(":id", rol.GetOneRol)
			roles.POST("", rol.AddNewRol)
			roles.PUT(":id", rol.PutOneRol)
			roles.DELETE(":id", rol.DeleteRol)
		}
	}

	return r
}
