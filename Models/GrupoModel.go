package Models

import (
	"time"
)

type Grupo struct {
	ID int `json:"id"`
	// add fk evaluador
	Nombre_grupo string    `json:"nombre_grupo"`
	Sigla_grupo  string    `json:"sigla_grupo"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (u *Grupo) TableName() string {
	return "public.grupos"
}
