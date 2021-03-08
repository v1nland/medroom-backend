package Request

type ListPuntajes struct {
}

type GetOnePuntaje struct {
}

type AddNewPuntaje struct {
	Id_evaluacion              int    `json:"id_evaluacion"`
	Nombre_competencia_puntaje string `json:"nombre_competencia_puntaje"`
	Codigo_competencia_puntaje string `json:"codigo_competencia_puntaje"`
	Calificacion_puntaje       int    `json:"calificacion_puntaje"`
	Feedback_puntaje           string `json:"feedback_puntaje"`
}

type PutOnePuntaje struct {
	Id_evaluacion              int    `json:"id_evaluacion"`
	Nombre_competencia_puntaje string `json:"nombre_competencia_puntaje"`
	Codigo_competencia_puntaje string `json:"codigo_competencia_puntaje"`
	Calificacion_puntaje       int    `json:"calificacion_puntaje"`
	Feedback_puntaje           string `json:"feedback_puntaje"`
}

type DeletePuntaje struct {
}
