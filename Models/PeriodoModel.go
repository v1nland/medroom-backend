package Models

import (
	"time"
)

type Periodo struct {
	ID             int       `json:"id"`
	Nombre_periodo string    `json:"nombre_periodo"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (u *Periodo) TableName() string {
	return "public.periodos"
}
