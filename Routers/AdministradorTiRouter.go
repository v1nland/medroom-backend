package Routers

import (
	"medroom-backend/ApiHelpers"
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
		// router.GET("me", Controllers.GetMyAdministradorTi)
		// router.PUT("me", Controllers.PutMyAdministradorTi)

		// create users routes
		// router.POST("estudiantes", AddNewEstudiante)
		// router.POST("evaluadores", AddNewEvaluador)
		// router.POST("administradores-academicos", AddNewAdministradorAcademico)

		// courses routes
		// router.POST("cursos", AddNewCurso)

		// manage users routes
		// router.POST("estudiantes", AddNewEstudiante)
		// router.PUT("cursos", PutOneCurso)
		// router.DELETE("grupos", DeleteGrupo)
		// router.PUT("estudiantes", Controllers.AsociarEstudiantesConGrupo) // ???? not sure
	}

	return r
}
