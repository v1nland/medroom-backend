package models

import (
	"time"
)

type Periodo struct {
	Id             int       `json:"id"`
	Nombre_periodo string    `json:"nombre_periodo" gorm:"unique;not null"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (u *Periodo) TableName() string {
	return "public.periodos"
}
