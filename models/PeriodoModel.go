package models

import (
	"time"
)

type Periodo struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *Periodo) TableName() string {
	return "public.periodos"
}
