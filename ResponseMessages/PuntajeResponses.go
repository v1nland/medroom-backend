package ResponseMessages

type ListPuntajesResponse struct {
	// Evaluacion_puntaje   GetOneEvaluacionResponse  `json:"evaluacion_puntaje"`
	Competencia_puntaje  GetOneCompetenciaResponse `json:"competencia_puntaje"`
	Calificacion_puntaje int                       `json:"calificacion_puntaje"`
	Nivel_logro_puntaje  string                    `json:"nivel_logro_puntaje"`
}

type GetOnePuntajeResponse struct {
	// Evaluacion_puntaje   GetOneEvaluacionResponse  `json:"evaluacion_puntaje"`
	Competencia_puntaje  GetOneCompetenciaResponse `json:"competencia_puntaje"`
	Calificacion_puntaje int                       `json:"calificacion_puntaje"`
	Nivel_logro_puntaje  string                    `json:"nivel_logro_puntaje"`
}

type AddNewPuntajeResponse struct {
	Id_evaluacion        int    `json:"id_evaluacion"`
	Id_competencia       int    `json:"id_competencia"`
	Calificacion_puntaje int    `json:"calificacion_puntaje"`
	Nivel_logro_puntaje  string `json:"nivel_logro_puntaje"`
}

type PutOnePuntajeResponse struct {
	Id_evaluacion        int    `json:"id_evaluacion"`
	Id_competencia       int    `json:"id_competencia"`
	Calificacion_puntaje int    `json:"calificacion_puntaje"`
	Nivel_logro_puntaje  string `json:"nivel_logro_puntaje"`
}

type DeletePuntajeResponse struct {
	Id_evaluacion        int    `json:"id_evaluacion"`
	Id_competencia       int    `json:"id_competencia"`
	Calificacion_puntaje int    `json:"calificacion_puntaje"`
	Nivel_logro_puntaje  string `json:"nivel_logro_puntaje"`
}
