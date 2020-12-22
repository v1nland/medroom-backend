package Models

import (
	"time"
)

type Puntaje struct {
	Id                         int         `json:"id"`
	Id_competencia             string      `json:"id_competencia"`
	Competencia_puntaje        Competencia `json:"competencia_puntaje" gorm:"foreignKey:Id_competencia"`
	Id_calificacion_estudiante int         `json:"id_calificacion_estudiante"`
	Calificacion_puntaje       int         `json:"calificacion_puntaje"`
	Feedback_puntaje           string      `json:"feedback_puntaje"`
	CreatedAt                  time.Time   `json:"created_at"`
	UpdatedAt                  time.Time   `json:"updated_at"`
}

func (u *Puntaje) TableName() string {
	return "public.puntajes"
}
