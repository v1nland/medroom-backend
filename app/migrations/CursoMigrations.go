package migrations

import (
	"fmt"
	"medroom-backend/app/Utils"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"
)

func Cursomigrations() {
	fmt.Println("===== CURSO =====")

	container := &models.Curso{
		Id_periodo: 1,
		Administradores_academicos_curso: []models.AdministradorAcademico{
			{
				Id_rol:                                     3,
				Rut_administrador_academico:                "44.444.444-4",
				Nombres_administrador_academico:            "NOMBRE ADMINISTRADOR",
				Apellidos_administrador_academico:          "APELLIDO ACADEMICO",
				Hash_contrasena_administrador_academico:    "d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa",
				Correo_electronico_administrador_academico: "ADMINISTRADOR.ACADEMICO@MAIL.UDP.CL",
				Telefono_fijo_administrador_academico:      "12345678",
				Telefono_celular_administrador_academico:   "12345678",
			},
		},
		Grupos_curso: []models.Grupo{
			{
				Evaluaciones_grupo: []models.Evaluacion{},
				Evaluadores_grupo:  []models.Evaluador{},
				Estudiantes_grupo:  []models.Estudiante{},
				Nombre_grupo:       "SIN GRUPO",
				Sigla_grupo:        "SG",
			},
			{
				Evaluadores_grupo: []models.Evaluador{
					{
						Id_rol:                       2,
						Rut_evaluador:                "11.111.111-1",
						Nombres_evaluador:            "KAKASHI",
						Apellidos_evaluador:          "HATAKE",
						Hash_contrasena_evaluador:    "d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa",
						Correo_electronico_evaluador: "KAKASHI.HATAKE@MAIL.UDP.CL",
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
						Hash_contrasena_evaluador:    "d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa",
						Correo_electronico_evaluador: "TENSO.YAMATO@MAIL.UDP.CL",
						Telefono_fijo_evaluador:      "1234567",
						Telefono_celular_evaluador:   "1234567",
						Recinto_evaluador:            "ACADEMIA NINJA DE KONOHA",
						Cargo_evaluador:              "ASISTENTE DE SENSEI",
					},
				},
				Estudiantes_grupo: []models.Estudiante{
					{
						Id_rol:                        1,
						Rut_estudiante:                "44.444.444-1",
						Nombres_estudiante:            "NARUTO",
						Apellidos_estudiante:          "UZUMAKI",
						Hash_contrasena_estudiante:    "d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa",
						Correo_electronico_estudiante: "NARUTO.UZUMAKI@MAIL.UDP.CL",
						Telefono_fijo_estudiante:      "12345678",
						Telefono_celular_estudiante:   "12345678",
					},
					{
						Id_rol:                        1,
						Rut_estudiante:                "55.555.555-1",
						Nombres_estudiante:            "SASUKE",
						Apellidos_estudiante:          "UCHIHA",
						Hash_contrasena_estudiante:    "d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa",
						Correo_electronico_estudiante: "SASUKE.UCHIHA@MAIL.UDP.CL",
						Telefono_fijo_estudiante:      "12345678",
						Telefono_celular_estudiante:   "12345678",
					},
					{
						Id_rol:                        1,
						Rut_estudiante:                "66.666.666-1",
						Nombres_estudiante:            "SAKURA",
						Apellidos_estudiante:          "HARUNO",
						Hash_contrasena_estudiante:    "d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa",
						Correo_electronico_estudiante: "SAKURA.HARUNO@MAIL.UDP.CL",
						Telefono_fijo_estudiante:      "12345678",
						Telefono_celular_estudiante:   "12345678",
					},
					{
						Id_rol:                        1,
						Rut_estudiante:                "77.777.777-1",
						Nombres_estudiante:            "SAI",
						Apellidos_estudiante:          "YAMANAKA",
						Hash_contrasena_estudiante:    "d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa",
						Correo_electronico_estudiante: "SAI.YAMANAKA@MAIL.UDP.CL",
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

	if err := repositories.AddNewCurso(container); err != nil {
		panic("NO SE PUDO MIGRAR CURSO 'NINJUTSU'")
	}

	container = &models.Curso{
		Id_periodo: 1,
		Grupos_curso: []models.Grupo{
			{
				Evaluaciones_grupo: []models.Evaluacion{},
				Evaluadores_grupo:  []models.Evaluador{},
				Estudiantes_grupo:  []models.Estudiante{},
				Nombre_grupo:       "SIN GRUPO",
				Sigla_grupo:        "SG",
			},
			{
				Evaluadores_grupo: []models.Evaluador{
					{
						Id_rol:                       2,
						Rut_evaluador:                "22.222.222-1",
						Nombres_evaluador:            "MIGHT",
						Apellidos_evaluador:          "GAI",
						Hash_contrasena_evaluador:    "d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa",
						Correo_electronico_evaluador: "MIGHT.GAI@MAIL.UDP.CL",
						Telefono_fijo_evaluador:      "1234567",
						Telefono_celular_evaluador:   "1234567",
						Recinto_evaluador:            "ACADEMIA NINJA DE KONOHA",
						Cargo_evaluador:              "SENSEI",
					},
				},
				Estudiantes_grupo: []models.Estudiante{
					{
						Id_rol:                        1,
						Rut_estudiante:                "88.888.888-1",
						Nombres_estudiante:            "ROCK",
						Apellidos_estudiante:          "LEE",
						Hash_contrasena_estudiante:    "d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa",
						Correo_electronico_estudiante: "ROCK.LEE@MAIL.UDP.CL",
						Telefono_fijo_estudiante:      "12345678",
						Telefono_celular_estudiante:   "12345678",
					},
					{
						Id_rol:                        1,
						Rut_estudiante:                "99.999.999-1",
						Nombres_estudiante:            "TEN",
						Apellidos_estudiante:          "TEN",
						Hash_contrasena_estudiante:    "d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa",
						Correo_electronico_estudiante: "TEN.TEN@MAIL.UDP.CL",
						Telefono_fijo_estudiante:      "12345678",
						Telefono_celular_estudiante:   "12345678",
					},
					{
						Id_rol:                        1,
						Rut_estudiante:                "11.111.111-2",
						Nombres_estudiante:            "NEJI",
						Apellidos_estudiante:          "HYUGA",
						Hash_contrasena_estudiante:    "d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa",
						Correo_electronico_estudiante: "NEJI.HYUGA@MAIL.UDP.CL",
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

	if err := repositories.AddNewCurso(container); err != nil {
		panic("NO SE PUDO MIGRAR CURSO 'NINJUTSU AVANZADO'")
	}

	container = &models.Curso{
		Id_periodo: 1,
		Grupos_curso: []models.Grupo{
			{
				Evaluaciones_grupo: []models.Evaluacion{},
				Evaluadores_grupo:  []models.Evaluador{},
				Estudiantes_grupo:  []models.Estudiante{},
				Nombre_grupo:       "SIN GRUPO",
				Sigla_grupo:        "SG",
			},
			{
				Evaluadores_grupo: []models.Evaluador{
					{
						Id_rol:                       2,
						Rut_evaluador:                "33.333.333-1",
						Nombres_evaluador:            "ASUMA",
						Apellidos_evaluador:          "SARUTOBI",
						Hash_contrasena_evaluador:    "d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa",
						Correo_electronico_evaluador: "ASUMA.SARUTOBI@MAIL.UDP.CL",
						Telefono_fijo_evaluador:      "1234567",
						Telefono_celular_evaluador:   "1234567",
						Recinto_evaluador:            "ACADEMIA NINJA DE KONOHA",
						Cargo_evaluador:              "SENSEI",
					},
				},
				Estudiantes_grupo: []models.Estudiante{
					{
						Id_rol:                        1,
						Rut_estudiante:                "22.222.222-2",
						Nombres_estudiante:            "CHOJI",
						Apellidos_estudiante:          "AKIMICHI",
						Hash_contrasena_estudiante:    "d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa",
						Correo_electronico_estudiante: "CHOJI.AKIMICHI@MAIL.UDP.CL",
						Telefono_fijo_estudiante:      "12345678",
						Telefono_celular_estudiante:   "12345678",
					},
					{
						Id_rol:                        1,
						Rut_estudiante:                "33.333.333-2",
						Nombres_estudiante:            "SHIKAMARU",
						Apellidos_estudiante:          "NARA",
						Hash_contrasena_estudiante:    "d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa",
						Correo_electronico_estudiante: "SHIKAMARU.NARA@MAIL.UDP.CL",
						Telefono_fijo_estudiante:      "12345678",
						Telefono_celular_estudiante:   "12345678",
					},
					{
						Id_rol:                        1,
						Rut_estudiante:                "44.444.444-2",
						Nombres_estudiante:            "INO",
						Apellidos_estudiante:          "YAMANAKA",
						Hash_contrasena_estudiante:    "d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa",
						Correo_electronico_estudiante: "INO.YAMANAKA@MAIL.UDP.CL",
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

	if err := repositories.AddNewCurso(container); err != nil {
		panic("NO SE PUDO MIGRAR CURSO 'NINJUTSU MEDICO'")
	}

	Utils.StructToString(container)
}
