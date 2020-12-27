package Response

type ListAdministradoresAcademicosResponse struct {
	Id                                         string            `json:"id"`
	Rol_administrador_academico                GetOneRolResponse `json:"rol_administrador_academico"`
	Rut_administrador_academico                string            `json:"rut_administrador_academico"`
	Nombres_administrador_academico            string            `json:"nombres_administrador_academico"`
	Apellidos_administrador_academico          string            `json:"apellidos_administrador_academico"`
	Correo_electronico_administrador_academico string            `json:"correo_electronico_administrador_academico"`
	Telefono_fijo_administrador_academico      string            `json:"telefono_fijo_administrador_academico"`
	Telefono_celular_administrador_academico   string            `json:"telefono_celular_administrador_academico"`
}

type GetOneAdministradorAcademicoResponse struct {
	Id                                         string            `json:"id"`
	Rol_administrador_academico                GetOneRolResponse `json:"rol_administrador_academico"`
	Rut_administrador_academico                string            `json:"rut_administrador_academico"`
	Nombres_administrador_academico            string            `json:"nombres_administrador_academico"`
	Apellidos_administrador_academico          string            `json:"apellidos_administrador_academico"`
	Correo_electronico_administrador_academico string            `json:"correo_electronico_administrador_academico"`
	Telefono_fijo_administrador_academico      string            `json:"telefono_fijo_administrador_academico"`
	Telefono_celular_administrador_academico   string            `json:"telefono_celular_administrador_academico"`
}

type GetMyAdministradorAcademicoResponse struct {
	Id                                         string            `json:"id"`
	Rol_administrador_academico                GetOneRolResponse `json:"rol_administrador_academico"`
	Rut_administrador_academico                string            `json:"rut_administrador_academico"`
	Nombres_administrador_academico            string            `json:"nombres_administrador_academico"`
	Apellidos_administrador_academico          string            `json:"apellidos_administrador_academico"`
	Correo_electronico_administrador_academico string            `json:"correo_electronico_administrador_academico"`
	Telefono_fijo_administrador_academico      string            `json:"telefono_fijo_administrador_academico"`
	Telefono_celular_administrador_academico   string            `json:"telefono_celular_administrador_academico"`
}

type AddNewAdministradorAcademicoResponse struct {
	Id string `json:"id"`
}

type PutOneAdministradorAcademicoResponse struct {
	Id string `json:"id"`
}

type PutMyAdministradorAcademicoResponse struct {
	Id string `json:"id"`
}

type DeleteAdministradorAcademicoResponse struct {
	Id string `json:"id"`
}
