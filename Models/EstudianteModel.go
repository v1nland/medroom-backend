package Models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Estudiante struct {
	ID                            string       `json:"id" sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Id_rol                        int          `json:"id_rol" sql:"type:int REFERENCES public.roles(id)"`
	Rol_estudiante                Rol          `json:"rol_estudiante" gorm:"foreignKey:Id_rol"`
	Id_grupo                      int          `json:"id_grupo" sql:"type:int REFERENCES public.grupos(id)"`
	Evaluaciones_estudiante       []Evaluacion `json:"evaluaciones_estudiante" gorm:"foreignKey:Id_estudiante"`
	Rut_estudiante                string       `json:"rut_estudiante"`
	Nombres_estudiante            string       `json:"nombres_estudiante"`
	Apellidos_estudiante          string       `json:"apellidos_estudiante"`
	Hash_contrasena_estudiante    string       `json:"hash_contrasena_estudiante"`
	Correo_electronico_estudiante string       `json:"correo_electronico_estudiante"`
	Telefono_fijo_estudiante      string       `json:"telefono_fijo_estudiante"`
	Telefono_celular_estudiante   string       `json:"telefono_celular_estudiante"`
	CreatedAt                     time.Time    `json:"created_at"`
	UpdatedAt                     time.Time    `json:"updated_at"`
}

func (u *Estudiante) TableName() string {
	return "public.estudiantes"
}

func (u *Estudiante) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4().String())
}
