package Models

import (
	"time"
)

type Evaluacion struct {
	Id                int       `json:"id"`
	Id_grupo          int       `json:"id_grupo"`
	Nombre_evaluacion string    `json:"nombre_evaluacion"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (u *Evaluacion) TableName() string {
	return "public.evaluaciones"
}
