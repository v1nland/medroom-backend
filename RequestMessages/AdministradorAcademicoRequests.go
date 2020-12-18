package RequestMessages

type ListAdministradoresAcademicosPayload struct {
}

type GetOneAdministradorAcademicoPayload struct {
}

type GetMyAdministradorAcademicoPayload struct {
}

type AddNewAdministradorAcademicoPayload struct {
	Id_rol                                     int    `json:"id_rol"`
	Rut_administrador_academico                string `json:"rut_administrador_academico"`
	Nombres_administrador_academico            string `json:"nombres_administrador_academico"`
	Apellidos_administrador_academico          string `json:"apellidos_administrador_academico"`
	Hash_contrasena_administrador_academico    string `json:"hash_contrasena_administrador_academico"`
	Correo_electronico_administrador_academico string `json:"correo_electronico_administrador_academico"`
	Telefono_fijo_administrador_academico      string `json:"telefono_fijo_administrador_academico"`
	Telefono_celular_administrador_academico   string `json:"telefono_celular_administrador_academico"`
}

type PutOneAdministradorAcademicoPayload struct {
	Id_rol                                     int    `json:"id_rol"`
	Rut_administrador_academico                string `json:"rut_administrador_academico"`
	Nombres_administrador_academico            string `json:"nombres_administrador_academico"`
	Apellidos_administrador_academico          string `json:"apellidos_administrador_academico"`
	Hash_contrasena_administrador_academico    string `json:"hash_contrasena_administrador_academico"`
	Correo_electronico_administrador_academico string `json:"correo_electronico_administrador_academico"`
	Telefono_fijo_administrador_academico      string `json:"telefono_fijo_administrador_academico"`
	Telefono_celular_administrador_academico   string `json:"telefono_celular_administrador_academico"`
}

type PutMyAdministradorAcademicoPayload struct {
	Id_rol                                     int    `json:"id_rol"`
	Rut_administrador_academico                string `json:"rut_administrador_academico"`
	Nombres_administrador_academico            string `json:"nombres_administrador_academico"`
	Apellidos_administrador_academico          string `json:"apellidos_administrador_academico"`
	Hash_contrasena_administrador_academico    string `json:"hash_contrasena_administrador_academico"`
	Correo_electronico_administrador_academico string `json:"correo_electronico_administrador_academico"`
	Telefono_fijo_administrador_academico      string `json:"telefono_fijo_administrador_academico"`
	Telefono_celular_administrador_academico   string `json:"telefono_celular_administrador_academico"`
}

type DeleteAdministradorAcademicoPayload struct {
}
