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

		grupos := router.Group("/grupos")
		{
			grupos.GET("", Controllers.ListGrupos)
			grupos.GET(":id", Controllers.GetOneGrupo)
			grupos.POST("", Controllers.AddNewGrupo)
			grupos.PUT(":id", Controllers.PutOneGrupo)
			grupos.PUT(":id/grupos/:id_grupo/estudiantes/:id_estudiante", Controllers.AddEstudianteToGrupo)
			grupos.PUT(":id/grupos/:id_grupo/evaluadores/:id_evaluador", Controllers.AddEvaluadorToGrupo)
			grupos.DELETE(":id", Controllers.DeleteGrupo)
		}

		// cursos routes
		// router.POST("grupos", Controllers.AddNewGrupo)
		// router.PUT("grupos", Controllers.PutOneGrupo)
		// router.DELETE("grupos", Controllers.DeleteGrupo)
		// router.PUT("estudiantes", Controllers.AsociarEstudiantesConGrupo)

	}

	return r
}
