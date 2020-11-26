package RequestMessages

type ListPuntajesPayload struct {
}

type GetOnePuntajePayload struct {
}

type AddNewPuntajePayload struct {
	Id_evaluacion        int    `json:"id_evaluacion"`
	Id_competencia       int    `json:"id_competencia"`
	Calificacion_puntaje int    `json:"calificacion_puntaje"`
	Nivel_logro_puntaje  string `json:"nivel_logro_puntaje"`
}

type PutOnePuntajePayload struct {
	Id_evaluacion        int    `json:"id_evaluacion"`
	Id_competencia       int    `json:"id_competencia"`
	Calificacion_puntaje int    `json:"calificacion_puntaje"`
	Nivel_logro_puntaje  string `json:"nivel_logro_puntaje"`
}

type DeletePuntajePayload struct {
}
