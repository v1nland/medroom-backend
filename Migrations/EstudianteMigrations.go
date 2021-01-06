package Migrations

// import (
// 	"fmt"
// 	"medroom-backend/Models"
// 	"medroom-backend/Repositories"
// 	"medroom-backend/Utils"
// )

// func EstudianteMigrations() {
// 	fmt.Println("===== ESTUDIANTE =====")

// 	var grupos []Models.Grupo
// 	Repositories.GetAllGrupos(&grupos)

// 	container := &Models.Estudiante{
// 		Id_rol:                        1,
// 		Rut_estudiante:                "22.222.222-2",
// 		Nombres_estudiante:            "PRIMER ESTUDIANTE",
// 		Apellidos_estudiante:          "PRIMER APELLIDO",
// 		Hash_contrasena_estudiante:    "d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa",
// 		Correo_electronico_estudiante: "primer.estudiante@udp.cl",
// 		Telefono_fijo_estudiante:      "12345678",
// 		Telefono_celular_estudiante:   "12345678",
// 	}

// 	if err := Repositories.AddNewEstudiante(container); err != nil {
// 		panic("No se ha migrado estudiante")
// 	}

// 	Utils.StructToString(container)
// }
