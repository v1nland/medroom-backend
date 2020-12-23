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

	if os.Getenv("MIGRATE_TABLES") == "1" {
		Config.DB.AutoMigrate(&Models.Periodo{}, &Models.Rol{}, &Models.AdministradorTi{},
			&Models.AdministradorAcademico{}, &Models.Evaluador{}, &Models.Competencia{},
			&Models.Curso{}, &Models.Grupo{}, &Models.CalificacionEstudiante{}, &Models.Estudiante{}, &Models.Puntaje{})
	}

	if os.Getenv("MIGRATE_VALUES") == "1" {
		Migrations.PeriodoMigrations()
		Migrations.RolMigrations()
		Migrations.CompetenciaMigrations()
		Migrations.AdministradorTiMigrations()
		Migrations.AdministradorAcademicoMigrations()
		Migrations.EvaluadorMigrations()
		Migrations.CursoMigrations()
		Migrations.EvaluacionMigrations()
		// Migrations.GrupoMigrations()
		// Migrations.EstudianteMigrations()
		Migrations.CalificacionEstudianteMigrations()
		// Migrations.PuntajeMigrations()
	}

	// var curso Models.Curso
	// if err := Repositories.GetOneCurso(&curso, "1"); err != nil {
	// 	panic("Curso no existe")
	// }

	// fmt.Println("=============")
	// fmt.Println("=============")
	// Utils.StructToString(curso)

	// var grupo Models.Grupo
	// if err := Repositories.GetOneGrupo(&grupo, "1"); err != nil {
	// 	panic("Grupo no existe")
	// }

	// fmt.Println("=============")
	// fmt.Println("=============")
	// Utils.StructToString(grupo)

	// var administradores_academicos []Models.AdministradorAcademico
	// if err := Repositories.GetAllAdministradoresAcademicos(&administradores_academicos); err != nil {
	// 	panic("AdministradorAcademico no existe")
	// }

	// fmt.Println("=============")
	// fmt.Println("=============")
	// Utils.StructToString(administradores_academicos[0])

	// setup router
	r := Routers.SetupRouter()

	// swagger
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	url := gin_swagger.URL(os.Getenv("SWAGGER_PROTOCOL") + "://" + os.Getenv("SWAGGER_HOST") + "/docs/v1/doc.json")
	r.GET("/docs/v1/*any", gin_swagger.WrapHandler(swagger_files.Handler, url))

	// run routes
	r.Run()
}
