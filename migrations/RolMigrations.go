package migrations

import (
	"fmt"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
)

func Rolmigrations() {
	fmt.Println("===== ROL =====")

	// estudiante
	container := &models.Rol{
		Id:         1,
		Nombre_rol: "ESTUDIANTE",
	}

	if err := repositories.AddNewRol(container); err != nil {
		fmt.Println("NO SE PUDO MIGRAR ROL 'ESTUDIANTE'")
	}

	utils.StructToString(container)

	// evaluador
	container = &models.Rol{
		Id:         2,
		Nombre_rol: "EVALUADOR",
	}

	if err := repositories.AddNewRol(container); err != nil {
		fmt.Println("NO SE PUDO MIGRAR ROL 'EVALUADOR'")
	}

	utils.StructToString(container)

	// administrador academico
	container = &models.Rol{
		Id:         3,
		Nombre_rol: "ADMINISTRADOR ACADEMICO",
	}

	if err := repositories.AddNewRol(container); err != nil {
		fmt.Println("NO SE PUDO MIGRAR ROL 'ADMINISTRADOR ACADEMICO'")
	}

	utils.StructToString(container)

	// administrador ti
	container = &models.Rol{
		Id:         4,
		Nombre_rol: "ADMINISTRADOR TI",
	}

	if err := repositories.AddNewRol(container); err != nil {
		fmt.Println("NO SE PUDO MIGRAR ROL 'ADMINISTRADOR TI'")
	}

	utils.StructToString(container)
}
