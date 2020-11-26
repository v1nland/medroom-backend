package RequestMessages

type ListPuntajesPayload struct {
}

type GetOnePuntajePayload struct {
}

type AddNewPuntajePayload struct {
	Id_competencia       int    `json:"id_competencia"`
	Calificacion_puntaje int    `json:"calificacion_puntaje"`
	Feedback_puntaje     string `json:"feedback_puntaje"`
}

type PutOnePuntajePayload struct {
	Id_competencia       int    `json:"id_competencia"`
	Calificacion_puntaje int    `json:"calificacion_puntaje"`
	Feedback_puntaje     string `json:"feedback_puntaje"`
}

type DeletePuntajePayload struct {
}
