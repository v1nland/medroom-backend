package Models

import (
	"time"

	"github.com/google/uuid"
)

type Estudiante struct {
	Id                            uuid.UUID                `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Id_rol                        int                      `json:"id_rol"`
	Rol_estudiante                Rol                      `json:"rol_estudiante" gorm:"foreignKey:Id_rol"`
	Calificaciones_estudiante     []CalificacionEstudiante `json:"calificaciones_estudiante" gorm:"foreignKey:Id_estudiante;references:id"`
	Grupos_estudiante             []Grupo                  `json:"grupos_estudiante" gorm:"many2many:estudiantes_grupos;joinForeignKey:id_estudiante;joinReferences:id_grupo"`
	Rut_estudiante                string                   `json:"rut_estudiante" gorm:"unique;not null"`
	Nombres_estudiante            string                   `json:"nombres_estudiante"`
	Apellidos_estudiante          string                   `json:"apellidos_estudiante"`
	Hash_contrasena_estudiante    string                   `json:"hash_contrasena_estudiante"`
	Correo_electronico_estudiante string                   `json:"correo_electronico_estudiante" gorm:"unique;not null"`
	Telefono_fijo_estudiante      string                   `json:"telefono_fijo_estudiante"`
	Telefono_celular_estudiante   string                   `json:"telefono_celular_estudiante"`
	CreatedAt                     time.Time                `json:"created_at"`
	UpdatedAt                     time.Time                `json:"updated_at"`
}

func (u *Estudiante) TableName() string {
	return "public.estudiantes"
}

// func (u *Estudiante) BeforeCreate(tx *gorm.DB) (err error) {
// 	u.Id = uuid.NewV4().String()
// 	return
// }
