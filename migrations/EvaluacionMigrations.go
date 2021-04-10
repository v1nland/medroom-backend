package migrations

import (
	"fmt"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
)

func Evaluacionmigrations() {
	fmt.Println("===== EVALUACION =====")

	var grupos []models.Grupo
	repositories.GetAllGrupos(&grupos)

	// control 1
	container := &models.Evaluacion{
		Sigla_curso_grupo:      "CIT-1000",
		Id_periodo_curso_grupo: "2021-1",
		Sigla_grupo:            "E-07",
		Nombre_evaluacion:      "EVALUACION NUMERO 1",
	}
	if err := repositories.AddNewEvaluacion(container); err != nil {
		panic("NO SE PUDO MIGRAR EVALUACION 'CONTROL 1'")
	}
	utils.StructToString(container)

	// control 2
	container = &models.Evaluacion{
		Sigla_curso_grupo:      "CIT-1000",
		Id_periodo_curso_grupo: "2021-1",
		Sigla_grupo:            "E-07",
		Nombre_evaluacion:      "EVALUACION NUMERO 2",
	}
	if err := repositories.AddNewEvaluacion(container); err != nil {
		panic("NO SE PUDO MIGRAR EVALUACION 'CONTROL 2'")
	}
	utils.StructToString(container)
}
