package ResponseMessages

type ListEvaluadoresResponse struct {
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
	Id_rol                       int    `json:"id_rol"`
	Rut_evaluador                string `json:"rut_evaluador"`
	Nombres_evaluador            string `json:"nombres_evaluador"`
	Apellidos_evaluador          string `json:"apellidos_evaluador"`
	Correo_electronico_evaluador string `json:"correo_electronico_evaluador"`
	Telefono_fijo_evaluador      string `json:"telefono_fijo_evaluador"`
	Telefono_celular_evaluador   string `json:"telefono_celular_evaluador"`
	Recinto_evaluador            string `json:"recinto_evaluador"`
	Cargo_evaluador              string `json:"cargo_evaluador"`
}

type PutOneEvaluadorResponse struct {
	Id_rol                       int    `json:"id_rol"`
	Rut_evaluador                string `json:"rut_evaluador"`
	Nombres_evaluador            string `json:"nombres_evaluador"`
	Apellidos_evaluador          string `json:"apellidos_evaluador"`
	Hash_contrasena_evaluador    string `json:"hash_contrasena_evaluador"`
	Correo_electronico_evaluador string `json:"correo_electronico_evaluador"`
	Telefono_fijo_evaluador      string `json:"telefono_fijo_evaluador"`
	Telefono_celular_evaluador   string `json:"telefono_celular_evaluador"`
	Recinto_evaluador            string `json:"recinto_evaluador"`
	Cargo_evaluador              string `json:"cargo_evaluador"`
}

type DeleteEvaluadorResponse struct {
	Id_rol                       int    `json:"id_rol"`
	Rut_evaluador                string `json:"rut_evaluador"`
	Nombres_evaluador            string `json:"nombres_evaluador"`
	Apellidos_evaluador          string `json:"apellidos_evaluador"`
	Correo_electronico_evaluador string `json:"correo_electronico_evaluador"`
	Telefono_fijo_evaluador      string `json:"telefono_fijo_evaluador"`
	Telefono_celular_evaluador   string `json:"telefono_celular_evaluador"`
	Recinto_evaluador            string `json:"recinto_evaluador"`
	Cargo_evaluador              string `json:"cargo_evaluador"`
}

type GetMyEvaluadorResponse struct {
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

type PutMyEvaluadorResponse struct {
	Id_rol                       int    `json:"id_rol"`
	Rut_evaluador                string `json:"rut_evaluador"`
	Nombres_evaluador            string `json:"nombres_evaluador"`
	Apellidos_evaluador          string `json:"apellidos_evaluador"`
	Hash_contrasena_evaluador    string `json:"hash_contrasena_evaluador"`
	Correo_electronico_evaluador string `json:"correo_electronico_evaluador"`
	Telefono_fijo_evaluador      string `json:"telefono_fijo_evaluador"`
	Telefono_celular_evaluador   string `json:"telefono_celular_evaluador"`
	Recinto_evaluador            string `json:"recinto_evaluador"`
	Cargo_evaluador              string `json:"cargo_evaluador"`
}
