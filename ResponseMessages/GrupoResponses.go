package ResponseMessages

type ListGruposResponse struct {
	Nombre_grupo string `json:"nombre_grupo"`
	Sigla_grupo  string `json:"sigla_grupo"`
}

type GetOneGrupoResponse struct {
	Nombre_grupo string `json:"nombre_grupo"`
	Sigla_grupo  string `json:"sigla_grupo"`
}

type AddNewGrupoResponse struct {
	Nombre_grupo string `json:"nombre_grupo"`
	Sigla_grupo  string `json:"sigla_grupo"`
}

type PutOneGrupoResponse struct {
	Nombre_grupo string `json:"nombre_grupo"`
	Sigla_grupo  string `json:"sigla_grupo"`
}

type DeleteGrupoResponse struct {
	Nombre_grupo string `json:"nombre_grupo"`
	Sigla_grupo  string `json:"sigla_grupo"`
}
