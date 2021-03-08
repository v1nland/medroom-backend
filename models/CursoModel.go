package models

import (
	"time"
)

// agregar Nombre_profesor_asociado_curso
type Curso struct {
	Id                               int                      `json:"id"`
	Id_periodo                       int                      `json:"id_periodo"`
	Administradores_academicos_curso []AdministradorAcademico `json:"administradores_academicos_curso" gorm:"many2many:administradores_academicos_cursos;joinForeignKey:id_curso;joinReferences:id_administrador_academico"`
	Periodo_curso                    Periodo                  `json:"periodo_curso" gorm:"foreignKey:Id_periodo"`
	Grupos_curso                     []Grupo                  `json:"grupos_curso" gorm:"foreignKey:Id_curso;references:id"`
	Nombre_curso                     string                   `json:"nombre_curso"`
	Sigla_curso                      string                   `json:"sigla_curso" gorm:"unique;not null"`
	Estado_curso                     bool                     `json:"estado_curso"`
	CreatedAt                        time.Time                `json:"created_at"`
	UpdatedAt                        time.Time                `json:"updated_at"`
}

func (u *Curso) TableName() string {
	return "public.cursos"
}
