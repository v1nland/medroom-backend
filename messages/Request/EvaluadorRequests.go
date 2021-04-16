package Request

type ListEvaluadores struct {
}

type GetOneEvaluador struct {
}

type AddNewEvaluador struct {
	Id_rol                       *int    `json:"id_rol"`
	Rut_evaluador                *string `json:"rut_evaluador"`
	Nombres_evaluador            *string `json:"nombres_evaluador"`
	Apellidos_evaluador          *string `json:"apellidos_evaluador"`
	Hash_contrasena_evaluador    *string `json:"hash_contrasena_evaluador"`
	Correo_electronico_evaluador *string `json:"correo_electronico_evaluador"`
	Telefono_fijo_evaluador      *string `json:"telefono_fijo_evaluador"`
	Telefono_celular_evaluador   *string `json:"telefono_celular_evaluador"`
	Recinto_evaluador            *string `json:"recinto_evaluador"`
	Cargo_evaluador              *string `json:"cargo_evaluador"`
}

type PutOneEvaluador struct {
	Rut_evaluador                *string `json:"rut_evaluador"`
	Nombres_evaluador            *string `json:"nombres_evaluador"`
	Apellidos_evaluador          *string `json:"apellidos_evaluador"`
	Hash_contrasena_evaluador    *string `json:"hash_contrasena_evaluador"`
	Correo_electronico_evaluador *string `json:"correo_electronico_evaluador"`
	Telefono_fijo_evaluador      *string `json:"telefono_fijo_evaluador"`
	Telefono_celular_evaluador   *string `json:"telefono_celular_evaluador"`
	Recinto_evaluador            *string `json:"recinto_evaluador"`
	Cargo_evaluador              *string `json:"cargo_evaluador"`
}

type DeleteEvaluador struct {
}

type GetMyEvaluador struct {
}

type PutMyEvaluador struct {
	Hash_contrasena_evaluador       *string `json:"hash_contrasena_evaluador"`
	Hash_nueva_contrasena_evaluador *string `json:"hash_nueva_contrasena_evaluador"`
	Telefono_fijo_evaluador         *string `json:"telefono_fijo_evaluador"`
	Telefono_celular_evaluador      *string `json:"telefono_celular_evaluador"`
	Cargo_evaluador                 *string `json:"cargo_evaluador"`
}

type AddEvaluadorToGrupo struct {
	Id_grupos []int `json:"id_grupos"`
}
