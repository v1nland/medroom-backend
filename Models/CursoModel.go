package Models

import (
	"time"
)

type Curso struct {
	ID            int       `json:"id"`
	Id_periodo    int       `json:"id_periodo" sql:"type:int REFERENCES public.periodos(id)"`
	Periodo_curso Periodo   `json:"periodo_curso" gorm:"foreignKey:Id_periodo"`
	Nombre_curso  string    `json:"nombre_curso"`
	Sigla_curso   string    `json:"sigla_curso"`
	Estado_curso  bool      `json:"estado_curso"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (u *Curso) TableName() string {
	return "public.cursos"
}
