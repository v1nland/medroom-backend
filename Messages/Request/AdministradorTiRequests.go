package Request

type ListAdministradoresTi struct {
}

type GetOneAdministradorTi struct {
}

type GetMyAdministradorTi struct {
}

type AddNewAdministradorTi struct {
	Id_rol                              *int    `json:"id_rol"`
	Id_grupo                            *int    `json:"id_grupo"`
	Rut_administrador_ti                *string `json:"rut_administrador_ti"`
	Nombres_administrador_ti            *string `json:"nombres_administrador_ti"`
	Apellidos_administrador_ti          *string `json:"apellidos_administrador_ti"`
	Hash_contrasena_administrador_ti    *string `json:"hash_contrasena_administrador_ti"`
	Correo_electronico_administrador_ti *string `json:"correo_electronico_administrador_ti"`
	Telefono_fijo_administrador_ti      *string `json:"telefono_fijo_administrador_ti"`
	Telefono_celular_administrador_ti   *string `json:"telefono_celular_administrador_ti"`
}

type PutOneAdministradorTi struct {
	Id_rol                              *int    `json:"id_rol"`
	Id_grupo                            *int    `json:"id_grupo"`
	Rut_administrador_ti                *string `json:"rut_administrador_ti"`
	Nombres_administrador_ti            *string `json:"nombres_administrador_ti"`
	Apellidos_administrador_ti          *string `json:"apellidos_administrador_ti"`
	Hash_contrasena_administrador_ti    *string `json:"hash_contrasena_administrador_ti"`
	Correo_electronico_administrador_ti *string `json:"correo_electronico_administrador_ti"`
	Telefono_fijo_administrador_ti      *string `json:"telefono_fijo_administrador_ti"`
	Telefono_celular_administrador_ti   *string `json:"telefono_celular_administrador_ti"`
}

type PutMyAdministradorTi struct {
	Id_rol                              *int    `json:"id_rol"`
	Id_grupo                            *int    `json:"id_grupo"`
	Rut_administrador_ti                *string `json:"rut_administrador_ti"`
	Nombres_administrador_ti            *string `json:"nombres_administrador_ti"`
	Apellidos_administrador_ti          *string `json:"apellidos_administrador_ti"`
	Hash_contrasena_administrador_ti    *string `json:"hash_contrasena_administrador_ti"`
	Correo_electronico_administrador_ti *string `json:"correo_electronico_administrador_ti"`
	Telefono_fijo_administrador_ti      *string `json:"telefono_fijo_administrador_ti"`
	Telefono_celular_administrador_ti   *string `json:"telefono_celular_administrador_ti"`
}

type DeleteAdministradorTi struct {
}
