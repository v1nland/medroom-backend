package Request

type ListEvaluacionesPayload struct {
}

type ListEvaluacionesEstudiantePayload struct {
}

type GetOneEvaluacionPayload struct {
}

type AddNewEvaluacionPayload struct {
	Id_estudiante                           string `json:"id_estudiante"`
	Id_evaluador                            string `json:"id_evaluador"`
	Id_periodo                              int    `json:"id_periodo"`
	Nombre_evaluacion                       string `json:"nombre_evaluacion"`
	Entorno_clinico_evaluacion              string `json:"entorno_clinico_evaluacion"`
	Paciente_evaluacion                     string `json:"paciente_evaluacion"`
	Asunto_principal_consulta_evaluacion    string `json:"asunto_principal_consulta_evaluacion"`
	Complejidad_caso_evaluacion             string `json:"complejidad_caso_evaluacion"`
	Numero_observaciones_previas_evaluacion string `json:"numero_observaciones_previas_evaluacion"`
	Categoria_observador_evaluacion         string `json:"categoria_observador_evaluacion"`
	Observacion_calificacion_evaluacion     string `json:"observacion_calificacion_evaluacion"`
	Tiempo_utilizado_evaluacion             int    `json:"tiempo_utilizado_evaluacion"`
}

type GenerarEvaluacionPayload struct {
	Id_estudiante                           string               `json:"id_estudiante"`
	Id_evaluador                            string               `json:"id_evaluador"`
	Id_periodo                              int                  `json:"id_periodo"`
	Nombre_evaluacion                       string               `json:"nombre_evaluacion"`
	Entorno_clinico_evaluacion              string               `json:"entorno_clinico_evaluacion"`
	Paciente_evaluacion                     string               `json:"paciente_evaluacion"`
	Asunto_principal_consulta_evaluacion    string               `json:"asunto_principal_consulta_evaluacion"`
	Complejidad_caso_evaluacion             string               `json:"complejidad_caso_evaluacion"`
	Numero_observaciones_previas_evaluacion string               `json:"numero_observaciones_previas_evaluacion"`
	Categoria_observador_evaluacion         string               `json:"categoria_observador_evaluacion"`
	Observacion_calificacion_evaluacion     string               `json:"observacion_calificacion_evaluacion"`
	Tiempo_utilizado_evaluacion             int                  `json:"tiempo_utilizado_evaluacion"`
	Puntajes_evaluacion                     []PuntajesEvaluacion `json:"puntajes_evaluacion"`
}

type PutOneEvaluacionPayload struct {
	Id_estudiante                           string `json:"id_estudiante"`
	Id_evaluador                            string `json:"id_evaluador"`
	Id_periodo                              int    `json:"id_periodo"`
	Nombre_evaluacion                       string `json:"nombre_evaluacion"`
	Entorno_clinico_evaluacion              string `json:"entorno_clinico_evaluacion"`
	Paciente_evaluacion                     string `json:"paciente_evaluacion"`
	Asunto_principal_consulta_evaluacion    string `json:"asunto_principal_consulta_evaluacion"`
	Complejidad_caso_evaluacion             string `json:"complejidad_caso_evaluacion"`
	Numero_observaciones_previas_evaluacion string `json:"numero_observaciones_previas_evaluacion"`
	Categoria_observador_evaluacion         string `json:"categoria_observador_evaluacion"`
	Observacion_calificacion_evaluacion     string `json:"observacion_calificacion_evaluacion"`
	Tiempo_utilizado_evaluacion             int    `json:"tiempo_utilizado_evaluacion"`
}

type DeleteEvaluacionPayload struct {
}

// aux entity
type PuntajesEvaluacion struct {
	Nombre_competencia   string `json:"nombre_competencia"`
	Codigo_competencia   string `json:"codigo_competencia"`
	Puntaje_competencia  int    `json:"puntaje_competencia"`
	Feedback_competencia string `json:"feedback_competencia"`
}
