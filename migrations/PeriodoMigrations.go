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

	container = &models.Periodo{
		Id: "2021-2",
	}

	if err := repositories.AddNewPeriodo(container); err != nil {
		fmt.Println("NO SE PUDO MIGRAR PERIODO '2021-2'")
	}

	utils.StructToString(container)

	container = &models.Periodo{
		Id: "2022-1",
	}

	if err := repositories.AddNewPeriodo(container); err != nil {
		fmt.Println("NO SE PUDO MIGRAR PERIODO '2022-1'")
	}

	utils.StructToString(container)

	container = &models.Periodo{
		Id: "2022-2",
	}

	if err := repositories.AddNewPeriodo(container); err != nil {
		fmt.Println("NO SE PUDO MIGRAR PERIODO '2022-2'")
	}

	utils.StructToString(container)

	container = &models.Periodo{
		Id: "2023-1",
	}

	if err := repositories.AddNewPeriodo(container); err != nil {
		fmt.Println("NO SE PUDO MIGRAR PERIODO '2023-1'")
	}

	utils.StructToString(container)

	container = &models.Periodo{
		Id: "2023-2",
	}

	if err := repositories.AddNewPeriodo(container); err != nil {
		fmt.Println("NO SE PUDO MIGRAR PERIODO '2023-2'")
	}

	utils.StructToString(container)

	container = &models.Periodo{
		Id: "2024-1",
	}

	if err := repositories.AddNewPeriodo(container); err != nil {
		fmt.Println("NO SE PUDO MIGRAR PERIODO '2024-1'")
	}

	utils.StructToString(container)

	container = &models.Periodo{
		Id: "2024-2",
	}

	if err := repositories.AddNewPeriodo(container); err != nil {
		fmt.Println("NO SE PUDO MIGRAR PERIODO '2024-2'")
	}

	utils.StructToString(container)

	container = &models.Periodo{
		Id: "2025-1",
	}

	if err := repositories.AddNewPeriodo(container); err != nil {
		fmt.Println("NO SE PUDO MIGRAR PERIODO '2025-1'")
	}

	utils.StructToString(container)

	container = &models.Periodo{
		Id: "2025-2",
	}

	if err := repositories.AddNewPeriodo(container); err != nil {
		fmt.Println("NO SE PUDO MIGRAR PERIODO '2025-2'")
	}

	utils.StructToString(container)
}
