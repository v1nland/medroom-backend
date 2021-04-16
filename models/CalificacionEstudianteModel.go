package models

import (
	"time"

	"github.com/google/uuid"
)

type CalificacionEstudiante struct {
	Id                                               int        `json:"id" gorm:"unique;autoIncrement:true"`
	Id_estudiante                                    uuid.UUID  `json:"id_estudiante" gorm:"primaryKey;autoIncrement:false"`
	Id_evaluador                                     uuid.UUID  `json:"id_evaluador"`
	Evaluador_calificacion_estudiante                Evaluador  `json:"evaluador_calificacion_estudiante" gorm:"foreignKey:Id_evaluador"`
	Id_periodo                                       int        `json:"id_periodo" gorm:"primaryKey;autoIncrement:false"`
	Periodo_calificacion_estudiante                  Periodo    `json:"periodo_calificacion_estudiante" gorm:"foreignKey:Id_periodo"`
	Id_evaluacion                                    int        `json:"id_evaluacion" gorm:"primaryKey;autoIncrement:false"`
	Evaluacion_calificacion_estudiante               Evaluacion `json:"evaluacion_calificacion_estudiante" gorm:"foreignKey:Id_evaluacion"`
	Puntajes_calificacion_estudiante                 []Puntaje  `json:"puntajes_calificacion_estudiante" gorm:"foreignKey:Id_calificacion_estudiante;references:id"`
	Nombre_calificacion_estudiante                   string     `json:"nombre_calificacion_estudiante"`
	Observacion_calificacion_calificacion_estudiante string     `json:"observacion_calificacion_calificacion_estudiante"`
	Valoracion_general_calificacion_estudiante       int        `json:"valoracion_general_calificacion_estudiante"`
	CreatedAt                                        time.Time  `json:"created_at"`
	UpdatedAt                                        time.Time  `json:"updated_at"`
}

func (u *CalificacionEstudiante) TableName() string {
	return "public.calificaciones_estudiantes"
}
