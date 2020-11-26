package InputFormats

import (
	"medroom-backend/RequestMessages"
	"medroom-backend/Utils"
	"strings"
)

func GetEvaluacionsInput(u *RequestMessages.ListEvaluacionsPayload) {

}

func GetOneEvaluacionInput(u *RequestMessages.GetOneEvaluacionPayload) {

}

func AddNewEvaluacionInput(u *RequestMessages.AddNewEvaluacionPayload) {
	u.Nombre_evaluacion = strings.TrimSpace(u.Nombre_evaluacion)
	u.Nombre_evaluacion = strings.ToUpper(u.Nombre_evaluacion)
	u.Nombre_evaluacion = Utils.RemoveAccents(u.Nombre_evaluacion)

	u.Entorno_clinico_evaluacion = strings.TrimSpace(u.Entorno_clinico_evaluacion)
	u.Entorno_clinico_evaluacion = strings.ToUpper(u.Entorno_clinico_evaluacion)
	u.Entorno_clinico_evaluacion = Utils.RemoveAccents(u.Entorno_clinico_evaluacion)

	u.Paciente_evaluacion = strings.TrimSpace(u.Paciente_evaluacion)
	u.Paciente_evaluacion = strings.ToUpper(u.Paciente_evaluacion)
	u.Paciente_evaluacion = Utils.RemoveAccents(u.Paciente_evaluacion)

	u.Entorno_clinico_evaluacion = strings.TrimSpace(u.Entorno_clinico_evaluacion)
	u.Entorno_clinico_evaluacion = strings.ToUpper(u.Entorno_clinico_evaluacion)
	u.Entorno_clinico_evaluacion = Utils.RemoveAccents(u.Entorno_clinico_evaluacion)

	u.Asunto_principal_consulta_evaluacion = strings.TrimSpace(u.Asunto_principal_consulta_evaluacion)
	u.Asunto_principal_consulta_evaluacion = strings.ToUpper(u.Asunto_principal_consulta_evaluacion)
	u.Asunto_principal_consulta_evaluacion = Utils.RemoveAccents(u.Asunto_principal_consulta_evaluacion)

	u.Complejidad_caso_evaluacion = strings.TrimSpace(u.Complejidad_caso_evaluacion)
	u.Complejidad_caso_evaluacion = strings.ToUpper(u.Complejidad_caso_evaluacion)
	u.Complejidad_caso_evaluacion = Utils.RemoveAccents(u.Complejidad_caso_evaluacion)

	u.Numero_observaciones_previas_evaluacion = strings.TrimSpace(u.Numero_observaciones_previas_evaluacion)
	u.Numero_observaciones_previas_evaluacion = strings.ToUpper(u.Numero_observaciones_previas_evaluacion)
	u.Numero_observaciones_previas_evaluacion = Utils.RemoveAccents(u.Numero_observaciones_previas_evaluacion)

	u.Categoria_observador_evaluacion = strings.TrimSpace(u.Categoria_observador_evaluacion)
	u.Categoria_observador_evaluacion = strings.ToUpper(u.Categoria_observador_evaluacion)
	u.Categoria_observador_evaluacion = Utils.RemoveAccents(u.Categoria_observador_evaluacion)

	u.Observacion_calificacion_evaluacion = strings.TrimSpace(u.Observacion_calificacion_evaluacion)
	u.Observacion_calificacion_evaluacion = strings.ToUpper(u.Observacion_calificacion_evaluacion)
	u.Observacion_calificacion_evaluacion = Utils.RemoveAccents(u.Observacion_calificacion_evaluacion)
}

func PutOneEvaluacionInput(u *RequestMessages.PutOneEvaluacionPayload) {
	u.Nombre_evaluacion = strings.TrimSpace(u.Nombre_evaluacion)
	u.Nombre_evaluacion = strings.ToUpper(u.Nombre_evaluacion)
	u.Nombre_evaluacion = Utils.RemoveAccents(u.Nombre_evaluacion)

	u.Entorno_clinico_evaluacion = strings.TrimSpace(u.Entorno_clinico_evaluacion)
	u.Entorno_clinico_evaluacion = strings.ToUpper(u.Entorno_clinico_evaluacion)
	u.Entorno_clinico_evaluacion = Utils.RemoveAccents(u.Entorno_clinico_evaluacion)

	u.Paciente_evaluacion = strings.TrimSpace(u.Paciente_evaluacion)
	u.Paciente_evaluacion = strings.ToUpper(u.Paciente_evaluacion)
	u.Paciente_evaluacion = Utils.RemoveAccents(u.Paciente_evaluacion)

	u.Entorno_clinico_evaluacion = strings.TrimSpace(u.Entorno_clinico_evaluacion)
	u.Entorno_clinico_evaluacion = strings.ToUpper(u.Entorno_clinico_evaluacion)
	u.Entorno_clinico_evaluacion = Utils.RemoveAccents(u.Entorno_clinico_evaluacion)

	u.Asunto_principal_consulta_evaluacion = strings.TrimSpace(u.Asunto_principal_consulta_evaluacion)
	u.Asunto_principal_consulta_evaluacion = strings.ToUpper(u.Asunto_principal_consulta_evaluacion)
	u.Asunto_principal_consulta_evaluacion = Utils.RemoveAccents(u.Asunto_principal_consulta_evaluacion)

	u.Complejidad_caso_evaluacion = strings.TrimSpace(u.Complejidad_caso_evaluacion)
	u.Complejidad_caso_evaluacion = strings.ToUpper(u.Complejidad_caso_evaluacion)
	u.Complejidad_caso_evaluacion = Utils.RemoveAccents(u.Complejidad_caso_evaluacion)

	u.Numero_observaciones_previas_evaluacion = strings.TrimSpace(u.Numero_observaciones_previas_evaluacion)
	u.Numero_observaciones_previas_evaluacion = strings.ToUpper(u.Numero_observaciones_previas_evaluacion)
	u.Numero_observaciones_previas_evaluacion = Utils.RemoveAccents(u.Numero_observaciones_previas_evaluacion)

	u.Categoria_observador_evaluacion = strings.TrimSpace(u.Categoria_observador_evaluacion)
	u.Categoria_observador_evaluacion = strings.ToUpper(u.Categoria_observador_evaluacion)
	u.Categoria_observador_evaluacion = Utils.RemoveAccents(u.Categoria_observador_evaluacion)

	u.Observacion_calificacion_evaluacion = strings.TrimSpace(u.Observacion_calificacion_evaluacion)
	u.Observacion_calificacion_evaluacion = strings.ToUpper(u.Observacion_calificacion_evaluacion)
	u.Observacion_calificacion_evaluacion = Utils.RemoveAccents(u.Observacion_calificacion_evaluacion)
}

func DeleteEvaluacionInput(u *RequestMessages.DeleteEvaluacionPayload) {

}
