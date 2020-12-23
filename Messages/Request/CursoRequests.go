package Request

type ListCursosPayload struct {
}

type GetOneCursoPayload struct {
}

type GetCursoEstudiantePayload struct {
}

type GetCursoEvaluadorPayload struct {
}

type AddNewCursoPayload struct {
	Id_periodo   *int    `json:"id_periodo"`
	Nombre_curso *string `json:"nombre_curso"`
	Sigla_curso  *string `json:"sigla_curso"`
}

type PutOneCursoPayload struct {
	Id_periodo   *int    `json:"id_periodo"`
	Nombre_curso *string `json:"nombre_curso"`
	Sigla_curso  *string `json:"sigla_curso"`
}

type DeleteCursoPayload struct {
}
