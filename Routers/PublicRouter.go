package Routers

import (
	"medroom-backend/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupPublicRouter(r *gin.Engine) *gin.Engine {

	router := r.Group("api/v1")
	{
		// login routes
		router.POST("estudiantes/login", Controllers.AutenticarEstudiante)
		router.POST("evaluadores/login", Controllers.AutenticarEvaluador)
		router.POST("administracion-academica/login", Controllers.AutenticarAdministradorAcademico)
		router.POST("administracion-ti/login", Controllers.AutenticarAdministradorTi)

		// traversal
		router.GET("periodos", Controllers.ListPeriodos)

		// testing
		// router.GET("gruposs", Controllers.ListGrupos)
		// router.GET("gruposs/:id", Controllers.GetOneGrupo)
		// router.GET("estudiantess", Controllers.ListEstudiantes)

		// estudiantes := router.Group("/estudiantes")
		// {
		// 	estudiantes.GET("", Controllers.ListEstudiantes)
		// 	estudiantes.GET(":id", Controllers.GetOneEstudiante)
		// 	estudiantes.POST("", Controllers.AddNewEstudiante)
		// 	estudiantes.PUT(":id", Controllers.PutOneEstudiante)
		// 	estudiantes.DELETE(":id", Controllers.DeleteEstudiante)
		// }

		// roles := router.Group("/roles")
		// {
		// 	roles.GET("", Controllers.ListRoles)
		// 	roles.GET(":id", Controllers.GetOneRol)
		// 	roles.POST("", Controllers.AddNewRol)
		// 	roles.PUT(":id", Controllers.PutOneRol)
		// 	roles.DELETE(":id", Controllers.DeleteRol)
		// }

		// grupos := router.Group("/grupos")
		// {
		// 	grupos.GET("", Controllers.ListGrupos)
		// 	grupos.GET(":id", Controllers.GetOneGrupo)
		// 	grupos.POST("", Controllers.AddNewGrupo)
		// 	grupos.PUT(":id", Controllers.PutOneGrupo)
		// 	grupos.DELETE(":id", Controllers.DeleteGrupo)
		// }

		// evaluadores := router.Group("/evaluadores")
		// {
		// 	evaluadores.GET("", Controllers.ListEvaluadores)
		// 	evaluadores.GET(":id", Controllers.GetOneEvaluador)
		// 	evaluadores.POST("", Controllers.AddNewEvaluador)
		// 	evaluadores.PUT(":id", Controllers.PutOneEvaluador)
		// 	evaluadores.DELETE(":id", Controllers.DeleteEvaluador)
		// }

		// cursos := router.Group("/cursos")
		// {
		// 	cursos.GET("", Controllers.ListCursos)
		// 	cursos.GET(":id", Controllers.GetOneCurso)
		// 	cursos.POST("", Controllers.AddNewCurso)
		// 	cursos.PUT(":id", Controllers.PutOneCurso)
		// 	cursos.DELETE(":id", Controllers.DeleteCurso)
		// }

		// periodos := router.Group("/periodos")
		// {
		// 	periodos.GET("", Controllers.ListPeriodos)
		// 	periodos.GET(":id", Controllers.GetOnePeriodo)
		// 	periodos.POST("", Controllers.AddNewPeriodo)
		// 	periodos.PUT(":id", Controllers.PutOnePeriodo)
		// 	periodos.DELETE(":id", Controllers.DeletePeriodo)
		// }

		// competencias := router.Group("/competencias")
		// {
		// 	competencias.GET("", Controllers.ListCompetencias)
		// 	competencias.GET(":id", Controllers.GetOneCompetencia)
		// 	competencias.POST("", Controllers.AddNewCompetencia)
		// 	competencias.PUT(":id", Controllers.PutOneCompetencia)
		// 	competencias.DELETE(":id", Controllers.DeleteCompetencia)
		// }

		// puntajes := router.Group("/puntajes")
		// {
		// 	puntajes.GET("", Controllers.ListPuntajes)
		// 	puntajes.GET(":id", Controllers.GetOnePuntaje)
		// 	puntajes.POST("", Controllers.AddNewPuntaje)
		// 	puntajes.PUT(":id", Controllers.PutOnePuntaje)
		// 	puntajes.DELETE(":id", Controllers.DeletePuntaje)
		// }

		// evaluaciones := router.Group("/evaluaciones")
		// {
		// 	evaluaciones.GET("", Controllers.ListEvaluaciones)
		// 	evaluaciones.GET(":id", Controllers.GetOneEvaluacion)
		// 	evaluaciones.POST("", Controllers.AddNewEvaluacion)
		// 	evaluaciones.PUT(":id", Controllers.PutOneEvaluacion)
		// 	evaluaciones.DELETE(":id", Controllers.DeleteEvaluacion)
		// }
	}

	return r
}
