package Migrations

import (
	"fmt"
	"medroom-backend/Models"
	"medroom-backend/Repositories"
	"medroom-backend/Utils"
)

func EvaluadorMigrations() {
	fmt.Println("===== EVALUADOR =====")

	var rol Models.Rol
	if err := Repositories.GetOneRol(&rol, "2"); err != nil {
		panic("ROL EVALUADOR NO EXISTE")
	}

	container := &Models.Evaluador{
		Id_rol:                       2,
		Rut_evaluador:                "33.333.333-3",
		Nombres_evaluador:            "PRIMER EVALUADOR",
		Apellidos_evaluador:          "PRIMER APELLIDO",
		Hash_contrasena_evaluador:    "hash",
		Correo_electronico_evaluador: "primer.evaluador@udp.cl",
		Telefono_fijo_evaluador:      "1234567",
		Telefono_celular_evaluador:   "1234567",
		Recinto_evaluador:            "POSTA CENTRAL",
		Cargo_evaluador:              "MEDICO GENERAL",
	}

	if err := Repositories.AddNewEvaluador(container); err != nil {
		panic("NO SE PUDO MIGRAR EVALUADOR")
	}

	Utils.StructToString(container)
}
