package Migrations

import (
	"fmt"
	"medroom-backend/Models"
	"medroom-backend/Repositories"
	"medroom-backend/Utils"
)

func CalificacionEstudianteMigrations() {
	fmt.Println("===== CALIFICACION ESTUDIANTE =====")

	var grupos []Models.Grupo
	if err := Repositories.GetAllGrupos(&grupos); err != nil {
		panic("NO EXISTEN GRUPOS")
	}

	container := &Models.CalificacionEstudiante{
		Id_estudiante: grupos[0].Estudiantes_grupo[0].Id,
		Id_evaluador:  grupos[0].Evaluadores_grupo[0].Id,
		Id_evaluacion: 1,
		Id_periodo:    1,
		Puntajes_calificacion_estudiante: []Models.Puntaje{
			{
				Id_competencia:       "ANAM",
				Calificacion_puntaje: 9,
				Feedback_puntaje:     "BASTANTE BIEN",
			},
			{
				Id_competencia:       "EXFI",
				Calificacion_puntaje: 5,
				Feedback_puntaje:     "BASTANTE BIEN",
			},
			{
				Id_competencia:       "PROF",
				Calificacion_puntaje: 5,
				Feedback_puntaje:     "BASTANTE BIEN",
			},
			{
				Id_competencia:       "JUCL",
				Calificacion_puntaje: 4,
				Feedback_puntaje:     "BASTANTE BIEN",
			},
			{
				Id_competencia:       "HACO",
				Calificacion_puntaje: 7,
				Feedback_puntaje:     "BASTANTE BIEN",
			},
			{
				Id_competencia:       "OREF",
				Calificacion_puntaje: 1,
				Feedback_puntaje:     "BASTANTE BIEN",
			},
			{
				Id_competencia:       "VAGL",
				Calificacion_puntaje: 2,
				Feedback_puntaje:     "BASTANTE BIEN",
			},
		},
		Nombre_calificacion_estudiante:                       "CONTROL 1",
		Entorno_clinico_calificacion_estudiante:              "KINESIOLOGIA",
		Paciente_calificacion_estudiante:                     "PEDRO EL PACIENTE PEREZ",
		Asunto_principal_consulta_calificacion_estudiante:    "LESION DE TOBILLO",
		Complejidad_caso_calificacion_estudiante:             "ALTA",
		Numero_observaciones_previas_calificacion_estudiante: "5",
		Categoria_observador_calificacion_estudiante:         "COMPLEJA",
		Observacion_calificacion_calificacion_estudiante:     "MUY BIEN LOGRADO",
		Tiempo_utilizado_calificacion_estudiante:             10,
	}

	if err := Repositories.AddNewCalificacionEstudiante(container); err != nil {
		panic("NO SE PUDO MIGRAR CALIFICACION ESTUDIANTE")
	}

	Utils.StructToString(container)
}
