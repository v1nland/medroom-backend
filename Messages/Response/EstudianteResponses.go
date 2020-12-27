package Response

type ListEstudiantesResponse struct {
	Id                            string                               `json:"id"`
	Rol_estudiante                GetOneRolResponse                    `json:"rol_estudiante"`
	Evaluaciones_estudiante       []ListEvaluacionesEstudianteResponse `json:"evaluaciones_estudiante"`
	Id_grupo                      int                                  `json:"id_grupo"`
	Rut_estudiante                string                               `json:"rut_estudiante"`
	Nombres_estudiante            string                               `json:"nombres_estudiante"`
	Apellidos_estudiante          string                               `json:"apellidos_estudiante"`
	Correo_electronico_estudiante string                               `json:"correo_electronico_estudiante"`
	Telefono_fijo_estudiante      string                               `json:"telefono_fijo_estudiante"`
	Telefono_celular_estudiante   string                               `json:"telefono_celular_estudiante"`
}

type GetOneEstudianteResponse struct {
	Id                            string                               `json:"id"`
	Rol_estudiante                GetOneRolResponse                    `json:"rol_estudiante"`
	Evaluaciones_estudiante       []ListEvaluacionesEstudianteResponse `json:"evaluaciones_estudiante"`
	Id_grupo                      int                                  `json:"id_grupo"`
	Rut_estudiante                string                               `json:"rut_estudiante"`
	Nombres_estudiante            string                               `json:"nombres_estudiante"`
	Apellidos_estudiante          string                               `json:"apellidos_estudiante"`
	Correo_electronico_estudiante string                               `json:"correo_electronico_estudiante"`
	Telefono_fijo_estudiante      string                               `json:"telefono_fijo_estudiante"`
	Telefono_celular_estudiante   string                               `json:"telefono_celular_estudiante"`
}

type GetMyEstudianteResponse struct {
	Id                            string                               `json:"id"`
	Rol_estudiante                GetOneRolResponse                    `json:"rol_estudiante"`
	Evaluaciones_estudiante       []ListEvaluacionesEstudianteResponse `json:"evaluaciones_estudiante"`
	Id_grupo                      int                                  `json:"id_grupo"`
	Rut_estudiante                string                               `json:"rut_estudiante"`
	Nombres_estudiante            string                               `json:"nombres_estudiante"`
	Apellidos_estudiante          string                               `json:"apellidos_estudiante"`
	Correo_electronico_estudiante string                               `json:"correo_electronico_estudiante"`
	Telefono_fijo_estudiante      string                               `json:"telefono_fijo_estudiante"`
	Telefono_celular_estudiante   string                               `json:"telefono_celular_estudiante"`
}

type AddNewEstudianteResponse struct {
	Id string `json:"id"`
}

type PutOneEstudianteResponse struct {
	Id string `json:"id"`
}

type DeleteEstudianteResponse struct {
	Id string `json:"id"`
}

type PutMyEstudianteResponse struct {
	Id string `json:"id"`
}
