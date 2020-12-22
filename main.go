package main

import (
	"fmt"
	"medroom-backend/Config"
	"medroom-backend/Migrations"
	"medroom-backend/Models"
	"medroom-backend/Routers"
	"medroom-backend/docs"
	"os"

	"github.com/joho/godotenv"
	swagger_files "github.com/swaggo/files"
	gin_swagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

// @title MedRoom API
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
	// load environment
	err := godotenv.Load()
	if err != nil {
		fmt.Println("status error: ", err)
	}

	// connection string
	connection_string := os.Getenv("POSTGRESQL_CONNECTION_STRING")

	// open db
	Config.DB, err = gorm.Open(postgres.Open(connection_string), &gorm.Config{})
	if err != nil {
		fmt.Println("status error: ", err)
	}

	// migrations
	migrate_tables := os.Getenv("MIGRATE_TABLES")
	migrate_values := os.Getenv("MIGRATE_VALUES")

	if migrate_tables == "1" {
		Config.DB.AutoMigrate(&Models.Periodo{}, &Models.Rol{},
			&Models.Evaluador{}, &Models.Curso{}, &Models.Grupo{},
			&Models.Estudiante{}, &Models.Evaluacion{}, &Models.Puntaje{},
			&Models.AdministradorAcademico{}, &Models.AdministradorTi{})
	}

	if migrate_values == "1" {
		Migrations.RolMigrations()
		Migrations.PeriodoMigrations()
		Migrations.AdministradorTiMigrations()
		Migrations.AdministradorAcademicoMigrations()
		Migrations.EvaluadorMigrations()
		Migrations.CursoMigrations()
		Migrations.GrupoMigrations()
		// Migrations.EstudianteMigrations()
		Migrations.EvaluacionMigrations()
		// Migrations.PuntajeMigrations()
	}

	r := Routers.SetupRouter()

	// swagger
	swagger_protocol := os.Getenv("SWAGGER_PROTOCOL")
	swagger_host := os.Getenv("SWAGGER_HOST")

	docs.SwaggerInfo.Host = swagger_host
	url := gin_swagger.URL(swagger_protocol + "://" + swagger_host + "/docs/v1/doc.json")
	r.GET("/docs/v1/*any", gin_swagger.WrapHandler(swagger_files.Handler, url))

	// run routes
	r.Run()
}
