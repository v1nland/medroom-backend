package ResponseMessages

type ListPuntajesResponse struct {
	Competencia_puntaje  GetOneCompetenciaResponse `json:"competencia_puntaje"`
	Calificacion_puntaje int                       `json:"calificacion_puntaje"`
	Feedback_puntaje     string                    `json:"feedback_puntaje"`
}

type GetOnePuntajeResponse struct {
	Competencia_puntaje  GetOneCompetenciaResponse `json:"competencia_puntaje"`
	Calificacion_puntaje int                       `json:"calificacion_puntaje"`
	Feedback_puntaje     string                    `json:"feedback_puntaje"`
}

type AddNewPuntajeResponse struct {
	Id_competencia       int    `json:"id_competencia"`
	Calificacion_puntaje int    `json:"calificacion_puntaje"`
	Feedback_puntaje     string `json:"feedback_puntaje"`
}

type PutOnePuntajeResponse struct {
	Id_competencia       int    `json:"id_competencia"`
	Calificacion_puntaje int    `json:"calificacion_puntaje"`
	Feedback_puntaje     string `json:"feedback_puntaje"`
}

type DeletePuntajeResponse struct {
	Id_competencia       int    `json:"id_competencia"`
	Calificacion_puntaje int    `json:"calificacion_puntaje"`
	Feedback_puntaje     string `json:"feedback_puntaje"`
}
