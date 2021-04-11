package routers

import (
	"medroom-backend/api_helpers"
	"medroom-backend/controllers/administrador_academico"
	"medroom-backend/controllers/curso"
	"medroom-backend/controllers/estudiante"
	"medroom-backend/controllers/evaluador"
	"medroom-backend/controllers/grupo"
	"medroom-backend/utils"

	"github.com/gin-gonic/gin"
)

func administradorAcademicoAuthMiddleware(c *gin.Context) {
	authorization := c.GetHeader("authorization")

	if utils.ValidarToken(authorization, "SECRET_KEY_ADMINISTRADOR_ACADEMICO") {
		c.Next()
	} else {
		api_helpers.RespondError(c, 401, "default")
		c.Abort()
		return
	}
}

func SetupAdministradorAcademicoRouter(r *gin.Engine) *gin.Engine {

	router := r.Group("api/v1/administracion-academica")
	router.Use(administradorAcademicoAuthMiddleware)
	{
		router.GET("me", administrador_academico.Profile)
		router.PUT("me", administrador_academico.PutProfile)

		estudiantes := router.Group("/estudiantes")
		{
			estudiantes.GET("", estudiante.List)
			estudiantes.GET(":id", estudiante.Get)
		}

		evaluadores := router.Group("/evaluadores")
		{
			evaluadores.GET("", evaluador.List)
			evaluadores.GET(":id", evaluador.Get)
		}

		cursos := router.Group("/me/cursos")
		{
			cursos.GET("", curso.GetCursosAdministradorAcademico)
			cursos.GET(":id", curso.GetOneCursoAdministradorAcademico)
			cursos.GET(":id/grupos", grupo.GetGruposAdministradorAcademico)
			cursos.GET(":id/grupos/:id_grupo", grupo.GetOneGrupoAdministradorAcademico)

			cursos.POST(":id/grupos", grupo.Add)
			cursos.PUT(":id/grupos/:id_grupo", grupo.Put)
			cursos.DELETE(":id/grupos/:id_grupo", grupo.Delete)

			cursos.PUT(":id/grupos/:id_grupo/estudiantes/:id_estudiante", grupo.AddEstudianteToGrupo)
			cursos.PUT(":id/grupos/:id_grupo/evaluadores/:id_evaluador", grupo.AddEvaluadorToGrupo)

			cursos.DELETE(":id/grupos/:id_grupo/estudiantes/:id_estudiante", grupo.RemoveEstudianteFromGrupo)
			cursos.DELETE(":id/grupos/:id_grupo/evaluadores/:id_evaluador", grupo.RemoveEvaluadorFromGrupo)

			cursos.GET(":id/estudiantes", estudiante.ListEstudiantesCurso)
			cursos.GET(":id/estudiantes/sin-grupo", estudiante.ListEstudiantesCursoSinGrupo)
		}
	}

	return r
}
