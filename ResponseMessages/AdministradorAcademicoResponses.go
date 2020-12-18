package ResponseMessages

type ListAdministradoresAcademicosResponse struct {
	Rol_administrador_academico                GetOneRolResponse `json:"rol_administrador_academico"`
	Rut_administrador_academico                string            `json:"rut_administrador_academico"`
	Nombres_administrador_academico            string            `json:"nombres_administrador_academico"`
	Apellidos_administrador_academico          string            `json:"apellidos_administrador_academico"`
	Correo_electronico_administrador_academico string            `json:"correo_electronico_administrador_academico"`
	Telefono_fijo_administrador_academico      string            `json:"telefono_fijo_administrador_academico"`
	Telefono_celular_administrador_academico   string            `json:"telefono_celular_administrador_academico"`
}

type GetOneAdministradorAcademicoResponse struct {
	Rol_administrador_academico                GetOneRolResponse `json:"rol_administrador_academico"`
	Rut_administrador_academico                string            `json:"rut_administrador_academico"`
	Nombres_administrador_academico            string            `json:"nombres_administrador_academico"`
	Apellidos_administrador_academico          string            `json:"apellidos_administrador_academico"`
	Correo_electronico_administrador_academico string            `json:"correo_electronico_administrador_academico"`
	Telefono_fijo_administrador_academico      string            `json:"telefono_fijo_administrador_academico"`
	Telefono_celular_administrador_academico   string            `json:"telefono_celular_administrador_academico"`
}

type GetMyAdministradorAcademicoResponse struct {
	Rol_administrador_academico                GetOneRolResponse `json:"rol_administrador_academico"`
	Rut_administrador_academico                string            `json:"rut_administrador_academico"`
	Nombres_administrador_academico            string            `json:"nombres_administrador_academico"`
	Apellidos_administrador_academico          string            `json:"apellidos_administrador_academico"`
	Correo_electronico_administrador_academico string            `json:"correo_electronico_administrador_academico"`
	Telefono_fijo_administrador_academico      string            `json:"telefono_fijo_administrador_academico"`
	Telefono_celular_administrador_academico   string            `json:"telefono_celular_administrador_academico"`
}

type AddNewAdministradorAcademicoResponse struct {
	Id_rol                                     int    `json:"id_rol"`
	Rut_administrador_academico                string `json:"rut_administrador_academico"`
	Nombres_administrador_academico            string `json:"nombres_administrador_academico"`
	Apellidos_administrador_academico          string `json:"apellidos_administrador_academico"`
	Correo_electronico_administrador_academico string `json:"correo_electronico_administrador_academico"`
	Telefono_fijo_administrador_academico      string `json:"telefono_fijo_administrador_academico"`
	Telefono_celular_administrador_academico   string `json:"telefono_celular_administrador_academico"`
}

type PutOneAdministradorAcademicoResponse struct {
	Id_rol                                     int    `json:"id_rol"`
	Rut_administrador_academico                string `json:"rut_administrador_academico"`
	Nombres_administrador_academico            string `json:"nombres_administrador_academico"`
	Apellidos_administrador_academico          string `json:"apellidos_administrador_academico"`
	Hash_contrasena_administrador_academico    string `json:"hash_contrasena_administrador_academico"`
	Correo_electronico_administrador_academico string `json:"correo_electronico_administrador_academico"`
	Telefono_fijo_administrador_academico      string `json:"telefono_fijo_administrador_academico"`
	Telefono_celular_administrador_academico   string `json:"telefono_celular_administrador_academico"`
}

type PutMyAdministradorAcademicoResponse struct {
	Id_rol                                     int    `json:"id_rol"`
	Rut_administrador_academico                string `json:"rut_administrador_academico"`
	Nombres_administrador_academico            string `json:"nombres_administrador_academico"`
	Apellidos_administrador_academico          string `json:"apellidos_administrador_academico"`
	Hash_contrasena_administrador_academico    string `json:"hash_contrasena_administrador_academico"`
	Correo_electronico_administrador_academico string `json:"correo_electronico_administrador_academico"`
	Telefono_fijo_administrador_academico      string `json:"telefono_fijo_administrador_academico"`
	Telefono_celular_administrador_academico   string `json:"telefono_celular_administrador_academico"`
}

type DeleteAdministradorAcademicoResponse struct {
	Id_rol                                     int    `json:"id_rol"`
	Rut_administrador_academico                string `json:"rut_administrador_academico"`
	Nombres_administrador_academico            string `json:"nombres_administrador_academico"`
	Apellidos_administrador_academico          string `json:"apellidos_administrador_academico"`
	Correo_electronico_administrador_academico string `json:"correo_electronico_administrador_academico"`
	Telefono_fijo_administrador_academico      string `json:"telefono_fijo_administrador_academico"`
	Telefono_celular_administrador_academico   string `json:"telefono_celular_administrador_academico"`
}
