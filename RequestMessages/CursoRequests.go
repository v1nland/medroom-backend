package RequestMessages

type ListCursosPayload struct {
}

type GetOneCursoPayload struct {
}

type AddNewCursoPayload struct {
	Id_periodo   int    `json:"id_periodo"`
	Nombre_curso string `json:"nombre_curso"`
	Sigla_curso  string `json:"sigla_curso"`
}

type PutOneCursoPayload struct {
	Id_periodo   int    `json:"id_periodo"`
	Nombre_curso string `json:"nombre_curso"`
	Sigla_curso  string `json:"sigla_curso"`
}

type DeleteCursoPayload struct {
}
