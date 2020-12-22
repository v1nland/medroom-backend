package Response

type ListCursosResponse struct {
	Periodo_curso GetOnePeriodoResponse `json:"periodo_curso"`
	Grupos_curso  []ListGruposResponse  `json:"grupos_curso"`
	Nombre_curso  string                `json:"nombre_curso"`
	Sigla_curso   string                `json:"sigla_curso"`
}

type GetOneCursoResponse struct {
	Periodo_curso GetOnePeriodoResponse `json:"periodo_curso"`
	Grupos_curso  []ListGruposResponse  `json:"grupos_curso"`
	Nombre_curso  string                `json:"nombre_curso"`
	Sigla_curso   string                `json:"sigla_curso"`
}

type GetCursoEstudianteResponse struct {
	Periodo_curso GetOnePeriodoResponse `json:"periodo_curso"`
	Grupos_curso  []ListGruposResponse  `json:"grupos_curso"`
	Nombre_curso  string                `json:"nombre_curso"`
	Sigla_curso   string                `json:"sigla_curso"`
}

type GetCursoEvaluadorResponse struct {
	Periodo_curso GetOnePeriodoResponse `json:"periodo_curso"`
	Grupos_curso  []ListGruposResponse  `json:"grupos_curso"`
	Nombre_curso  string                `json:"nombre_curso"`
	Sigla_curso   string                `json:"sigla_curso"`
}

type AddNewCursoResponse struct {
	Id_periodo   int    `json:"id_periodo"`
	Nombre_curso string `json:"nombre_curso"`
	Sigla_curso  string `json:"sigla_curso"`
}

type PutOneCursoResponse struct {
	Id_periodo   int    `json:"id_periodo"`
	Nombre_curso string `json:"nombre_curso"`
	Sigla_curso  string `json:"sigla_curso"`
}

type DeleteCursoResponse struct {
	Id_periodo   int    `json:"id_periodo"`
	Nombre_curso string `json:"nombre_curso"`
	Sigla_curso  string `json:"sigla_curso"`
}
