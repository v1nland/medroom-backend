package migrations

import (
	"fmt"
	"medroom-backend/app/Utils"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"
)

func Rolmigrations() {
	fmt.Println("===== ROL =====")

	// estudiante
	container := &models.Rol{
		Nombre_rol: "ESTUDIANTE",
	}

	if err := repositories.AddNewRol(container); err != nil {
		panic("NO SE PUDO MIGRAR ROL 'ESTUDIANTE'")
	}

	Utils.StructToString(container)

	// evaluador
	container = &models.Rol{
		Nombre_rol: "EVALUADOR",
	}

	if err := repositories.AddNewRol(container); err != nil {
		panic("NO SE PUDO MIGRAR ROL 'EVALUADOR'")
	}

	Utils.StructToString(container)

	// administrador academico
	container = &models.Rol{
		Nombre_rol: "ADMINISTRADOR ACADEMICO",
	}

	if err := repositories.AddNewRol(container); err != nil {
		panic("NO SE PUDO MIGRAR ROL 'ADMINISTRADOR ACADEMICO'")
	}

	Utils.StructToString(container)

	// administrador ti
	container = &models.Rol{
		Nombre_rol: "ADMINISTRADOR TI",
	}

	if err := repositories.AddNewRol(container); err != nil {
		panic("NO SE PUDO MIGRAR ROL 'ADMINISTRADOR TI'")
	}

	Utils.StructToString(container)
}
