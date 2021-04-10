package models

import (
	"time"
)

type Curso struct {
	Sigla_curso                      string                   `json:"sigla_curso" gorm:"primaryKey;not null"`
	Id_periodo                       string                   `json:"id_periodo" gorm:"primaryKey;not null"`
	Periodo_curso                    Periodo                  `json:"periodo_curso" gorm:"foreignKey:Id_periodo"`
	Administradores_academicos_curso []AdministradorAcademico `json:"administradores_academicos_curso" gorm:"many2many:administradores_academicos_cursos;joinForeignKey:Sigla_curso,Id_periodo;joinReferences:id_administrador_academico"`
	Grupos_curso                     []Grupo                  `json:"grupos_curso" gorm:"ForeignKey:Sigla_curso,Id_periodo_curso;References:Sigla_curso,Id_periodo"`
	Nombre_curso                     string                   `json:"nombre_curso"`
	Estado_curso                     bool                     `json:"estado_curso"`
	CreatedAt                        time.Time                `json:"created_at"`
	UpdatedAt                        time.Time                `json:"updated_at"`
}

func (u *Curso) TableName() string {
	return "public.cursos"
}
