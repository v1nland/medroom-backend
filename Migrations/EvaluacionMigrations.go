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
	if err := Repositories.GetAllGrupos(&grupos); err != nil {
		panic("Grupos no existe")
	}

	container := &Models.Evaluacion{
		Id_estudiante: grupos[0].Estudiantes_grupo[0].Id,
		Id_evaluador:  grupos[0].Id_evaluador,
		Id_periodo:    1,
		Puntajes_evaluacion: []Models.Puntaje{
			{
				Nombre_competencia_puntaje: "COMPETENCIA 1",
				Codigo_competencia_puntaje: "C1",
				Calificacion_puntaje:       9,
				Feedback_puntaje:           "BASTANTE BIEN",
			},
			{
				Nombre_competencia_puntaje: "COMPETENCIA 2",
				Codigo_competencia_puntaje: "C2",
				Calificacion_puntaje:       5,
				Feedback_puntaje:           "BASTANTE BIEN",
			},
		},
		Nombre_evaluacion:                       "CONTROL 1",
		Entorno_clinico_evaluacion:              "KINESIOLOGIA",
		Paciente_evaluacion:                     "PEDRO EL PACIENTE PEREZ",
		Asunto_principal_consulta_evaluacion:    "LESION DE TOBILLO",
		Complejidad_caso_evaluacion:             "ALTA",
		Numero_observaciones_previas_evaluacion: "5",
		Categoria_observador_evaluacion:         "COMPLEJA",
		Observacion_calificacion_evaluacion:     "MUY BIEN LOGRADO",
		Tiempo_utilizado_evaluacion:             10,
	}

	if err := Repositories.AddNewEvaluacion(container); err != nil {
		panic("No se ha migrado evaluación")
	}

	Utils.StructToString(container)
}
