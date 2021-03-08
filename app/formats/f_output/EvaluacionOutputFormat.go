package f_output

import (
	"medroom-backend/app/Messages/Response"
	"medroom-backend/app/models"
)

func ListEvaluacionesGrupoEstudiante(u []models.CalificacionEstudiante) (output []Response.ListEvaluacionesGrupoEstudianteResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListEvaluacionesGrupoEstudianteResponse{
			// Evaluador_evaluacion:                    GetOneEvaluadorf_output(u[i].Evaluador_evaluacion),
			// Periodo_evaluacion:                      GetOnePeriodof_output(u[i].Periodo_evaluacion),
			// Puntajes_evaluacion:                     ListPuntajesf_output(u[i].Puntajes_evaluacion),
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

func ListEvaluacionesGrupoEvaluador(u []models.CalificacionEstudiante) (output []Response.ListEvaluacionesGrupoEvaluadorResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListEvaluacionesGrupoEvaluadorResponse{
			// Evaluador_evaluacion:                    GetOneEvaluadorf_output(u[i].Evaluador_evaluacion),
			// Periodo_evaluacion:                      GetOnePeriodof_output(u[i].Periodo_evaluacion),
			// Puntajes_evaluacion:                     ListPuntajesf_output(u[i].Puntajes_evaluacion),
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

func ListEvaluaciones(u []models.CalificacionEstudiante) (output []Response.ListEvaluacionesResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListEvaluacionesResponse{
			// Evaluador_evaluacion:                    GetOneEvaluadorf_output(u[i].Evaluador_evaluacion),
			// Periodo_evaluacion:                      GetOnePeriodof_output(u[i].Periodo_evaluacion),
			// Puntajes_evaluacion:                     ListPuntajesf_output(u[i].Puntajes_evaluacion),
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

func ListEvaluacionesEstudiante(u []models.CalificacionEstudiante) (output []Response.ListEvaluacionesEstudianteResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListEvaluacionesEstudianteResponse{
			// Evaluador_evaluacion:                    GetOneEvaluadorf_output(u[i].Evaluador_evaluacion),
			// Periodo_evaluacion:                      GetOnePeriodof_output(u[i].Periodo_evaluacion),
			// Puntajes_evaluacion:                     ListPuntajesf_output(u[i].Puntajes_evaluacion),
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

func GetOneEvaluacion(u models.CalificacionEstudiante) (output Response.GetOneEvaluacionResponse) {
	return Response.GetOneEvaluacionResponse{
		// Evaluador_evaluacion:                    GetOneEvaluadorf_output(u.Evaluador_evaluacion),
		// Periodo_evaluacion:                      GetOnePeriodof_output(u.Periodo_evaluacion),
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

func AddNewEvaluacion(u models.CalificacionEstudiante) (output Response.AddNewEvaluacionResponse) {
	return Response.AddNewEvaluacionResponse{
		Id: u.Id,
	}
}

func PutOneEvaluacion(u models.CalificacionEstudiante) (output Response.PutOneEvaluacionResponse) {
	return Response.PutOneEvaluacionResponse{
		Id: u.Id,
	}
}

func DeleteEvaluacion(u models.CalificacionEstudiante) (output Response.DeleteEvaluacionResponse) {
	return Response.DeleteEvaluacionResponse{
		Id: u.Id,
	}
}
