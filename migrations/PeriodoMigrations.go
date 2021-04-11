package migrations

import (
	"fmt"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
)

func Periodo() {
	fmt.Println("===== PERIODO =====")

	container := &models.Periodo{
		Id: "2021-1",
	}

	if err := repositories.AddNewPeriodo(container); err != nil {
		fmt.Println("NO SE PUDO MIGRAR PERIODO '2021-1'")
	}

	utils.StructToString(container)

	// container = &models.Periodo{
	// 	Nombre_periodo: "2020-2",
	// }

	// if err := repositories.AddNewPeriodo(container); err != nil {
	// 	panic("NO SE PUDO MIGRAR PERIODO '2020-2'")
	// }

	// utils.StructToString(container)

	// container = &models.Periodo{
	// 	Nombre_periodo: "2021-1",
	// }

	// if err := repositories.AddNewPeriodo(container); err != nil {
	// 	panic("NO SE PUDO MIGRAR PERIODO '2021-1'")
	// }

	// utils.StructToString(container)
}
