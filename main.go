package main

import (
	"fmt"
	"medroom-backend/config"
	"medroom-backend/docs"
	"medroom-backend/migrations"
	"medroom-backend/models"
	"medroom-backend/routers"

	swagger_files "github.com/swaggo/files"
	gin_swagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

// @title medRoom
// @version 1.0
// @description Swagger definition for MedRoom backend.
// @termsOfService https://swagger.io/terms/

// @contact.name Martín Saavedra
// @contact.url https://www.swagger.io/support
// @contact.email martin.saavedra@mail.udp.cl

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1
func main() {
	config.LoadSettings("smartSide", "microservice")

	connection_string := config.GetString("POSTGRESQL_CONNECTION_STRING")

	config.DB, err = gorm.Open(postgres.Open(connection_string), &gorm.Config{})
	if err != nil {
		fmt.Println("status error: ", err)
	}

	if config.GetBool("MIGRATE_TABLES") {
		config.DB.AutoMigrate(&models.Periodo{}, &models.Rol{}, &models.AdministradorTi{},
			&models.AdministradorAcademico{}, &models.Evaluador{}, &models.Competencia{},
			&models.Curso{}, &models.Grupo{}, &models.CalificacionEstudiante{}, &models.Estudiante{}, &models.Puntaje{})
	}

	if config.GetBool("MIGRATE_VALUES") {
		migrations.Periodo()
		migrations.Rol()
		migrations.Competencia()
		migrations.AdministradorTi()
		// migrations.AdministradorAcademico()
		// migrations.Evaluadormigrations()
		// migrations.Curso()
		// migrations.Evaluacion()
		// migrations.Grupo()
		// migrations.Estudiante()
		// migrations.CalificacionEstudiante()
		// migrations.Puntaje()
	}

	// setup router
	r := routers.SetupRouter()

	// swagger
	docs.SwaggerInfo.Host = config.GetString("SWAGGER_HOST")
	url := gin_swagger.URL(config.GetString("SWAGGER_PROTOCOL") + "://" + config.GetString("SWAGGER_HOST") + "/docs/v1/doc.json")
	r.GET("/docs/v1/*any", gin_swagger.WrapHandler(swagger_files.Handler, url))

	// run routes
	r.Run()
}
