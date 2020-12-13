package ResponseMessages

type ListAdministradoresTiResponse struct {
	Rol_administrador_ti                GetOneRolResponse `json:"rol_administrador_ti"`
	Rut_administrador_ti                string            `json:"rut_administrador_ti"`
	Nombres_administrador_ti            string            `json:"nombres_administrador_ti"`
	Apellidos_administrador_ti          string            `json:"apellidos_administrador_ti"`
	Correo_electronico_administrador_ti string            `json:"correo_electronico_administrador_ti"`
	Telefono_fijo_administrador_ti      string            `json:"telefono_fijo_administrador_ti"`
	Telefono_celular_administrador_ti   string            `json:"telefono_celular_administrador_ti"`
}

type GetOneAdministradorTiResponse struct {
	Rol_administrador_ti                GetOneRolResponse `json:"rol_administrador_ti"`
	Rut_administrador_ti                string            `json:"rut_administrador_ti"`
	Nombres_administrador_ti            string            `json:"nombres_administrador_ti"`
	Apellidos_administrador_ti          string            `json:"apellidos_administrador_ti"`
	Correo_electronico_administrador_ti string            `json:"correo_electronico_administrador_ti"`
	Telefono_fijo_administrador_ti      string            `json:"telefono_fijo_administrador_ti"`
	Telefono_celular_administrador_ti   string            `json:"telefono_celular_administrador_ti"`
}

type AddNewAdministradorTiResponse struct {
	Id_rol                              int    `json:"id_rol"`
	Rut_administrador_ti                string `json:"rut_administrador_ti"`
	Nombres_administrador_ti            string `json:"nombres_administrador_ti"`
	Apellidos_administrador_ti          string `json:"apellidos_administrador_ti"`
	Correo_electronico_administrador_ti string `json:"correo_electronico_administrador_ti"`
	Telefono_fijo_administrador_ti      string `json:"telefono_fijo_administrador_ti"`
	Telefono_celular_administrador_ti   string `json:"telefono_celular_administrador_ti"`
}

type PutOneAdministradorTiResponse struct {
	Id_rol                              int    `json:"id_rol"`
	Rut_administrador_ti                string `json:"rut_administrador_ti"`
	Nombres_administrador_ti            string `json:"nombres_administrador_ti"`
	Apellidos_administrador_ti          string `json:"apellidos_administrador_ti"`
	Hash_contrasena_administrador_ti    string `json:"hash_contrasena_administrador_ti"`
	Correo_electronico_administrador_ti string `json:"correo_electronico_administrador_ti"`
	Telefono_fijo_administrador_ti      string `json:"telefono_fijo_administrador_ti"`
	Telefono_celular_administrador_ti   string `json:"telefono_celular_administrador_ti"`
}

type DeleteAdministradorTiResponse struct {
	Id_rol                              int    `json:"id_rol"`
	Rut_administrador_ti                string `json:"rut_administrador_ti"`
	Nombres_administrador_ti            string `json:"nombres_administrador_ti"`
	Apellidos_administrador_ti          string `json:"apellidos_administrador_ti"`
	Correo_electronico_administrador_ti string `json:"correo_electronico_administrador_ti"`
	Telefono_fijo_administrador_ti      string `json:"telefono_fijo_administrador_ti"`
	Telefono_celular_administrador_ti   string `json:"telefono_celular_administrador_ti"`
}
