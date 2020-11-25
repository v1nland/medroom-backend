package Routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"medroom-backend/Controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("authorization")

	r.Use(cors.New(config))

	// segment API by version
	v1 := r.Group("api/v1")
	{
		// segment by business domain
		estudiantes := v1.Group("/estudiantes")
		{
			estudiantes.GET("", Controllers.ListEstudiantes)
			estudiantes.GET(":id", Controllers.GetOneEstudiante)
			estudiantes.POST("", Controllers.AddNewEstudiante)
			estudiantes.PUT(":id", Controllers.PutOneEstudiante)
			estudiantes.DELETE(":id", Controllers.DeleteEstudiante)
		}

		// segment by business domain
		roles := v1.Group("/roles")
		{
			roles.GET("", Controllers.ListRols)
			roles.GET(":id", Controllers.GetOneRol)
			roles.POST("", Controllers.AddNewRol)
			roles.PUT(":id", Controllers.PutOneRol)
			roles.DELETE(":id", Controllers.DeleteRol)
		}

		// segment by business domain
		grupos := v1.Group("/grupos")
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
