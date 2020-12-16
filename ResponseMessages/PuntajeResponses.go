package ResponseMessages

type ListPuntajesResponse struct {
	Id_evaluacion        int                       `json:"id_evaluacion"`
	Competencia_puntaje  GetOneCompetenciaResponse `json:"competencia_puntaje"`
	Calificacion_puntaje int                       `json:"calificacion_puntaje"`
	Nivel_logro_puntaje  string                    `json:"nivel_logro_puntaje"`
}

type GetOnePuntajeResponse struct {
	Id_evaluacion        int                       `json:"id_evaluacion"`
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
