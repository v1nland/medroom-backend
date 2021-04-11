package migrations

// func Grupomigrations() {
// 	fmt.Println("===== GRUPO =====")

// 	var evaluadores []models.Evaluador
// 	if err := repositories.GetAllEvaluadores(&evaluadores); err != nil {
// 		panic("Evaluadores no existe")
// 	}

// 	container := &models.Grupo{
// 		Id_curso:     1,
// 		Id_evaluador: evaluadores[0].Id,
// 		Estudiantes_grupo: []models.Estudiante{
// 			{
// 				Id_rol:                        1,
// 				Rut_estudiante:                "22.222.222-2",
// 				Nombres_estudiante:            "PRIMER ESTUDIANTE",
// 				Apellidos_estudiante:          "PRIMER APELLIDO",
// 				Hash_contrasena_estudiante:    "d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa",
// 				Correo_electronico_estudiante: "primer.estudiante@udp.cl",
// 				Telefono_fijo_estudiante:      "12345678",
// 				Telefono_celular_estudiante:   "12345678",
// 			},
// 		},
// 		Nombre_grupo: "EQUIPO 7",
// 		Sigla_grupo:  "KAKASHI",
// 	}

// 	if err := repositories.AddNewGrupo(container); err != nil {
// 		panic("No se ha migrado grupo")
// 	}

// 	utils.StructToString(container)
// }
