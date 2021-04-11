package routers

import (
	"medroom-backend/api_helpers"
	"medroom-backend/controllers/administrador_academico"
	"medroom-backend/controllers/administrador_ti"
	"medroom-backend/controllers/curso"
	"medroom-backend/controllers/estudiante"
	"medroom-backend/controllers/evaluador"
	"medroom-backend/controllers/grupo"
	"medroom-backend/controllers/periodo"
	"medroom-backend/controllers/rol"
	"medroom-backend/utils"

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
		router.GET("me", administrador_ti.Profile)
		router.PUT("me", administrador_ti.PutProfile)

		estudiantes := router.Group("/estudiantes")
		{
			estudiantes.GET("", estudiante.List)
			estudiantes.GET(":id", estudiante.Get)
			estudiantes.POST("", estudiante.Add)
			estudiantes.POST("carga-masiva", estudiante.AddNewEstudiantes)
			estudiantes.PUT(":id", estudiante.Put)
			estudiantes.DELETE(":id", estudiante.Delete)
		}

		evaluadores := router.Group("/evaluadores")
		{
			evaluadores.GET("", evaluador.List)
			evaluadores.GET(":id", evaluador.Get)
			evaluadores.POST("", evaluador.Add)
			evaluadores.PUT(":id", evaluador.Put)
			evaluadores.DELETE(":id", evaluador.Delete)
		}

		administradores_academicos := router.Group("/administradores-academicos")
		{
			administradores_academicos.GET("", administrador_academico.List)
			administradores_academicos.GET(":id", administrador_academico.Get)
			administradores_academicos.POST("", administrador_academico.Add)
			administradores_academicos.POST("/carga-masiva", administrador_academico.MassiveAdd)
			administradores_academicos.PUT(":id", administrador_academico.Put)
			administradores_academicos.DELETE(":id", administrador_academico.Delete)
		}

		administradores_ti := router.Group("/administradores-ti")
		{
			administradores_ti.GET("", administrador_ti.List)
			administradores_ti.GET(":id", administrador_ti.Get)
			administradores_ti.POST("", administrador_ti.Add)
			administradores_ti.PUT(":id", administrador_ti.Put)
			administradores_ti.DELETE(":id", administrador_ti.Delete)
		}

		cursos := router.Group("/cursos")
		{
			cursos.GET("", curso.ListCursos)
			cursos.GET(":id", curso.GetOneCurso)
			cursos.GET(":id/reporte-notas", curso.GetMarksReport)
			cursos.POST("", curso.Add)
			cursos.POST("carga-masiva", curso.AddNewCursos)
			cursos.PUT(":id", curso.Put)
			cursos.PUT(":id/estudiantes/:id_estudiante", curso.AddEstudianteToCurso)
			cursos.PUT(":id/evaluadores/:id_evaluador", curso.AddEvaluadorToCurso)
			cursos.PUT(":id/administradores-academicos/:id_administrador_academico", curso.AddAdministradorAcademicoToCurso)
			cursos.DELETE(":id", curso.Delete)
		}

		grupos := router.Group("/grupos")
		{
			grupos.GET("", grupo.ListGrupos)
			grupos.GET(":id", grupo.GetOneGrupo)
			grupos.POST("", grupo.Add)
			grupos.POST("carga-masiva", grupo.AddNewGrupos)
			grupos.PUT(":id", grupo.Put)
			grupos.DELETE(":id", grupo.Delete)
		}

		periodos := router.Group("/periodos")
		{
			periodos.GET("", periodo.List)
			periodos.GET(":id", periodo.Get)
			periodos.POST("", periodo.Add)
			periodos.PUT(":id", periodo.Put)
			periodos.DELETE(":id", periodo.Delete)
		}

		roles := router.Group("/roles")
		{
			roles.GET("", rol.List)
			roles.GET(":id", rol.Get)
			roles.POST("", rol.Add)
			roles.PUT(":id", rol.PutOneRol)
			roles.DELETE(":id", rol.Delete)
		}
	}

	return r
}
