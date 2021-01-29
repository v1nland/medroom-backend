package Request

type ListGrupos struct {
}

type GetOneGrupo struct {
}

type GetGrupoEstudiante struct {
}

type GetGrupoEvaluador struct {
}

type AddNewGrupo struct {
	Id_curso     *int    `json:"id_curso"`
	Id_evaluador *string `json:"id_evaluador"`
	Nombre_grupo *string `json:"nombre_grupo"`
	Sigla_grupo  *string `json:"sigla_grupo"`
}

type PutOneGrupo struct {
	Id_curso     *int    `json:"id_curso"`
	Id_evaluador *string `json:"id_evaluador"`
	Nombre_grupo *string `json:"nombre_grupo"`
	Sigla_grupo  *string `json:"sigla_grupo"`
}

type DeleteGrupo struct {
}
