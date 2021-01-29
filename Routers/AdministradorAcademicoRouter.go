package Routers

import (
	"medroom-backend/ApiHelpers"
	"medroom-backend/Controllers"
	"medroom-backend/Utils"

	"github.com/gin-gonic/gin"
)

func administradorAcademicoAuthMiddleware(c *gin.Context) {
	authorization := c.GetHeader("authorization")

	if Utils.ValidarToken(authorization, "SECRET_KEY_ADMINISTRADOR_ACADEMICO") {
		c.Next()
	} else {
		ApiHelpers.RespondError(c, 401, "default")
		c.Abort()
		return
	}
}

func SetupAdministradorAcademicoRouter(r *gin.Engine) *gin.Engine {

	router := r.Group("api/v1/administracion-academica")
	router.Use(administradorAcademicoAuthMiddleware)
	{
		// profile routes
		router.GET("me", Controllers.GetMyAdministradorAcademico)
		router.PUT("me", Controllers.PutMyAdministradorAcademico)

		estudiantes := router.Group("/estudiantes")
		{
			estudiantes.GET("", Controllers.ListEstudiantes)
			estudiantes.GET(":id", Controllers.GetOneEstudiante)
		}

		evaluadores := router.Group("/evaluadores")
		{
			evaluadores.GET("", Controllers.ListEvaluadores)
			evaluadores.GET(":id", Controllers.GetOneEvaluador)
		}

		cursos := router.Group("/cursos")
		{
			cursos.GET("", Controllers.ListCursos)
			cursos.GET(":id", Controllers.GetOneCurso)
			cursos.PUT(":id/grupos/:id_grupo/estudiantes/:id_estudiante", Controllers.AddEstudianteToGrupo)
			cursos.PUT(":id/grupos/:id_grupo/evaluadores/:id_evaluador", Controllers.AddEvaluadorToGrupo)

			cursos.GET(":id/estudiantes", Controllers.ListEstudiantesCurso)
			cursos.GET(":id/estudiantes/sin-grupo", Controllers.ListEstudiantesCursoSinGrupo)
		}

		grupos := router.Group("/grupos")
		{
			grupos.GET("", Controllers.ListGrupos)
			grupos.GET(":id", Controllers.GetOneGrupo)
			grupos.POST("", Controllers.AddNewGrupo)
			grupos.PUT(":id", Controllers.PutOneGrupo)
			grupos.DELETE(":id", Controllers.DeleteGrupo)
		}
	}

	return r
}
