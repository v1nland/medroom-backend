package Migrations

import (
	"fmt"
	"medroom-backend/Models"
	"medroom-backend/Repositories"
	"medroom-backend/Utils"
)

func CursoMigrations() {
	fmt.Println("===== CURSO =====")

	var evaluadores []Models.Evaluador
	if err := Repositories.GetAllEvaluadores(&evaluadores); err != nil {
		panic("NO EXISTEN EVALUADORES")
	}

	container := &Models.Curso{
		Id_periodo: 1,
		Grupos_curso: []Models.Grupo{
			{
				Id_curso:     1,
				Id_evaluador: evaluadores[0].Id,
				Estudiantes_grupo: []Models.Estudiante{
					{
						Id_rol:                        1,
						Rut_estudiante:                "22.222.222-2",
						Nombres_estudiante:            "PRIMER ESTUDIANTE",
						Apellidos_estudiante:          "PRIMER APELLIDO",
						Hash_contrasena_estudiante:    "hash",
						Correo_electronico_estudiante: "primer.estudiante@udp.cl",
						Telefono_fijo_estudiante:      "12345678",
						Telefono_celular_estudiante:   "12345678",
					},
				},
				Nombre_grupo: "EQUIPO 7",
				Sigla_grupo:  "KAKASHI",
			},
		},
		Nombre_curso: "PROGRAMACION",
		Sigla_curso:  "CIT-1000",
		Estado_curso: true,
	}

	if err := Repositories.AddNewCurso(container); err != nil {
		panic("NO SE PUDO MIGRAR CURSO")
	}

	Utils.StructToString(container)
}
