package Request

type ListGruposPayload struct {
}

type GetOneGrupoPayload struct {
}

type GetGrupoEstudiantePayload struct {
}

type GetGrupoEvaluadorPayload struct {
}

type AddNewGrupoPayload struct {
	Id_curso     int    `json:"id_curso"`
	Id_evaluador string `json:"id_evaluador"`
	Nombre_grupo string `json:"nombre_grupo"`
	Sigla_grupo  string `json:"sigla_grupo"`
}

type PutOneGrupoPayload struct {
	Id_curso     int    `json:"id_curso"`
	Id_evaluador string `json:"id_evaluador"`
	Nombre_grupo string `json:"nombre_grupo"`
	Sigla_grupo  string `json:"sigla_grupo"`
}

type DeleteGrupoPayload struct {
}
