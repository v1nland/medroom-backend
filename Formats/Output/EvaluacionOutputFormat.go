package Output

import (
	"medroom-backend/Messages/Response"
	"medroom-backend/Models"
)

func ListEvaluacionesOutput(u []Models.CalificacionEstudiante) (output []Response.ListEvaluacionesResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListEvaluacionesResponse{
			// Evaluador_evaluacion:                    GetOneEvaluadorOutput(u[i].Evaluador_evaluacion),
			// Periodo_evaluacion:                      GetOnePeriodoOutput(u[i].Periodo_evaluacion),
			// Puntajes_evaluacion:                     ListPuntajesOutput(u[i].Puntajes_evaluacion),
			// Id_estudiante:                           u[i].Id_estudiante,
			// Nombre_evaluacion:                       u[i].Nombre_evaluacion,
			// Entorno_clinico_evaluacion:              u[i].Entorno_clinico_evaluacion,
			// Paciente_evaluacion:                     u[i].Paciente_evaluacion,
			// Asunto_principal_consulta_evaluacion:    u[i].Asunto_principal_consulta_evaluacion,
			// Complejidad_caso_evaluacion:             u[i].Complejidad_caso_evaluacion,
			// Numero_observaciones_previas_evaluacion: u[i].Numero_observaciones_previas_evaluacion,
			// Categoria_observador_evaluacion:         u[i].Categoria_observador_evaluacion,
			// Observacion_calificacion_evaluacion:     u[i].Observacion_calificacion_evaluacion,
			// Tiempo_utilizado_evaluacion:             u[i].Tiempo_utilizado_evaluacion,
		})
	}

	return output
}

func ListEvaluacionesEstudianteOutput(u []Models.CalificacionEstudiante) (output []Response.ListEvaluacionesEstudianteResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListEvaluacionesEstudianteResponse{
			// Evaluador_evaluacion:                    GetOneEvaluadorOutput(u[i].Evaluador_evaluacion),
			// Periodo_evaluacion:                      GetOnePeriodoOutput(u[i].Periodo_evaluacion),
			// Puntajes_evaluacion:                     ListPuntajesOutput(u[i].Puntajes_evaluacion),
			// Id_estudiante:                           u[i].Id_estudiante,
			// Nombre_evaluacion:                       u[i].Nombre_evaluacion,
			// Entorno_clinico_evaluacion:              u[i].Entorno_clinico_evaluacion,
			// Paciente_evaluacion:                     u[i].Paciente_evaluacion,
			// Asunto_principal_consulta_evaluacion:    u[i].Asunto_principal_consulta_evaluacion,
			// Complejidad_caso_evaluacion:             u[i].Complejidad_caso_evaluacion,
			// Numero_observaciones_previas_evaluacion: u[i].Numero_observaciones_previas_evaluacion,
			// Categoria_observador_evaluacion:         u[i].Categoria_observador_evaluacion,
			// Observacion_calificacion_evaluacion:     u[i].Observacion_calificacion_evaluacion,
			// Tiempo_utilizado_evaluacion:             u[i].Tiempo_utilizado_evaluacion,
		})
	}

	return output
}

func GetOneEvaluacionOutput(u Models.CalificacionEstudiante) (output Response.GetOneEvaluacionResponse) {
	return Response.GetOneEvaluacionResponse{
		// Evaluador_evaluacion:                    GetOneEvaluadorOutput(u.Evaluador_evaluacion),
		// Periodo_evaluacion:                      GetOnePeriodoOutput(u.Periodo_evaluacion),
		// Id_estudiante:                           u.Id_estudiante,
		// Nombre_evaluacion:                       u.Nombre_evaluacion,
		// Entorno_clinico_evaluacion:              u.Entorno_clinico_evaluacion,
		// Paciente_evaluacion:                     u.Paciente_evaluacion,
		// Asunto_principal_consulta_evaluacion:    u.Asunto_principal_consulta_evaluacion,
		// Complejidad_caso_evaluacion:             u.Complejidad_caso_evaluacion,
		// Numero_observaciones_previas_evaluacion: u.Numero_observaciones_previas_evaluacion,
		// Categoria_observador_evaluacion:         u.Categoria_observador_evaluacion,
		// Observacion_calificacion_evaluacion:     u.Observacion_calificacion_evaluacion,
		// Tiempo_utilizado_evaluacion:             u.Tiempo_utilizado_evaluacion,
	}
}

func AddNewEvaluacionOutput(u Models.CalificacionEstudiante) (output Response.AddNewEvaluacionResponse) {
	return Response.AddNewEvaluacionResponse{
		// Id_estudiante:                           u.Id_estudiante,
		// Id_evaluador:                            u.Id_evaluador,
		// Id_periodo:                              u.Id_periodo,
		// Nombre_evaluacion:                       u.Nombre_evaluacion,
		// Entorno_clinico_evaluacion:              u.Entorno_clinico_evaluacion,
		// Paciente_evaluacion:                     u.Paciente_evaluacion,
		// Asunto_principal_consulta_evaluacion:    u.Asunto_principal_consulta_evaluacion,
		// Complejidad_caso_evaluacion:             u.Complejidad_caso_evaluacion,
		// Numero_observaciones_previas_evaluacion: u.Numero_observaciones_previas_evaluacion,
		// Categoria_observador_evaluacion:         u.Categoria_observador_evaluacion,
		// Observacion_calificacion_evaluacion:     u.Observacion_calificacion_evaluacion,
		// Tiempo_utilizado_evaluacion:             u.Tiempo_utilizado_evaluacion,
	}
}

func GenerarEvaluacionOutput(u Models.CalificacionEstudiante) (output Response.GenerarEvaluacionResponse) {
	return Response.GenerarEvaluacionResponse{
		// Puntajes_evaluacion:                     ListPuntajesOutput(u.Puntajes_evaluacion),
		// Id_estudiante:                           u.Id_estudiante,
		// Id_evaluador:                            u.Id_evaluador,
		// Id_periodo:                              u.Id_periodo,
		// Nombre_evaluacion:                       u.Nombre_evaluacion,
		// Entorno_clinico_evaluacion:              u.Entorno_clinico_evaluacion,
		// Paciente_evaluacion:                     u.Paciente_evaluacion,
		// Asunto_principal_consulta_evaluacion:    u.Asunto_principal_consulta_evaluacion,
		// Complejidad_caso_evaluacion:             u.Complejidad_caso_evaluacion,
		// Numero_observaciones_previas_evaluacion: u.Numero_observaciones_previas_evaluacion,
		// Categoria_observador_evaluacion:         u.Categoria_observador_evaluacion,
		// Observacion_calificacion_evaluacion:     u.Observacion_calificacion_evaluacion,
		// Tiempo_utilizado_evaluacion:             u.Tiempo_utilizado_evaluacion,
	}
}

func PutOneEvaluacionOutput(u Models.CalificacionEstudiante) (output Response.PutOneEvaluacionResponse) {
	return Response.PutOneEvaluacionResponse{
		// Id_estudiante:                           u.Id_estudiante,
		// Id_evaluador:                            u.Id_evaluador,
		// Id_periodo:                              u.Id_periodo,
		// Nombre_evaluacion:                       u.Nombre_evaluacion,
		// Entorno_clinico_evaluacion:              u.Entorno_clinico_evaluacion,
		// Paciente_evaluacion:                     u.Paciente_evaluacion,
		// Asunto_principal_consulta_evaluacion:    u.Asunto_principal_consulta_evaluacion,
		// Complejidad_caso_evaluacion:             u.Complejidad_caso_evaluacion,
		// Numero_observaciones_previas_evaluacion: u.Numero_observaciones_previas_evaluacion,
		// Categoria_observador_evaluacion:         u.Categoria_observador_evaluacion,
		// Observacion_calificacion_evaluacion:     u.Observacion_calificacion_evaluacion,
		// Tiempo_utilizado_evaluacion:             u.Tiempo_utilizado_evaluacion,
	}
}

func DeleteEvaluacionOutput(u Models.CalificacionEstudiante) (output Response.DeleteEvaluacionResponse) {
	return Response.DeleteEvaluacionResponse{
		// Id_estudiante:                           u.Id_estudiante,
		// Id_evaluador:                            u.Id_evaluador,
		// Id_periodo:                              u.Id_periodo,
		// Nombre_evaluacion:                       u.Nombre_evaluacion,
		// Entorno_clinico_evaluacion:              u.Entorno_clinico_evaluacion,
		// Paciente_evaluacion:                     u.Paciente_evaluacion,
		// Asunto_principal_consulta_evaluacion:    u.Asunto_principal_consulta_evaluacion,
		// Complejidad_caso_evaluacion:             u.Complejidad_caso_evaluacion,
		// Numero_observaciones_previas_evaluacion: u.Numero_observaciones_previas_evaluacion,
		// Categoria_observador_evaluacion:         u.Categoria_observador_evaluacion,
		// Observacion_calificacion_evaluacion:     u.Observacion_calificacion_evaluacion,
		// Tiempo_utilizado_evaluacion:             u.Tiempo_utilizado_evaluacion,
	}
}
