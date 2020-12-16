package Models

import (
	"time"
)

type Grupo struct {
	ID                int          `json:"id"`
	Id_curso          int          `json:"id_curso" sql:"type:int REFERENCES public.cursos(id)"`
	Curso_grupo       Curso        `json:"curso_grupo" gorm:"foreignKey:Id_curso"`
	Id_evaluador      string       `json:"id_evaluador" sql:"type:uuid REFERENCES public.evaluadores(id)"`
	Evaluador_grupo   Evaluador    `json:"evaluador_grupo" gorm:"foreignKey:Id_evaluador"`
	Estudiantes_grupo []Estudiante `json:"estudiantes_grupo" gorm:"foreignKey:Id_grupo"`
	Nombre_grupo      string       `json:"nombre_grupo"`
	Sigla_grupo       string       `json:"sigla_grupo"`
	CreatedAt         time.Time    `json:"created_at"`
	UpdatedAt         time.Time    `json:"updated_at"`
}

func (u *Grupo) TableName() string {
	return "public.grupos"
}
