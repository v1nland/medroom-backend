package Response

type PromedioCalificacionEvaluacionResponse struct {
	Nombre_evaluacion   string  `json:"nombre_evaluacion"`
	Promedio_estudiante float64 `json:"promedio_estudiante"`
	Promedio_grupo      float64 `json:"promedio_grupo"`
	Promedio_curso      float64 `json:"promedio_curso"`
}
