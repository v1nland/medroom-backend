package Routers

import (
	"medroom-backend/ApiHelpers"
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
		// router.GET("me", Controllers.GetMyAdministradorAcademico)
		// router.PUT("me", Controllers.PutMyAdministradorAcademico)

		// cursos routes
		// router.POST("grupos", Controllers.AddNewGrupo)
		// router.PUT("grupos", Controllers.PutOneGrupo)
		// router.DELETE("grupos", Controllers.DeleteGrupo)
		// router.PUT("estudiantes", Controllers.AsociarEstudiantesConGrupo)

	}

	return r
}
