package migrations

import (
	"fmt"
	"medroom-backend/app/Utils"
	"medroom-backend/app/models"
	"medroom-backend/app/repositories"
)

func Evaluacionmigrations() {
	fmt.Println("===== EVALUACION =====")

	var grupos []models.Grupo
	repositories.GetAllGrupos(&grupos)

	// control 1
	container := &models.Evaluacion{
		Id_grupo:          grupos[1].Id,
		Nombre_evaluacion: "CONTROL NUMERO 1",
	}
	if err := repositories.AddNewEvaluacion(container); err != nil {
		panic("NO SE PUDO MIGRAR EVALUACION 'CONTROL 1'")
	}
	Utils.StructToString(container)

	// control 2
	container = &models.Evaluacion{
		Id_grupo:          grupos[1].Id,
		Nombre_evaluacion: "CONTROL NUMERO 2",
	}
	if err := repositories.AddNewEvaluacion(container); err != nil {
		panic("NO SE PUDO MIGRAR EVALUACION 'CONTROL 2'")
	}
	Utils.StructToString(container)
}
