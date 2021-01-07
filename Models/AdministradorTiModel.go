package Models

import (
	"time"

	"github.com/google/uuid"
)

type AdministradorTi struct {
	Id                                  uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Id_rol                              int       `json:"id_rol"`
	Rol_administrador_ti                Rol       `json:"rol_administrador_ti" gorm:"foreignKey:Id_rol"`
	Rut_administrador_ti                string    `json:"rut_administrador_ti" gorm:"unique;not null"`
	Nombres_administrador_ti            string    `json:"nombres_administrador_ti"`
	Apellidos_administrador_ti          string    `json:"apellidos_administrador_ti"`
	Hash_contrasena_administrador_ti    string    `json:"hash_contrasena_administrador_ti"`
	Correo_electronico_administrador_ti string    `json:"correo_electronico_administrador_ti" gorm:"unique;not null"`
	Telefono_fijo_administrador_ti      string    `json:"telefono_fijo_administrador_ti"`
	Telefono_celular_administrador_ti   string    `json:"telefono_celular_administrador_ti"`
	CreatedAt                           time.Time `json:"created_at"`
	UpdatedAt                           time.Time `json:"updated_at"`
}

func (u *AdministradorTi) TableName() string {
	return "public.administradores_ti"
}
