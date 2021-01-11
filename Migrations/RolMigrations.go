package Migrations

import (
	"fmt"
	"medroom-backend/Models"
	"medroom-backend/Repositories"
	"medroom-backend/Utils"
)

func RolMigrations() {
	fmt.Println("===== ROL =====")

	// estudiante
	container := &Models.Rol{
		Nombre_rol: "ESTUDIANTE",
	}

	if err := Repositories.AddNewRol(container); err != nil {
		panic("NO SE PUDO MIGRAR ROL 'ESTUDIANTE'")
	}

	Utils.StructToString(container)

	// evaluador
	container = &Models.Rol{
		Nombre_rol: "EVALUADOR",
	}

	if err := Repositories.AddNewRol(container); err != nil {
		panic("NO SE PUDO MIGRAR ROL 'EVALUADOR'")
	}

	Utils.StructToString(container)

	// administrador academico
	container = &Models.Rol{
		Nombre_rol: "ADMINISTRADOR ACADEMICO",
	}

	if err := Repositories.AddNewRol(container); err != nil {
		panic("NO SE PUDO MIGRAR ROL 'ADMINISTRADOR ACADEMICO'")
	}

	Utils.StructToString(container)

	// administrador ti
	container = &Models.Rol{
		Nombre_rol: "ADMINISTRADOR TI",
	}

	if err := Repositories.AddNewRol(container); err != nil {
		panic("NO SE PUDO MIGRAR ROL 'ADMINISTRADOR TI'")
	}

	Utils.StructToString(container)
}
