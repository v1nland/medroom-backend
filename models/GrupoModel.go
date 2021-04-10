package models

import (
	"time"
)

type Grupo struct {
	Sigla_curso        string       `json:"sigla_curso" gorm:"primaryKey;not null"`
	Id_periodo_curso   string       `json:"id_periodo_curso" gorm:"primaryKey;not null"`
	Sigla_grupo        string       `json:"sigla_grupo" gorm:"primaryKey;not null"`
	Evaluaciones_grupo []Evaluacion `json:"evaluaciones_grupo" gorm:"ForeignKey:Sigla_curso_grupo,Id_periodo_curso_grupo,Sigla_grupo;References:Sigla_curso,Id_periodo_curso,Sigla_grupo"`
	Evaluadores_grupo  []Evaluador  `json:"evaluadores_grupo" gorm:"many2many:evaluadores_grupos;joinForeignKey:Sigla_curso,Id_periodo_curso,Sigla_grupo;joinReferences:id_evaluador"`
	Estudiantes_grupo  []Estudiante `json:"estudiantes_grupo" gorm:"many2many:estudiantes_grupos;joinForeignKey:Sigla_curso,Id_periodo_curso,Sigla_grupo;joinReferences:id_estudiante"`
	Nombre_grupo       string       `json:"nombre_grupo"`
	CreatedAt          time.Time    `json:"created_at"`
	UpdatedAt          time.Time    `json:"updated_at"`
}

func (u *Grupo) TableName() string {
	return "public.grupos"
}
