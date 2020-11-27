package OutputFormats

import (
	"medroom-backend/Models"
	"medroom-backend/ResponseMessages"
)

func GetEvaluacionesOutput(u []Models.Evaluacion) (output []ResponseMessages.ListEvaluacionesResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, ResponseMessages.ListEvaluacionesResponse{
			Estudiante_evaluacion:                   GetOneEstudianteOutput(u[i].Estudiante_evaluacion),
			Evaluador_evaluacion:                    GetOneEvaluadorOutput(u[i].Evaluador_evaluacion),
			Competencia_evaluacion:                  GetOneCompetenciaOutput(u[i].Competencia_evaluacion),
			Periodo_evaluacion:                      GetOnePeriodoOutput(u[i].Periodo_evaluacion),
			Puntajes_evaluacion:                     GetPuntajesOutput(u[i].Puntajes_evaluacion),
			Nombre_evaluacion:                       u[i].Nombre_evaluacion,
			Entorno_clinico_evaluacion:              u[i].Entorno_clinico_evaluacion,
			Paciente_evaluacion:                     u[i].Paciente_evaluacion,
			Asunto_principal_consulta_evaluacion:    u[i].Asunto_principal_consulta_evaluacion,
			Complejidad_caso_evaluacion:             u[i].Complejidad_caso_evaluacion,
			Numero_observaciones_previas_evaluacion: u[i].Numero_observaciones_previas_evaluacion,
			Categoria_observador_evaluacion:         u[i].Categoria_observador_evaluacion,
			Observacion_calificacion_evaluacion:     u[i].Observacion_calificacion_evaluacion,
			Tiempo_utilizado_evaluacion:             u[i].Tiempo_utilizado_evaluacion,
		})
	}

	return output
}

func GetOneEvaluacionOutput(u Models.Evaluacion) (output ResponseMessages.GetOneEvaluacionResponse) {
	return ResponseMessages.GetOneEvaluacionResponse{
		Estudiante_evaluacion:                   GetOneEstudianteOutput(u.Estudiante_evaluacion),
		Evaluador_evaluacion:                    GetOneEvaluadorOutput(u.Evaluador_evaluacion),
		Competencia_evaluacion:                  GetOneCompetenciaOutput(u.Competencia_evaluacion),
		Periodo_evaluacion:                      GetOnePeriodoOutput(u.Periodo_evaluacion),
		Nombre_evaluacion:                       u.Nombre_evaluacion,
		Entorno_clinico_evaluacion:              u.Entorno_clinico_evaluacion,
		Paciente_evaluacion:                     u.Paciente_evaluacion,
		Asunto_principal_consulta_evaluacion:    u.Asunto_principal_consulta_evaluacion,
		Complejidad_caso_evaluacion:             u.Complejidad_caso_evaluacion,
		Numero_observaciones_previas_evaluacion: u.Numero_observaciones_previas_evaluacion,
		Categoria_observador_evaluacion:         u.Categoria_observador_evaluacion,
		Observacion_calificacion_evaluacion:     u.Observacion_calificacion_evaluacion,
		Tiempo_utilizado_evaluacion:             u.Tiempo_utilizado_evaluacion,
	}
}

func AddNewEvaluacionOutput(u Models.Evaluacion) (output ResponseMessages.AddNewEvaluacionResponse) {
	return ResponseMessages.AddNewEvaluacionResponse{
		Id_estudiante:                           u.Id_estudiante,
		Id_evaluador:                            u.Id_evaluador,
		Id_competencia:                          u.Id_competencia,
		Id_periodo:                              u.Id_periodo,
		Nombre_evaluacion:                       u.Nombre_evaluacion,
		Entorno_clinico_evaluacion:              u.Entorno_clinico_evaluacion,
		Paciente_evaluacion:                     u.Paciente_evaluacion,
		Asunto_principal_consulta_evaluacion:    u.Asunto_principal_consulta_evaluacion,
		Complejidad_caso_evaluacion:             u.Complejidad_caso_evaluacion,
		Numero_observaciones_previas_evaluacion: u.Numero_observaciones_previas_evaluacion,
		Categoria_observador_evaluacion:         u.Categoria_observador_evaluacion,
		Observacion_calificacion_evaluacion:     u.Observacion_calificacion_evaluacion,
		Tiempo_utilizado_evaluacion:             u.Tiempo_utilizado_evaluacion,
	}
}

func PutOneEvaluacionOutput(u Models.Evaluacion) (output ResponseMessages.PutOneEvaluacionResponse) {
	return ResponseMessages.PutOneEvaluacionResponse{
		Id_estudiante:                           u.Id_estudiante,
		Id_evaluador:                            u.Id_evaluador,
		Id_competencia:                          u.Id_competencia,
		Id_periodo:                              u.Id_periodo,
		Nombre_evaluacion:                       u.Nombre_evaluacion,
		Entorno_clinico_evaluacion:              u.Entorno_clinico_evaluacion,
		Paciente_evaluacion:                     u.Paciente_evaluacion,
		Asunto_principal_consulta_evaluacion:    u.Asunto_principal_consulta_evaluacion,
		Complejidad_caso_evaluacion:             u.Complejidad_caso_evaluacion,
		Numero_observaciones_previas_evaluacion: u.Numero_observaciones_previas_evaluacion,
		Categoria_observador_evaluacion:         u.Categoria_observador_evaluacion,
		Observacion_calificacion_evaluacion:     u.Observacion_calificacion_evaluacion,
		Tiempo_utilizado_evaluacion:             u.Tiempo_utilizado_evaluacion,
	}
}

func DeleteEvaluacionOutput(u Models.Evaluacion) (output ResponseMessages.DeleteEvaluacionResponse) {
	return ResponseMessages.DeleteEvaluacionResponse{
		Id_estudiante:                           u.Id_estudiante,
		Id_evaluador:                            u.Id_evaluador,
		Id_competencia:                          u.Id_competencia,
		Id_periodo:                              u.Id_periodo,
		Nombre_evaluacion:                       u.Nombre_evaluacion,
		Entorno_clinico_evaluacion:              u.Entorno_clinico_evaluacion,
		Paciente_evaluacion:                     u.Paciente_evaluacion,
		Asunto_principal_consulta_evaluacion:    u.Asunto_principal_consulta_evaluacion,
		Complejidad_caso_evaluacion:             u.Complejidad_caso_evaluacion,
		Numero_observaciones_previas_evaluacion: u.Numero_observaciones_previas_evaluacion,
		Categoria_observador_evaluacion:         u.Categoria_observador_evaluacion,
		Observacion_calificacion_evaluacion:     u.Observacion_calificacion_evaluacion,
		Tiempo_utilizado_evaluacion:             u.Tiempo_utilizado_evaluacion,
	}
}
