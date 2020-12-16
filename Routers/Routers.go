package Routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("authorization")

	r.Use(cors.New(config))

	r = SetupPublicRouter(r)
	r = SetupEstudianteRouter(r)
	r = SetupAdministradorAcademicoRouter(r)
	r = SetupAdministradorTiRouter(r)

	return r
}
