package Models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Evaluador struct {
	ID                           string    `json:"id" sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Id_rol                       int       `json:"id_rol" sql:"type:int REFERENCES public.roles(id)"`
	Rol_evaluador                Rol       `json:"rol_evaluador" gorm:"foreignKey:Id_rol"`
	Rut_evaluador                string    `json:"rut_evaluador"`
	Nombres_evaluador            string    `json:"nombres_evaluador"`
	Apellidos_evaluador          string    `json:"apellidos_evaluador"`
	Hash_contrasena_evaluador    string    `json:"hash_contrasena_evaluador"`
	Correo_electronico_evaluador string    `json:"correo_electronico_evaluador" gorm:"unique;not null"`
	Telefono_fijo_evaluador      string    `json:"telefono_fijo_evaluador"`
	Telefono_celular_evaluador   string    `json:"telefono_celular_evaluador"`
	Recinto_evaluador            string    `json:"recinto_evaluador"`
	Cargo_evaluador              string    `json:"cargo_evaluador"`
	CreatedAt                    time.Time `json:"created_at"`
	UpdatedAt                    time.Time `json:"updated_at"`
}

func (u *Evaluador) TableName() string {
	return "public.evaluadores"
}

func (u *Evaluador) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4().String())
}
