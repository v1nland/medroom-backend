package Response

import "github.com/google/uuid"

type ListAdministradoresTiResponse struct {
	Id                                  uuid.UUID         `json:"id"`
	Rol_administrador_ti                GetOneRolResponse `json:"rol_administrador_ti"`
	Rut_administrador_ti                string            `json:"rut_administrador_ti"`
	Nombres_administrador_ti            string            `json:"nombres_administrador_ti"`
	Apellidos_administrador_ti          string            `json:"apellidos_administrador_ti"`
	Correo_electronico_administrador_ti string            `json:"correo_electronico_administrador_ti"`
	Telefono_fijo_administrador_ti      string            `json:"telefono_fijo_administrador_ti"`
	Telefono_celular_administrador_ti   string            `json:"telefono_celular_administrador_ti"`
}

type GetOneAdministradorTiResponse struct {
	Id                                  uuid.UUID         `json:"id"`
	Rol_administrador_ti                GetOneRolResponse `json:"rol_administrador_ti"`
	Rut_administrador_ti                string            `json:"rut_administrador_ti"`
	Nombres_administrador_ti            string            `json:"nombres_administrador_ti"`
	Apellidos_administrador_ti          string            `json:"apellidos_administrador_ti"`
	Correo_electronico_administrador_ti string            `json:"correo_electronico_administrador_ti"`
	Telefono_fijo_administrador_ti      string            `json:"telefono_fijo_administrador_ti"`
	Telefono_celular_administrador_ti   string            `json:"telefono_celular_administrador_ti"`
}

type GetMyAdministradorTiResponse struct {
	Id                                  uuid.UUID         `json:"id"`
	Rol_administrador_ti                GetOneRolResponse `json:"rol_administrador_ti"`
	Rut_administrador_ti                string            `json:"rut_administrador_ti"`
	Nombres_administrador_ti            string            `json:"nombres_administrador_ti"`
	Apellidos_administrador_ti          string            `json:"apellidos_administrador_ti"`
	Correo_electronico_administrador_ti string            `json:"correo_electronico_administrador_ti"`
	Telefono_fijo_administrador_ti      string            `json:"telefono_fijo_administrador_ti"`
	Telefono_celular_administrador_ti   string            `json:"telefono_celular_administrador_ti"`
}

type AddNewAdministradorTiResponse struct {
	Id uuid.UUID `json:"id"`
}

type PutOneAdministradorTiResponse struct {
	Id uuid.UUID `json:"id"`
}

type PutMyAdministradorTiResponse struct {
	Id uuid.UUID `json:"id"`
}

type DeleteAdministradorTiResponse struct {
	Id uuid.UUID `json:"id"`
}
