package RequestMessages

type ListGruposPayload struct {
}

type GetOneGrupoPayload struct {
}

type AddNewGrupoPayload struct {
	Nombre_grupo string `json:"nombre_grupo"`
	Sigla_grupo  string `json:"sigla_grupo"`
}

type PutOneGrupoPayload struct {
	Nombre_grupo string `json:"nombre_grupo"`
	Sigla_grupo  string `json:"sigla_grupo"`
}

type DeleteGrupoPayload struct {
}
