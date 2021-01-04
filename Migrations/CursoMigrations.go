package Migrations

import (
	"fmt"
	"medroom-backend/Models"
	"medroom-backend/Repositories"
	"medroom-backend/Utils"
)

func CursoMigrations() {
	fmt.Println("===== CURSO =====")

	// var evaluadores []Models.Evaluador
	// if err := Repositories.GetAllEvaluadores(&evaluadores); err != nil {
	// 	panic("NO EXISTEN EVALUADORES")
	// }

	container := &Models.Curso{
		Id_periodo: 1,
		Grupos_curso: []Models.Grupo{
			{
				Id_curso: 1,
				Evaluadores_grupo: []Models.Evaluador{
					{
						Id_rol:                       2,
						Rut_evaluador:                "11.111.111-1",
						Nombres_evaluador:            "KAKASHI",
						Apellidos_evaluador:          "HATAKE",
						Hash_contrasena_evaluador:    "hash",
						Correo_electronico_evaluador: "kakashi.hatake@mail.udp.cl",
						Telefono_fijo_evaluador:      "1234567",
						Telefono_celular_evaluador:   "1234567",
						Recinto_evaluador:            "ACADEMIA NINJA DE KONOHA",
						Cargo_evaluador:              "SENSEI",
					},
					{
						Id_rol:                       2,
						Rut_evaluador:                "55.555.555-2",
						Nombres_evaluador:            "TENZO",
						Apellidos_evaluador:          "YAMATO",
						Hash_contrasena_evaluador:    "hash",
						Correo_electronico_evaluador: "tenso.yamato@mail.udp.cl",
						Telefono_fijo_evaluador:      "1234567",
						Telefono_celular_evaluador:   "1234567",
						Recinto_evaluador:            "ACADEMIA NINJA DE KONOHA",
						Cargo_evaluador:              "ASISTENTE DE SENSEI",
					},
				},
				Estudiantes_grupo: []Models.Estudiante{
					{
						Id_rol:                        1,
						Rut_estudiante:                "44.444.444-1",
						Nombres_estudiante:            "NARUTO",
						Apellidos_estudiante:          "UZUMAKI",
						Hash_contrasena_estudiante:    "hash",
						Correo_electronico_estudiante: "naruto.uzumaki@mail.udp.cl",
						Telefono_fijo_estudiante:      "12345678",
						Telefono_celular_estudiante:   "12345678",
					},
					{
						Id_rol:                        1,
						Rut_estudiante:                "55.555.555-1",
						Nombres_estudiante:            "SASUKE",
						Apellidos_estudiante:          "UCHIHA",
						Hash_contrasena_estudiante:    "hash",
						Correo_electronico_estudiante: "sasuke.uchiha@mail.udp.cl",
						Telefono_fijo_estudiante:      "12345678",
						Telefono_celular_estudiante:   "12345678",
					},
					{
						Id_rol:                        1,
						Rut_estudiante:                "66.666.666-1",
						Nombres_estudiante:            "SAKURA",
						Apellidos_estudiante:          "HARUNO",
						Hash_contrasena_estudiante:    "hash",
						Correo_electronico_estudiante: "sakura.haruno@mail.udp.cl",
						Telefono_fijo_estudiante:      "12345678",
						Telefono_celular_estudiante:   "12345678",
					},
					{
						Id_rol:                        1,
						Rut_estudiante:                "77.777.777-1",
						Nombres_estudiante:            "SAI",
						Apellidos_estudiante:          "YAMANAKA",
						Hash_contrasena_estudiante:    "hash",
						Correo_electronico_estudiante: "sai.yamanaka@mail.udp.cl",
						Telefono_fijo_estudiante:      "12345678",
						Telefono_celular_estudiante:   "12345678",
					},
				},
				Nombre_grupo: "EQUIPO 7",
				Sigla_grupo:  "E-07",
			},
		},
		Nombre_curso: "NINJUTSU",
		Sigla_curso:  "CIT-1000",
		Estado_curso: true,
	}

	if err := Repositories.AddNewCurso(container); err != nil {
		panic("NO SE PUDO MIGRAR CURSO 'NINJUTSU'")
	}

	container = &Models.Curso{
		Id_periodo: 1,
		Grupos_curso: []Models.Grupo{
			{
				Id_curso: 1,
				Evaluadores_grupo: []Models.Evaluador{
					{
						Id_rol:                       2,
						Rut_evaluador:                "22.222.222-1",
						Nombres_evaluador:            "MIGHT",
						Apellidos_evaluador:          "GAI",
						Hash_contrasena_evaluador:    "hash",
						Correo_electronico_evaluador: "might.gai@mail.udp.cl",
						Telefono_fijo_evaluador:      "1234567",
						Telefono_celular_evaluador:   "1234567",
						Recinto_evaluador:            "ACADEMIA NINJA DE KONOHA",
						Cargo_evaluador:              "SENSEI",
					},
				},
				Estudiantes_grupo: []Models.Estudiante{
					{
						Id_rol:                        1,
						Rut_estudiante:                "88.888.888-1",
						Nombres_estudiante:            "ROCK",
						Apellidos_estudiante:          "LEE",
						Hash_contrasena_estudiante:    "hash",
						Correo_electronico_estudiante: "rock.lee@mail.udp.cl",
						Telefono_fijo_estudiante:      "12345678",
						Telefono_celular_estudiante:   "12345678",
					},
					{
						Id_rol:                        1,
						Rut_estudiante:                "99.999.999-1",
						Nombres_estudiante:            "TEN",
						Apellidos_estudiante:          "TEN",
						Hash_contrasena_estudiante:    "hash",
						Correo_electronico_estudiante: "ten.ten@mail.udp.cl",
						Telefono_fijo_estudiante:      "12345678",
						Telefono_celular_estudiante:   "12345678",
					},
					{
						Id_rol:                        1,
						Rut_estudiante:                "11.111.111-2",
						Nombres_estudiante:            "NEJI",
						Apellidos_estudiante:          "HYUGA",
						Hash_contrasena_estudiante:    "hash",
						Correo_electronico_estudiante: "neji.hyuga@mail.udp.cl",
						Telefono_fijo_estudiante:      "12345678",
						Telefono_celular_estudiante:   "12345678",
					},
				},
				Nombre_grupo: "EQUIPO 1",
				Sigla_grupo:  "E-01",
			},
		},
		Nombre_curso: "NINJUTSU AVANZADO",
		Sigla_curso:  "CIT-1001",
		Estado_curso: true,
	}

	if err := Repositories.AddNewCurso(container); err != nil {
		panic("NO SE PUDO MIGRAR CURSO 'NINJUTSU AVANZADO'")
	}

	container = &Models.Curso{
		Id_periodo: 1,
		Grupos_curso: []Models.Grupo{
			{
				Id_curso: 1,
				Evaluadores_grupo: []Models.Evaluador{
					{
						Id_rol:                       2,
						Rut_evaluador:                "33.333.333-1",
						Nombres_evaluador:            "ASUMA",
						Apellidos_evaluador:          "SARUTOBI",
						Hash_contrasena_evaluador:    "hash",
						Correo_electronico_evaluador: "asuma.sarutobi@mail.udp.cl",
						Telefono_fijo_evaluador:      "1234567",
						Telefono_celular_evaluador:   "1234567",
						Recinto_evaluador:            "ACADEMIA NINJA DE KONOHA",
						Cargo_evaluador:              "SENSEI",
					},
				},
				Estudiantes_grupo: []Models.Estudiante{
					{
						Id_rol:                        1,
						Rut_estudiante:                "22.222.222-2",
						Nombres_estudiante:            "CHOJI",
						Apellidos_estudiante:          "AKIMICHI",
						Hash_contrasena_estudiante:    "hash",
						Correo_electronico_estudiante: "choji.akimichi@mail.udp.cl",
						Telefono_fijo_estudiante:      "12345678",
						Telefono_celular_estudiante:   "12345678",
					},
					{
						Id_rol:                        1,
						Rut_estudiante:                "33.333.333-2",
						Nombres_estudiante:            "SHIKAMARU",
						Apellidos_estudiante:          "NARA",
						Hash_contrasena_estudiante:    "hash",
						Correo_electronico_estudiante: "shikamaru.nara@mail.udp.cl",
						Telefono_fijo_estudiante:      "12345678",
						Telefono_celular_estudiante:   "12345678",
					},
					{
						Id_rol:                        1,
						Rut_estudiante:                "44.444.444-2",
						Nombres_estudiante:            "INO",
						Apellidos_estudiante:          "YAMANAKA",
						Hash_contrasena_estudiante:    "hash",
						Correo_electronico_estudiante: "ino.yamanaka@mail.udp.cl",
						Telefono_fijo_estudiante:      "12345678",
						Telefono_celular_estudiante:   "12345678",
					},
				},
				Nombre_grupo: "EQUIPO 10",
				Sigla_grupo:  "E-10",
			},
		},
		Nombre_curso: "NINJUTSU MEDICO",
		Sigla_curso:  "CIT-2000",
		Estado_curso: true,
	}

	if err := Repositories.AddNewCurso(container); err != nil {
		panic("NO SE PUDO MIGRAR CURSO 'NINJUTSU MEDICO'")
	}

	Utils.StructToString(container)
}
