package Query

type CalificacionesEstudiantePorCompetencia struct {
	Id_evaluacion                       int     `json:"id_evaluacion"`
	Nombre_evaluacion                   string  `json:"nombre_evaluacion"`
	Id_competencia                      string  `json:"id_competencia"`
	Promedio_calificacion_puntaje_grupo float64 `json:"promedio_calificacion_puntaje_grupo"`
	Calificacion_puntaje_estudiante     int     `json:"calificacion_puntaje_estudiante"`
}

type CalificacionesEstudiantePorEvaluacion struct {
	Id_evaluacion                            int     `json:"id_evaluacion"`
	Nombre_evaluacion                        string  `json:"nombre_evaluacion"`
	Promedio_calificacion_puntaje_grupo      float64 `json:"promedio_calificacion_puntaje_grupo"`
	Promedio_calificacion_puntaje_estudiante float64 `json:"promedio_calificacion_puntaje_estudiante"`
}

type CalificacionesGrupoPorCompetencia struct {
	Id_evaluacion                       int     `json:"id_evaluacion"`
	Nombre_evaluacion                   string  `json:"nombre_evaluacion"`
	Id_competencia                      string  `json:"id_competencia"`
	Promedio_calificacion_puntaje_grupo float64 `json:"promedio_calificacion_puntaje_grupo"`
}

type CalificacionesGrupoPorEvaluacion struct {
	Id_evaluacion                       int     `json:"id_evaluacion"`
	Nombre_evaluacion                   string  `json:"nombre_evaluacion"`
	Promedio_calificacion_puntaje_grupo float64 `json:"promedio_calificacion_puntaje_grupo"`
}
