package Request

type GetOneCalificacionEstudiantePayload struct {
}
type AddNewCalificacionEstudiantePayload struct {
	Id_periodo                                           *int                             `json:"id_periodo"`
	Nombre_calificacion_estudiante                       *string                          `json:"nombre_calificacion_estudiante"`
	Entorno_clinico_calificacion_estudiante              *string                          `json:"entorno_clinico_calificacion_estudiante"`
	Paciente_calificacion_estudiante                     *string                          `json:"paciente_calificacion_estudiante"`
	Asunto_principal_consulta_calificacion_estudiante    *string                          `json:"asunto_principal_consulta_calificacion_estudiante"`
	Complejidad_caso_calificacion_estudiante             *string                          `json:"complejidad_caso_calificacion_estudiante"`
	Numero_observaciones_previas_calificacion_estudiante *string                          `json:"numero_observaciones_previas_calificacion_estudiante"`
	Categoria_observador_calificacion_estudiante         *string                          `json:"categoria_observador_calificacion_estudiante"`
	Observacion_calificacion_calificacion_estudiante     *string                          `json:"observacion_calificacion_calificacion_estudiante"`
	Tiempo_utilizado_calificacion_estudiante             *int                             `json:"tiempo_utilizado_calificacion_estudiante"`
	Puntajes_calificacion_estudiante                     []PuntajesCalificacionEstudiante `json:"puntajes_calificacion_estudiante"`
}

// aux types
type PuntajesCalificacionEstudiante struct {
	Id_competencia       *string `json:"id_competencia"`
	Calificacion_puntaje *int    `json:"calificacion_puntaje"`
	Feedback_puntaje     *string `json:"feedback_puntaje"`
}
