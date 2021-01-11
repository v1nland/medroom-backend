package Migrations

import (
	"fmt"
	"medroom-backend/Models"
	"medroom-backend/Repositories"
	"medroom-backend/Utils"
)

func PeriodoMigrations() {
	fmt.Println("===== PERIODO =====")

	container := &Models.Periodo{
		Nombre_periodo: "2020-1",
	}

	if err := Repositories.AddNewPeriodo(container); err != nil {
		panic("NO SE PUDO MIGRAR PERIODO '2020-1'")
	}

	Utils.StructToString(container)

	container = &Models.Periodo{
		Nombre_periodo: "2020-2",
	}

	if err := Repositories.AddNewPeriodo(container); err != nil {
		panic("NO SE PUDO MIGRAR PERIODO '2020-2'")
	}

	Utils.StructToString(container)

	container = &Models.Periodo{
		Nombre_periodo: "2021-1",
	}

	if err := Repositories.AddNewPeriodo(container); err != nil {
		panic("NO SE PUDO MIGRAR PERIODO '2021-1'")
	}

	Utils.StructToString(container)
}
