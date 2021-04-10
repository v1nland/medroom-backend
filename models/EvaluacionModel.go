package models

import (
	"time"
)

type Evaluacion struct {
	Id                     int       `json:"id"`
	Sigla_curso_grupo      string    `json:"sigla_curso_grupo"`
	Id_periodo_curso_grupo string    `json:"id_periodo_curso_grupo"`
	Sigla_grupo            string    `json:"sigla_grupo"`
	Nombre_evaluacion      string    `json:"nombre_evaluacion"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

func (u *Evaluacion) TableName() string {
	return "public.evaluaciones"
}
