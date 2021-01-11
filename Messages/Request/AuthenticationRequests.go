package Request

type LoginEstudiantePayload struct {
	Correo_electronico_estudiante string `json:"correo_electronico_estudiante"`
	Hash_contrasena_estudiante    string `json:"hash_contrasena_estudiante"`
}

type LoginEvaluadorPayload struct {
	Correo_electronico_evaluador string `json:"correo_electronico_evaluador"`
	Hash_contrasena_evaluador    string `json:"hash_contrasena_evaluador"`
}

type LoginAdministradorAcademicoPayload struct {
	Correo_electronico_administrador_academico string `json:"correo_electronico_administrador_academico"`
	Hash_contrasena_administrador_academico    string `json:"hash_contrasena_administrador_academico"`
}

type LoginAdministradorTiPayload struct {
	Correo_electronico_administrador_ti string `json:"correo_electronico_administrador_ti"`
	Hash_contrasena_administrador_ti    string `json:"hash_contrasena_administrador_ti"`
}
