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
		panic("No se ha migrado rol ESTUDIANTE")
	}

	Utils.StructToString(container)

	// evaluador
	container = &Models.Rol{
		Nombre_rol: "EVALUADOR",
	}

	if err := Repositories.AddNewRol(container); err != nil {
		panic("No se ha migrado rol EVALUADOR")
	}

	Utils.StructToString(container)

	// administrador academico
	container = &Models.Rol{
		Nombre_rol: "ADMINISTRADOR ACADEMICO",
	}

	if err := Repositories.AddNewRol(container); err != nil {
		panic("No se ha migrado rol ADMINISTRADOR ACADEMICO")
	}

	Utils.StructToString(container)

	// administrador ti
	container = &Models.Rol{
		Nombre_rol: "ADMINISTRADOR TI",
	}

	if err := Repositories.AddNewRol(container); err != nil {
		panic("No se ha migrado rol ADMINISTRADOR TI")
	}

	Utils.StructToString(container)
}
