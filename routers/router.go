package routers

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	config := cors.New(cors.Config{
		AllowOrigins:           []string{"*"},
		AllowMethods:           []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:           []string{"Origin", "Authorization"},
		AllowCredentials:       true,
		ExposeHeaders:          []string{"Content-Length"},
		MaxAge:                 12 * time.Hour,
		AllowAllOrigins:        true,
		AllowWildcard:          false,
		AllowBrowserExtensions: false,
		AllowWebSockets:        false,
		AllowFiles:             false,
	})

	r.Use(config)

	r = SetupPublicRouter(r)
	r = SetupEstudianteRouter(r)
	r = SetupEvaluadorRouter(r)
	r = SetupAdministradorAcademicoRouter(r)
	r = SetupAdministradorTiRouter(r)

	return r
}
