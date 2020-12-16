package RequestMessages

type ListPuntajesPayload struct {
}

type GetOnePuntajePayload struct {
}

type AddNewPuntajePayload struct {
	Id_evaluacion              int    `json:"id_evaluacion"`
	Nombre_competencia_puntaje string `json:"nombre_competencia_puntaje"`
	Calificacion_puntaje       int    `json:"calificacion_puntaje"`
	Feedback_puntaje           string `json:"feedback_puntaje"`
}

type PutOnePuntajePayload struct {
	Id_evaluacion              int    `json:"id_evaluacion"`
	Nombre_competencia_puntaje string `json:"nombre_competencia_puntaje"`
	Calificacion_puntaje       int    `json:"calificacion_puntaje"`
	Feedback_puntaje           string `json:"feedback_puntaje"`
}

type DeletePuntajePayload struct {
}
