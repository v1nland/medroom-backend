package Response

type ListPuntajesResponse struct {
	Id                         int    `json:"id"`
	Id_evaluacion              int    `json:"id_evaluacion"`
	Nombre_competencia_puntaje string `json:"nombre_competencia_puntaje"`
	Codigo_competencia_puntaje string `json:"codigo_competencia_puntaje"`
	Calificacion_puntaje       int    `json:"calificacion_puntaje"`
	Feedback_puntaje           string `json:"feedback_puntaje"`
}

type GetOnePuntajeResponse struct {
	Id                         int    `json:"id"`
	Id_evaluacion              int    `json:"id_evaluacion"`
	Nombre_competencia_puntaje string `json:"nombre_competencia_puntaje"`
	Codigo_competencia_puntaje string `json:"codigo_competencia_puntaje"`
	Calificacion_puntaje       int    `json:"calificacion_puntaje"`
	Feedback_puntaje           string `json:"feedback_puntaje"`
}

type AddNewPuntajeResponse struct {
	Id int `json:"id"`
}

type PutOnePuntajeResponse struct {
	Id int `json:"id"`
}

type DeletePuntajeResponse struct {
	Id int `json:"id"`
}
