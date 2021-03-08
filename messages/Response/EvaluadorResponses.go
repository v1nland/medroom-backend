package Response

import "github.com/google/uuid"

type ListEvaluadoresResponse struct {
	Id                           uuid.UUID         `json:"id"`
	Rol_evaluador                GetOneRolResponse `json:"rol_evaluador"`
	Rut_evaluador                string            `json:"rut_evaluador"`
	Nombres_evaluador            string            `json:"nombres_evaluador"`
	Apellidos_evaluador          string            `json:"apellidos_evaluador"`
	Correo_electronico_evaluador string            `json:"correo_electronico_evaluador"`
	Telefono_fijo_evaluador      string            `json:"telefono_fijo_evaluador"`
	Telefono_celular_evaluador   string            `json:"telefono_celular_evaluador"`
	Recinto_evaluador            string            `json:"recinto_evaluador"`
	Cargo_evaluador              string            `json:"cargo_evaluador"`
}

type GetOneEvaluadorResponse struct {
	Id                           uuid.UUID         `json:"id"`
	Rol_evaluador                GetOneRolResponse `json:"rol_evaluador"`
	Rut_evaluador                string            `json:"rut_evaluador"`
	Nombres_evaluador            string            `json:"nombres_evaluador"`
	Apellidos_evaluador          string            `json:"apellidos_evaluador"`
	Correo_electronico_evaluador string            `json:"correo_electronico_evaluador"`
	Telefono_fijo_evaluador      string            `json:"telefono_fijo_evaluador"`
	Telefono_celular_evaluador   string            `json:"telefono_celular_evaluador"`
	Recinto_evaluador            string            `json:"recinto_evaluador"`
	Cargo_evaluador              string            `json:"cargo_evaluador"`
}

type GetMyEvaluadorResponse struct {
	Id                           uuid.UUID         `json:"id"`
	Rol_evaluador                GetOneRolResponse `json:"rol_evaluador"`
	Rut_evaluador                string            `json:"rut_evaluador"`
	Nombres_evaluador            string            `json:"nombres_evaluador"`
	Apellidos_evaluador          string            `json:"apellidos_evaluador"`
	Correo_electronico_evaluador string            `json:"correo_electronico_evaluador"`
	Telefono_fijo_evaluador      string            `json:"telefono_fijo_evaluador"`
	Telefono_celular_evaluador   string            `json:"telefono_celular_evaluador"`
	Recinto_evaluador            string            `json:"recinto_evaluador"`
	Cargo_evaluador              string            `json:"cargo_evaluador"`
}

type AddNewEvaluadorResponse struct {
	Id uuid.UUID `json:"id"`
}

type PutOneEvaluadorResponse struct {
	Id uuid.UUID `json:"id"`
}

type DeleteEvaluadorResponse struct {
	Id uuid.UUID `json:"id"`
}

type PutMyEvaluadorResponse struct {
	Id uuid.UUID `json:"id"`
}
