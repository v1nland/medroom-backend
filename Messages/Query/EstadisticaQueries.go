package Query

type PromedioCalificacionEvaluacion struct {
	Id_estudiante        string  `json:"id_estudiante"`
	Nombres_estudiante   string  `json:"nombres_estudiante"`
	Apellidos_estudiante string  `json:"apellidos_estudiante"`
	Nombre_evaluacion    string  `json:"nombre_evaluacion"`
	Promedio_estudiante  float64 `json:"promedio_estudiante"`
}
