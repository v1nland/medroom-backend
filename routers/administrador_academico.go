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
		router.GET("me", administrador_academico.GetMyAdministradorAcademico)
		router.PUT("me", administrador_academico.PutMyAdministradorAcademico)

		estudiantes := router.Group("/estudiantes")
		{
			estudiantes.GET("", estudiante.ListEstudiantes)
			estudiantes.GET(":id", estudiante.GetOneEstudiante)
		}

		evaluadores := router.Group("/evaluadores")
		{
			evaluadores.GET("", evaluador.ListEvaluadores)
			evaluadores.GET(":id", evaluador.GetOneEvaluador)
		}

		cursos := router.Group("/me/cursos")
		{
			cursos.GET("", curso.GetCursosAdministradorAcademico)
			cursos.GET(":id", curso.GetOneCursoAdministradorAcademico)
			cursos.GET(":id/grupos", grupo.GetGruposAdministradorAcademico)
			cursos.GET(":id/grupos/:id_grupo", grupo.GetOneGrupoAdministradorAcademico)

			cursos.PUT(":id/grupos/:id_grupo/estudiantes/:id_estudiante", grupo.AddEstudianteToGrupo)
			cursos.PUT(":id/grupos/:id_grupo/evaluadores/:id_evaluador", grupo.AddEvaluadorToGrupo)

			cursos.GET(":id/estudiantes", estudiante.ListEstudiantesCurso)
			cursos.GET(":id/estudiantes/sin-grupo", estudiante.ListEstudiantesCursoSinGrupo)
		}

		grupos := router.Group("/me/grupos")
		{
			// grupos.GET("", controllers.GetGruposAdministradorAcademico)
			// grupos.GET(":id", controllers.GetOneGrupoAdministradorAcademico)
			grupos.POST("", grupo.AddNewGrupo)
			grupos.PUT(":id", grupo.PutOneGrupo)
			grupos.DELETE(":id", grupo.DeleteGrupo)
		}
	}

	return r
}
