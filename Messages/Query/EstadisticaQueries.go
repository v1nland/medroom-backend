package Query

type CalificacionesEstudiantePorCompetencia struct {
	Id_evaluacion                 int     `json:"id_evaluacion"`
	Nombre_evaluacion             string  `json:"nombre_evaluacion"`
	Id_competencia                string  `json:"id_competencia"`
	Promedio_calificacion_puntaje float64 `json:"promedio_calificacion_puntaje"`
	Calificacion_puntaje          int     `json:"calificacion_puntaje"`
}
