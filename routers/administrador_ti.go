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
			estudiantes.GET(":id_estudiante", estudiante.Get)
			estudiantes.POST("", estudiante.Add)
			estudiantes.PUT(":id_estudiante", estudiante.Put)
			estudiantes.PUT(":id_estudiante/reestablecer", administrador_ti.ResetPasswordEstudiante)
			estudiantes.DELETE(":id_estudiante", estudiante.Delete)
		}

		evaluadores := router.Group("/evaluadores")
		{
			evaluadores.GET("", evaluador.List)
			evaluadores.GET(":id_evaluador", evaluador.Get)
			evaluadores.POST("", evaluador.Add)
			evaluadores.PUT(":id_evaluador", evaluador.Put)
			evaluadores.DELETE(":id_evaluador", evaluador.Delete)
		}

		administradores_academicos := router.Group("/administradores-academicos")
		{
			administradores_academicos.GET("", administrador_academico.List)
			administradores_academicos.GET(":id_administrador_academico", administrador_academico.Get)
			administradores_academicos.POST("", administrador_academico.Add)
			administradores_academicos.PUT(":id_administrador_academico", administrador_academico.Put)
			administradores_academicos.DELETE(":id_administrador_academico", administrador_academico.Delete)
		}

		administradores_ti := router.Group("/administradores-ti")
		{
			administradores_ti.GET("", administrador_ti.List)
			administradores_ti.GET(":id_administrador_ti", administrador_ti.Get)
			administradores_ti.POST("", administrador_ti.Add)
			administradores_ti.PUT(":id_administrador_ti", administrador_ti.Put)
			administradores_ti.DELETE(":id_administrador_ti", administrador_ti.Delete)
		}

		cursos := router.Group("/cursos")
		{
			cursos.GET("", curso.ListCursos)
			cursos.GET(":id_periodo/:sigla_curso", curso.GetOneCurso)
			cursos.GET(":id_periodo/:sigla_curso/reporte-notas", curso.GetMarksReport)
			cursos.POST(":id_periodo", curso.Add)
			cursos.PUT(":id_periodo/:sigla_curso", curso.Put)
			cursos.PUT(":id_periodo/:sigla_curso/estudiantes/:id_estudiante", curso.AddEstudianteToCurso)
			cursos.PUT(":id_periodo/:sigla_curso/evaluadores/:id_evaluador", curso.AddEvaluadorToCurso)
			cursos.PUT(":id_periodo/:sigla_curso/administradores-academicos/:id_administrador_academico", curso.AddAdministradorAcademicoToCurso)
			cursos.DELETE(":id_periodo/:sigla_curso", curso.Delete)

			cursos.GET(":id_periodo/:sigla_curso/grupos", grupo.ListGrupos)
			cursos.GET(":id_periodo/:sigla_curso/grupos/:sigla_grupo", grupo.GetOneGrupo)
			cursos.POST(":id_periodo/:sigla_curso/grupos", grupo.Add)
			cursos.PUT(":id_periodo/:sigla_curso/grupos/:sigla_grupo", grupo.Put)
			cursos.DELETE(":id_periodo/:sigla_curso/grupos/:sigla_grupo", grupo.Delete)
		}

		carga_masiva := router.Group("/carga-masiva")
		{
			carga_masiva.POST("cursos", curso.MassiveAdd)
			carga_masiva.POST("grupos", grupo.MassiveAdd)
			carga_masiva.POST("estudiantes", estudiante.MassiveAdd)
			carga_masiva.POST("administradores-academicos", administrador_academico.MassiveAdd)
		}

		periodos := router.Group("/periodos")
		{
			periodos.GET("", periodo.List)
			periodos.GET(":id_periodo", periodo.Get)
			periodos.POST("", periodo.Add)
			periodos.PUT(":id_periodo", periodo.Put)
			periodos.DELETE(":id_periodo", periodo.Delete)
		}

		roles := router.Group("/roles")
		{
			roles.GET("", rol.List)
			roles.GET(":id_rol", rol.Get)
			roles.POST("", rol.Add)
			roles.PUT(":id_rol", rol.PutOneRol)
			roles.DELETE(":id_rol", rol.Delete)
		}
	}

	return r
}
