package Request

type ListEstudiantes struct {
}

type GetOneEstudiante struct {
}

type AddNewEstudiante struct {
	Id_rol                        *int    `json:"id_rol"`
	Rut_estudiante                *string `json:"rut_estudiante"`
	Nombres_estudiante            *string `json:"nombres_estudiante"`
	Apellidos_estudiante          *string `json:"apellidos_estudiante"`
	Hash_contrasena_estudiante    *string `json:"hash_contrasena_estudiante"`
	Correo_electronico_estudiante *string `json:"correo_electronico_estudiante"`
	Telefono_fijo_estudiante      *string `json:"telefono_fijo_estudiante"`
	Telefono_celular_estudiante   *string `json:"telefono_celular_estudiante"`
}

type PutOneEstudiante struct {
	Rut_estudiante                *string `json:"rut_estudiante"`
	Nombres_estudiante            *string `json:"nombres_estudiante"`
	Apellidos_estudiante          *string `json:"apellidos_estudiante"`
	Hash_contrasena_estudiante    *string `json:"hash_contrasena_estudiante"`
	Correo_electronico_estudiante *string `json:"correo_electronico_estudiante"`
	Telefono_fijo_estudiante      *string `json:"telefono_fijo_estudiante"`
	Telefono_celular_estudiante   *string `json:"telefono_celular_estudiante"`
}

type DeleteEstudiante struct {
}

type GetMyEstudiante struct {
}

type PutMyEstudiante struct {
	Hash_contrasena_estudiante       *string `json:"hash_contrasena_estudiante"`
	Hash_nueva_contrasena_estudiante *string `json:"hash_nueva_contrasena_estudiante"`
	Telefono_fijo_estudiante         *string `json:"telefono_fijo_estudiante"`
	Telefono_celular_estudiante      *string `json:"telefono_celular_estudiante"`
}

type AddEstudianteToGrupo struct {
	Id_grupos []int `json:"id_grupos"`
}

type AddEstudianteToCurso struct {
}

type ListEstudiantesCurso struct {
}

type ListEstudiantesCursoSinGrupo struct {
}
