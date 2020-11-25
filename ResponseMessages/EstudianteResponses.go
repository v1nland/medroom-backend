package ResponseMessages

import (
	"medroom-backend/Models"
)

type ListEstudiantesResponse struct {
	Rol_estudiante                Models.Rol   `json:"rol_estudiante"`
	Grupo_estudiante              Models.Grupo `json:"grupo_estudiante"`
	Rut_estudiante                string       `json:"rut_estudiante"`
	Nombres_estudiante            string       `json:"nombres_estudiante"`
	Apellidos_estudiante          string       `json:"apellidos_estudiante"`
	Correo_electronico_estudiante string       `json:"correo_electronico_estudiante"`
	Telefono_fijo_estudiante      string       `json:"telefono_fijo_estudiante"`
	Telefono_celular_estudiante   string       `json:"telefono_celular_estudiante"`
}

type GetOneEstudianteResponse struct {
	Rol_estudiante                Models.Rol   `json:"rol_estudiante"`
	Grupo_estudiante              Models.Grupo `json:"grupo_estudiante"`
	Rut_estudiante                string       `json:"rut_estudiante"`
	Nombres_estudiante            string       `json:"nombres_estudiante"`
	Apellidos_estudiante          string       `json:"apellidos_estudiante"`
	Correo_electronico_estudiante string       `json:"correo_electronico_estudiante"`
	Telefono_fijo_estudiante      string       `json:"telefono_fijo_estudiante"`
	Telefono_celular_estudiante   string       `json:"telefono_celular_estudiante"`
}

type AddNewEstudianteResponse struct {
	Rol_estudiante                Models.Rol   `json:"rol_estudiante"`
	Grupo_estudiante              Models.Grupo `json:"grupo_estudiante"`
	Rut_estudiante                string       `json:"rut_estudiante"`
	Nombres_estudiante            string       `json:"nombres_estudiante"`
	Apellidos_estudiante          string       `json:"apellidos_estudiante"`
	Correo_electronico_estudiante string       `json:"correo_electronico_estudiante"`
	Telefono_fijo_estudiante      string       `json:"telefono_fijo_estudiante"`
	Telefono_celular_estudiante   string       `json:"telefono_celular_estudiante"`
}

type PutOneEstudianteResponse struct {
	Rol_estudiante                Models.Rol   `json:"rol_estudiante"`
	Grupo_estudiante              Models.Grupo `json:"grupo_estudiante"`
	Rut_estudiante                string       `json:"rut_estudiante"`
	Nombres_estudiante            string       `json:"nombres_estudiante"`
	Apellidos_estudiante          string       `json:"apellidos_estudiante"`
	Hash_contrasena_estudiante    string       `json:"hash_contrasena_estudiante"`
	Correo_electronico_estudiante string       `json:"correo_electronico_estudiante"`
	Telefono_fijo_estudiante      string       `json:"telefono_fijo_estudiante"`
	Telefono_celular_estudiante   string       `json:"telefono_celular_estudiante"`
}

type DeleteEstudianteResponse struct {
	Rol_estudiante                Models.Rol   `json:"rol_estudiante"`
	Grupo_estudiante              Models.Grupo `json:"grupo_estudiante"`
	Rut_estudiante                string       `json:"rut_estudiante"`
	Nombres_estudiante            string       `json:"nombres_estudiante"`
	Apellidos_estudiante          string       `json:"apellidos_estudiante"`
	Correo_electronico_estudiante string       `json:"correo_electronico_estudiante"`
	Telefono_fijo_estudiante      string       `json:"telefono_fijo_estudiante"`
	Telefono_celular_estudiante   string       `json:"telefono_celular_estudiante"`
}
