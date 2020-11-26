package Models

import (
	"time"
)

type Puntaje struct {
	ID                   int         `json:"id"`
	Id_competencia       int         `json:"id_competencia" sql:"type:int REFERENCES public.competencias(id)"`
	Competencia_puntaje  Competencia `json:"competencia_puntaje" gorm:"foreignKey:Id_competencia"`
	Calificacion_puntaje int         `json:"calificacion_puntaje"`
	Feedback_puntaje     string      `json:"feedback_puntaje"`
	CreatedAt            time.Time   `json:"created_at"`
	UpdatedAt            time.Time   `json:"updated_at"`
}

func (u *Puntaje) TableName() string {
	return "public.puntajes"
}
