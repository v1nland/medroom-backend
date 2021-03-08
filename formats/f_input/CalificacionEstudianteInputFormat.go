package f_input

import (
	"medroom-backend/messages/Request"
	"medroom-backend/utils"
	"strings"
)

func GetOneCalificacionEstudiante(u *Request.GetOneCalificacionEstudiante) {
}

func AddNewCalificacionEstudiante(u *Request.AddNewCalificacionEstudiante) {
	if u.Nombre_calificacion_estudiante != nil {
		*u.Nombre_calificacion_estudiante = strings.TrimSpace(*u.Nombre_calificacion_estudiante)
		*u.Nombre_calificacion_estudiante = strings.ToUpper(*u.Nombre_calificacion_estudiante)
		*u.Nombre_calificacion_estudiante = utils.RemoveAccents(*u.Nombre_calificacion_estudiante)
	}

	if u.Entorno_clinico_calificacion_estudiante != nil {
		*u.Entorno_clinico_calificacion_estudiante = strings.TrimSpace(*u.Entorno_clinico_calificacion_estudiante)
		*u.Entorno_clinico_calificacion_estudiante = strings.ToUpper(*u.Entorno_clinico_calificacion_estudiante)
		*u.Entorno_clinico_calificacion_estudiante = utils.RemoveAccents(*u.Entorno_clinico_calificacion_estudiante)
	}

	if u.Paciente_calificacion_estudiante != nil {
		*u.Paciente_calificacion_estudiante = strings.TrimSpace(*u.Paciente_calificacion_estudiante)
		*u.Paciente_calificacion_estudiante = strings.ToUpper(*u.Paciente_calificacion_estudiante)
		*u.Paciente_calificacion_estudiante = utils.RemoveAccents(*u.Paciente_calificacion_estudiante)
	}

	if u.Entorno_clinico_calificacion_estudiante != nil {
		*u.Entorno_clinico_calificacion_estudiante = strings.TrimSpace(*u.Entorno_clinico_calificacion_estudiante)
		*u.Entorno_clinico_calificacion_estudiante = strings.ToUpper(*u.Entorno_clinico_calificacion_estudiante)
		*u.Entorno_clinico_calificacion_estudiante = utils.RemoveAccents(*u.Entorno_clinico_calificacion_estudiante)
	}

	if u.Asunto_principal_consulta_calificacion_estudiante != nil {
		*u.Asunto_principal_consulta_calificacion_estudiante = strings.TrimSpace(*u.Asunto_principal_consulta_calificacion_estudiante)
		*u.Asunto_principal_consulta_calificacion_estudiante = strings.ToUpper(*u.Asunto_principal_consulta_calificacion_estudiante)
		*u.Asunto_principal_consulta_calificacion_estudiante = utils.RemoveAccents(*u.Asunto_principal_consulta_calificacion_estudiante)
	}

	if u.Complejidad_caso_calificacion_estudiante != nil {
		*u.Complejidad_caso_calificacion_estudiante = strings.TrimSpace(*u.Complejidad_caso_calificacion_estudiante)
		*u.Complejidad_caso_calificacion_estudiante = strings.ToUpper(*u.Complejidad_caso_calificacion_estudiante)
		*u.Complejidad_caso_calificacion_estudiante = utils.RemoveAccents(*u.Complejidad_caso_calificacion_estudiante)
	}

	if u.Numero_observaciones_previas_calificacion_estudiante != nil {
		*u.Numero_observaciones_previas_calificacion_estudiante = strings.TrimSpace(*u.Numero_observaciones_previas_calificacion_estudiante)
		*u.Numero_observaciones_previas_calificacion_estudiante = strings.ToUpper(*u.Numero_observaciones_previas_calificacion_estudiante)
		*u.Numero_observaciones_previas_calificacion_estudiante = utils.RemoveAccents(*u.Numero_observaciones_previas_calificacion_estudiante)
	}

	if u.Categoria_observador_calificacion_estudiante != nil {
		*u.Categoria_observador_calificacion_estudiante = strings.TrimSpace(*u.Categoria_observador_calificacion_estudiante)
		*u.Categoria_observador_calificacion_estudiante = strings.ToUpper(*u.Categoria_observador_calificacion_estudiante)
		*u.Categoria_observador_calificacion_estudiante = utils.RemoveAccents(*u.Categoria_observador_calificacion_estudiante)
	}

	if u.Observacion_calificacion_calificacion_estudiante != nil {
		*u.Observacion_calificacion_calificacion_estudiante = strings.TrimSpace(*u.Observacion_calificacion_calificacion_estudiante)
		*u.Observacion_calificacion_calificacion_estudiante = strings.ToUpper(*u.Observacion_calificacion_calificacion_estudiante)
		*u.Observacion_calificacion_calificacion_estudiante = utils.RemoveAccents(*u.Observacion_calificacion_calificacion_estudiante)
	}

	if u.Puntajes_calificacion_estudiante != nil {
		for i := 0; i < len(u.Puntajes_calificacion_estudiante); i++ {
			if u.Puntajes_calificacion_estudiante[i].Feedback_puntaje != nil {
				*u.Puntajes_calificacion_estudiante[i].Feedback_puntaje = strings.TrimSpace(*u.Puntajes_calificacion_estudiante[i].Feedback_puntaje)
				*u.Puntajes_calificacion_estudiante[i].Feedback_puntaje = strings.ToUpper(*u.Puntajes_calificacion_estudiante[i].Feedback_puntaje)
				*u.Puntajes_calificacion_estudiante[i].Feedback_puntaje = utils.RemoveAccents(*u.Puntajes_calificacion_estudiante[i].Feedback_puntaje)
			}
		}
	}
}
