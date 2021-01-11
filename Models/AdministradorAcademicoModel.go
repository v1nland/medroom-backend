package Models

import (
	"time"

	"github.com/google/uuid"
)

type AdministradorAcademico struct {
	Id                                         uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Id_rol                                     int       `json:"id_rol"`
	Rol_administrador_academico                Rol       `json:"rol_administrador_academico" gorm:"foreignKey:Id_rol"`
	Rut_administrador_academico                string    `json:"rut_administrador_academico" gorm:"unique;not null"`
	Nombres_administrador_academico            string    `json:"nombres_administrador_academico"`
	Apellidos_administrador_academico          string    `json:"apellidos_administrador_academico"`
	Hash_contrasena_administrador_academico    string    `json:"hash_contrasena_administrador_academico"`
	Correo_electronico_administrador_academico string    `json:"correo_electronico_administrador_academico" gorm:"unique;not null"`
	Telefono_fijo_administrador_academico      string    `json:"telefono_fijo_administrador_academico"`
	Telefono_celular_administrador_academico   string    `json:"telefono_celular_administrador_academico"`
	CreatedAt                                  time.Time `json:"created_at"`
	UpdatedAt                                  time.Time `json:"updated_at"`
}

func (u *AdministradorAcademico) TableName() string {
	return "public.administradores_academicos"
}
