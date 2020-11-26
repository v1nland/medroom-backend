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

		// segment by business domain
		evaluadores := v1.Group("/evaluadores")
		{
			evaluadores.GET("", Controllers.ListEvaluadors)
			evaluadores.GET(":id", Controllers.GetOneEvaluador)
			evaluadores.POST("", Controllers.AddNewEvaluador)
			evaluadores.PUT(":id", Controllers.PutOneEvaluador)
			evaluadores.DELETE(":id", Controllers.DeleteEvaluador)
		}

		// segment by business domain
		cursos := v1.Group("/cursos")
		{
			cursos.GET("", Controllers.ListCursos)
			cursos.GET(":id", Controllers.GetOneCurso)
			cursos.POST("", Controllers.AddNewCurso)
			cursos.PUT(":id", Controllers.PutOneCurso)
			cursos.DELETE(":id", Controllers.DeleteCurso)
		}

		// segment by business domain
		periodos := v1.Group("/periodos")
		{
			periodos.GET("", Controllers.ListPeriodos)
			periodos.GET(":id", Controllers.GetOnePeriodo)
			periodos.POST("", Controllers.AddNewPeriodo)
			periodos.PUT(":id", Controllers.PutOnePeriodo)
			periodos.DELETE(":id", Controllers.DeletePeriodo)
		}

		// segment by business domain
		competencias := v1.Group("/competencias")
		{
			competencias.GET("", Controllers.ListCompetencias)
			competencias.GET(":id", Controllers.GetOneCompetencia)
			competencias.POST("", Controllers.AddNewCompetencia)
			competencias.PUT(":id", Controllers.PutOneCompetencia)
			competencias.DELETE(":id", Controllers.DeleteCompetencia)
		}

		// segment by business domain
		puntajes := v1.Group("/puntajes")
		{
			puntajes.GET("", Controllers.ListPuntajes)
			puntajes.GET(":id", Controllers.GetOnePuntaje)
			puntajes.POST("", Controllers.AddNewPuntaje)
			puntajes.PUT(":id", Controllers.PutOnePuntaje)
			puntajes.DELETE(":id", Controllers.DeletePuntaje)
		}

		// segment by business domain
		evaluaciones := v1.Group("/evaluaciones")
		{
			evaluaciones.GET("", Controllers.ListEvaluacions)
			evaluaciones.GET(":id", Controllers.GetOneEvaluacion)
			evaluaciones.POST("", Controllers.AddNewEvaluacion)
			evaluaciones.PUT(":id", Controllers.PutOneEvaluacion)
			evaluaciones.DELETE(":id", Controllers.DeleteEvaluacion)
		}
	}

	return r
}
