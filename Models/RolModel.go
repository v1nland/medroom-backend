package Models

import (
	"time"
)

type Rol struct {
	ID         int       `json:"id"`
	Nombre_rol string    `json:"nombre_rol"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (u *Rol) TableName() string {
	return "public.roles"
}
