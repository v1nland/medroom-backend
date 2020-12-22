package Migrations

// import (
// 	"fmt"
// )

// func GrupoMigrations() {
// 	fmt.Println("===== GRUPO =====")

// 	var evaluadores []Models.Evaluador
// 	if err := Repositories.GetAllEvaluadores(&evaluadores); err != nil {
// 		panic("Evaluadores no existe")
// 	}

// 	container := &Models.Grupo{
// 		Id_curso:     1,
// 		Id_evaluador: evaluadores[0].Id,
// 		Estudiantes_grupo: []Models.Estudiante{
// 			{
// 				Id_rol:                        1,
// 				Rut_estudiante:                "22.222.222-2",
// 				Nombres_estudiante:            "PRIMER ESTUDIANTE",
// 				Apellidos_estudiante:          "PRIMER APELLIDO",
// 				Hash_contrasena_estudiante:    "hash",
// 				Correo_electronico_estudiante: "primer.estudiante@udp.cl",
// 				Telefono_fijo_estudiante:      "12345678",
// 				Telefono_celular_estudiante:   "12345678",
// 			},
// 		},
// 		Nombre_grupo: "EQUIPO 7",
// 		Sigla_grupo:  "KAKASHI",
// 	}

// 	if err := Repositories.AddNewGrupo(container); err != nil {
// 		panic("No se ha migrado grupo")
// 	}

// 	Utils.StructToString(container)
// }
