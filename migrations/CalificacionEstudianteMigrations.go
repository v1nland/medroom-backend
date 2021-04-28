package migrations

import (
	"fmt"
	"medroom-backend/models"
	"medroom-backend/repositories"
	"medroom-backend/utils"
)

func CalificacionEstudiante() {
	fmt.Println("===== CALIFICACION ESTUDIANTE =====")

	var grupos []models.Grupo
	if err := repositories.GetAllGrupos(&grupos); err != nil {
		panic("NO EXISTEN GRUPOS")
	}

	container := &models.CalificacionEstudiante{
		Id_estudiante: grupos[1].Estudiantes_grupo[0].Id,
		Id_evaluador:  grupos[1].Evaluadores_grupo[0].Id,
		Id_evaluacion: 1,
		Id_periodo:    "2021-1",
		Puntajes_calificacion_estudiante: []models.Puntaje{
			{
				Id_competencia:       "ANAM",
				Calificacion_puntaje: 9,
			},
			{
				Id_competencia:       "EXFI",
				Calificacion_puntaje: 5,
			},
			{
				Id_competencia:       "PROF",
				Calificacion_puntaje: 5,
			},
			{
				Id_competencia:       "JUCL",
				Calificacion_puntaje: 4,
			},
			{
				Id_competencia:       "HACO",
				Calificacion_puntaje: 7,
			},
			{
				Id_competencia:       "OREF",
				Calificacion_puntaje: 1,
			},
		},
		Valoracion_general_calificacion_estudiante: 4,
	}

	if err := repositories.AddNewCalificacionEstudiante(container); err != nil {
		fmt.Println("NO SE PUDO MIGRAR CALIFICACION ESTUDIANTE")
	}

	utils.StructToString(container)
}
