package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	swagger_files "github.com/swaggo/files"
	gin_swagger "github.com/swaggo/gin-swagger"
	"medroom-backend/Config"
	"medroom-backend/Models"
	"medroom-backend/Routers"
	_ "medroom-backend/docs"
	"os"
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

// @host localhost:8080
// @BasePath /api/v1
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("status error: ", err)
	}

	connection_string := os.Getenv("POSTGRESQL_CONNECTION_STRING")
	should_automigrate := os.Getenv("SHOULD_AUTOMIGRATE")

	Config.DB, err = gorm.Open("postgres", connection_string)
	if err != nil {
		fmt.Println("status error: ", err)
	}
	defer Config.DB.Close()

	if should_automigrate == "1" {
		// // Inicialización de tablas en DB
		Config.DB.AutoMigrate(&Models.Rol{}, &Models.Grupo{}, &Models.Estudiante{})

		// // Las FK de las relaciones N - N con datos adicionales se deben aplicar manualmente de la siguiente forma
		// // Para relaciones 1 - N y 1 - 1, se deben especificar únicamente en el Model.go
		// Config.DB.Model(&Models.PedidoProducto{}).AddForeignKey("pedido_id", "public.pedidos(id)", "CASCADE", "CASCADE")
		// Config.DB.Model(&Models.PedidoProducto{}).AddForeignKey("producto_id", "public.productos(id)", "CASCADE", "CASCADE")
	}

	r := Routers.SetupRouter()

	// auto swagger configuration
	url := gin_swagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", gin_swagger.WrapHandler(swagger_files.Handler, url))

	r.Run()
}
