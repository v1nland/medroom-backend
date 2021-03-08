package Request

type ListCursos struct {
}

type GetOneCurso struct {
}

type GetCursosEstudiante struct {
}

type GetOneCursoEstudiante struct {
}

type GetCursosEvaluador struct {
}

type GetOneCursoEvaluador struct {
}

type AddNewCurso struct {
	Id_periodo   *int    `json:"id_periodo"`
	Nombre_curso *string `json:"nombre_curso"`
	Sigla_curso  *string `json:"sigla_curso"`
}

type PutOneCurso struct {
	Id_periodo   *int    `json:"id_periodo"`
	Nombre_curso *string `json:"nombre_curso"`
	Sigla_curso  *string `json:"sigla_curso"`
}

type DeleteCurso struct {
}
