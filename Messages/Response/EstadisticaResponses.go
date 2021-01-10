package Response

type EvolucionEstudiantePorCompetenciaResponse struct {
	Eje_x   []string `json:"eje_x"`
	Eje_y   []int    `json:"eje_y"`
	Valores map[string][]struct {
		Puntaje_estudiante int     `json:"puntaje_estudiante"`
		Promedio_grupo     float64 `json:"promedio_grupo"`
	} `json:"valores"`
}
