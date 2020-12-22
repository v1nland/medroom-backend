package Models

import (
	"time"
)

type Puntaje struct {
	Id                         int       `json:"id"`
	Id_evaluacion              int       `json:"id_evaluacion"`
	Nombre_competencia_puntaje string    `json:"nombre_competencia_puntaje" sql:"primary_key"`
	Codigo_competencia_puntaje string    `json:"codigo_competencia_puntaje" sql:"primary_key"`
	Calificacion_puntaje       int       `json:"calificacion_puntaje"`
	Feedback_puntaje           string    `json:"feedback_puntaje"`
	CreatedAt                  time.Time `json:"created_at"`
	UpdatedAt                  time.Time `json:"updated_at"`
}

func (u *Puntaje) TableName() string {
	return "public.puntajes"
}
