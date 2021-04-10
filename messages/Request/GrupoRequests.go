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
	Id_curso     *string `json:"id_curso"`
	Nombre_grupo *string `json:"nombre_grupo"`
	Sigla_grupo  *string `json:"sigla_grupo"`
}

type PutOneGrupo struct {
	Id_curso     *string `json:"id_curso"`
	Nombre_grupo *string `json:"nombre_grupo"`
	Sigla_grupo  *string `json:"sigla_grupo"`
}

type DeleteGrupo struct {
}
