package ResponseMessages

type ListPuntajesResponse struct {
	Id_evaluacion              int    `json:"id_evaluacion"`
	Nombre_competencia_puntaje string `json:"nombre_competencia_puntaje"`
	Calificacion_puntaje       int    `json:"calificacion_puntaje"`
	Feedback_puntaje           string `json:"feedback_puntaje"`
}

type GetOnePuntajeResponse struct {
	Id_evaluacion              int    `json:"id_evaluacion"`
	Nombre_competencia_puntaje string `json:"nombre_competencia_puntaje"`
	Calificacion_puntaje       int    `json:"calificacion_puntaje"`
	Feedback_puntaje           string `json:"feedback_puntaje"`
}

type AddNewPuntajeResponse struct {
	Id_evaluacion              int    `json:"id_evaluacion"`
	Nombre_competencia_puntaje string `json:"nombre_competencia_puntaje"`
	Calificacion_puntaje       int    `json:"calificacion_puntaje"`
	Feedback_puntaje           string `json:"feedback_puntaje"`
}

type PutOnePuntajeResponse struct {
	Id_evaluacion              int    `json:"id_evaluacion"`
	Nombre_competencia_puntaje string `json:"nombre_competencia_puntaje"`
	Calificacion_puntaje       int    `json:"calificacion_puntaje"`
	Feedback_puntaje           string `json:"feedback_puntaje"`
}

type DeletePuntajeResponse struct {
	Id_evaluacion              int    `json:"id_evaluacion"`
	Nombre_competencia_puntaje string `json:"nombre_competencia_puntaje"`
	Calificacion_puntaje       int    `json:"calificacion_puntaje"`
	Feedback_puntaje           string `json:"feedback_puntaje"`
}
