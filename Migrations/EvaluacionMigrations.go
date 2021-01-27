package Migrations

import (
	"fmt"
	"medroom-backend/Models"
	"medroom-backend/Repositories"
	"medroom-backend/Utils"
)

func EvaluacionMigrations() {
	fmt.Println("===== EVALUACION =====")

	var grupos []Models.Grupo
	Repositories.GetAllGrupos(&grupos)

	// control 1
	container := &Models.Evaluacion{
		Id_grupo:          grupos[1].Id,
		Nombre_evaluacion: "CONTROL NUMERO 1",
	}
	if err := Repositories.AddNewEvaluacion(container); err != nil {
		panic("NO SE PUDO MIGRAR EVALUACION 'CONTROL 1'")
	}
	Utils.StructToString(container)

	// control 2
	container = &Models.Evaluacion{
		Id_grupo:          grupos[1].Id,
		Nombre_evaluacion: "CONTROL NUMERO 2",
	}
	if err := Repositories.AddNewEvaluacion(container); err != nil {
		panic("NO SE PUDO MIGRAR EVALUACION 'CONTROL 2'")
	}
	Utils.StructToString(container)
}
