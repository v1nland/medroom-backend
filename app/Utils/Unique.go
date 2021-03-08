package utils

import (
	"medroom-backend/app/messages/Query"
)

func BuildUniqueEvaluacionesCompetencia(result []Query.CalificacionesEstudiantePorCompetencia) []string {
	var nombres_evaluaciones []string
	for i := 0; i < len(result); i++ {
		nombres_evaluaciones = append(nombres_evaluaciones, result[i].Nombre_evaluacion)
	}

	return Unique(nombres_evaluaciones)
}

func BuildUniqueEvaluacionesEvaluacion(result []Query.CalificacionesEstudiantePorEvaluacion) []string {
	var nombres_evaluaciones []string
	for i := 0; i < len(result); i++ {
		nombres_evaluaciones = append(nombres_evaluaciones, result[i].Nombre_evaluacion)
	}

	return Unique(nombres_evaluaciones)
}

func Unique(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}
