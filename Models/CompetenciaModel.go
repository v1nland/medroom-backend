package Models

import (
	"time"
)

type Competencia struct {
	ID                 int       `json:"id"`
	Nombre_competencia string    `json:"nombre_competencia"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func (u *Competencia) TableName() string {
	return "public.competencias"
}
