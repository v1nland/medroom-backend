package Response

type ListCursosResponse struct {
	Id            int                   `json:"id"`
	Periodo_curso GetOnePeriodoResponse `json:"periodo_curso"`
	Grupos_curso  []ListGruposResponse  `json:"grupos_curso"`
	Nombre_curso  string                `json:"nombre_curso"`
	Sigla_curso   string                `json:"sigla_curso"`
}

type GetOneCursoResponse struct {
	Id            int                   `json:"id"`
	Periodo_curso GetOnePeriodoResponse `json:"periodo_curso"`
	Grupos_curso  []ListGruposResponse  `json:"grupos_curso"`
	Nombre_curso  string                `json:"nombre_curso"`
	Sigla_curso   string                `json:"sigla_curso"`
}

type GetCursosEstudianteResponse struct {
	Id            int                   `json:"id"`
	Periodo_curso GetOnePeriodoResponse `json:"periodo_curso"`
	Grupos_curso  []ListGruposResponse  `json:"grupos_curso"`
	Nombre_curso  string                `json:"nombre_curso"`
	Sigla_curso   string                `json:"sigla_curso"`
}

type GetOneCursoEstudianteResponse struct {
	Id            int                   `json:"id"`
	Periodo_curso GetOnePeriodoResponse `json:"periodo_curso"`
	Grupos_curso  []ListGruposResponse  `json:"grupos_curso"`
	Nombre_curso  string                `json:"nombre_curso"`
	Sigla_curso   string                `json:"sigla_curso"`
}

type GetCursosEvaluadorResponse struct {
	Id            int                   `json:"id"`
	Periodo_curso GetOnePeriodoResponse `json:"periodo_curso"`
	Grupos_curso  []ListGruposResponse  `json:"grupos_curso"`
	Nombre_curso  string                `json:"nombre_curso"`
	Sigla_curso   string                `json:"sigla_curso"`
}

type GetOneCursoEvaluadorResponse struct {
	Id            int                   `json:"id"`
	Periodo_curso GetOnePeriodoResponse `json:"periodo_curso"`
	Grupos_curso  []ListGruposResponse  `json:"grupos_curso"`
	Nombre_curso  string                `json:"nombre_curso"`
	Sigla_curso   string                `json:"sigla_curso"`
}

type AddNewCursoResponse struct {
	Id int `json:"id"`
}

type PutOneCursoResponse struct {
	Id int `json:"id"`
}

type DeleteCursoResponse struct {
	Id int `json:"id"`
}

type AddEstudianteToCursoResponse struct {
	Id int `json:"id"`
}
