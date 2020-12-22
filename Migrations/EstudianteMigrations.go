package Migrations

import "fmt"

// import (
// 	"medroom-backend/Models"
// 	"medroom-backend/Repositories"
// 	"medroom-backend/Utils"
// )

func EstudianteMigrations() {
	fmt.Println("===== ESTUDIANTE =====")

	// var grupos []Models.Grupo
	// Repositories.GetAllGrupos(&grupos)

	// container := &Models.Estudiante{
	// 	Id_rol:                        1,
	// 	Rut_estudiante:                "22.222.222-2",
	// 	Nombres_estudiante:            "PRIMER ESTUDIANTE",
	// 	Apellidos_estudiante:          "PRIMER APELLIDO",
	// 	Hash_contrasena_estudiante:    "hash",
	// 	Correo_electronico_estudiante: "primer.estudiante@udp.cl",
	// 	Telefono_fijo_estudiante:      "12345678",
	// 	Telefono_celular_estudiante:   "12345678",
	// }

	// if err := Repositories.AddNewEstudiante(container); err != nil {
	// 	panic("No se ha migrado estudiante")
	// }

	// Utils.StructToString(container)
}
