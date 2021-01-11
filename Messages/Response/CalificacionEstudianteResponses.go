package Response

type GetOneCalificacionEstudianteResponse struct {
	Id                                      int                     `json:"id"`
	Evaluador_evaluacion                    GetOneEvaluadorResponse `json:"evaluador_evaluacion"`
	Periodo_evaluacion                      GetOnePeriodoResponse   `json:"periodo_evaluacion"`
	Puntajes_evaluacion                     []ListPuntajesResponse  `json:"puntajes_evaluacion"`
	Id_estudiante                           string                  `json:"id_estudiante"`
	Nombre_evaluacion                       string                  `json:"nombre_evaluacion"`
	Entorno_clinico_evaluacion              string                  `json:"entorno_clinico_evaluacion"`
	Paciente_evaluacion                     string                  `json:"paciente_evaluacion"`
	Asunto_principal_consulta_evaluacion    string                  `json:"asunto_principal_consulta_evaluacion"`
	Complejidad_caso_evaluacion             string                  `json:"complejidad_caso_evaluacion"`
	Numero_observaciones_previas_evaluacion string                  `json:"numero_observaciones_previas_evaluacion"`
	Categoria_observador_evaluacion         string                  `json:"categoria_observador_evaluacion"`
	Observacion_calificacion_evaluacion     string                  `json:"observacion_calificacion_evaluacion"`
	Tiempo_utilizado_evaluacion             int                     `json:"tiempo_utilizado_evaluacion"`
}

type AddNewCalificacionEstudianteResponse struct {
	Id int `json:"id"`
}
