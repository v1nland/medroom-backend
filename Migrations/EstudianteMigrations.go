package migrations

import (
	"fmt"
	"medroom-backend/Utils"
	"medroom-backend/models"
	"medroom-backend/repositories"
)

func Estudiantemigrations() {
	fmt.Println("===== ESTUDIANTE =====")

	container := &models.Estudiante{
		Id_rol:                        1,
		Rut_estudiante:                "09.090.090-0",
		Nombres_estudiante:            "PRIMER ESTUDIANTE",
		Apellidos_estudiante:          "PRIMER APELLIDO",
		Hash_contrasena_estudiante:    "d04b98f48e8f8bcc15c6ae5ac050801cd6dcfd428fb5f9e65c4e16e7807340fa",
		Correo_electronico_estudiante: "PRIMER.ESTUDIANTE@MAIL.UDP.CL",
		Telefono_fijo_estudiante:      "12345678",
		Telefono_celular_estudiante:   "12345678",
	}

	if err := repositories.AddNewEstudiante(container); err != nil {
		panic("No se ha migrado estudiante")
	}

	Utils.StructToString(container)
}
