package migrations

import (
	"fmt"
	"medroom-backend/app/Utils"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"
)

func Periodomigrations() {
	fmt.Println("===== PERIODO =====")

	container := &models.Periodo{
		Nombre_periodo: "2020-1",
	}

	if err := repositories.AddNewPeriodo(container); err != nil {
		panic("NO SE PUDO MIGRAR PERIODO '2020-1'")
	}

	Utils.StructToString(container)

	container = &models.Periodo{
		Nombre_periodo: "2020-2",
	}

	if err := repositories.AddNewPeriodo(container); err != nil {
		panic("NO SE PUDO MIGRAR PERIODO '2020-2'")
	}

	Utils.StructToString(container)

	container = &models.Periodo{
		Nombre_periodo: "2021-1",
	}

	if err := repositories.AddNewPeriodo(container); err != nil {
		panic("NO SE PUDO MIGRAR PERIODO '2021-1'")
	}

	Utils.StructToString(container)
}
