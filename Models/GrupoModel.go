package Models

import (
	"time"
)

type Grupo struct {
	Id       int `json:"id"`
	Id_curso int `json:"id_curso"`
	// Evaluador_grupo    Evaluador    `json:"evaluador_grupo" gorm:"foreignKey:Id_evaluador"`
	Evaluaciones_grupo []Evaluacion `json:"evaluaciones_grupo" gorm:"foreignKey:Id_grupo;references:id"`
	Evaluadores_grupo  []Evaluador  `json:"evaluadores_grupo" gorm:"many2many:evaluadores_grupos;joinForeignKey:id_grupo;joinReferences:id_evaluador"`
	Estudiantes_grupo  []Estudiante `json:"estudiantes_grupo" gorm:"many2many:estudiantes_grupos;joinForeignKey:id_grupo;joinReferences:id_estudiante"`
	Nombre_grupo       string       `json:"nombre_grupo"`
	Sigla_grupo        string       `json:"sigla_grupo"`
	CreatedAt          time.Time    `json:"created_at"`
	UpdatedAt          time.Time    `json:"updated_at"`
}

func (u *Grupo) TableName() string {
	return "public.grupos"
}
