package ResponseMessages

type ListGruposResponse struct {
	Curso_grupo       GetOneCursoResponse       `json:"curso_grupo"`
	Evaluador_grupo   GetOneEvaluadorResponse   `json:"evaluador_grupo"`
	Estudiantes_grupo []ListEstudiantesResponse `json:"estudiantes_grupo"`
	Nombre_grupo      string                    `json:"nombre_grupo"`
	Sigla_grupo       string                    `json:"sigla_grupo"`
}

type GetOneGrupoResponse struct {
	Curso_grupo       GetOneCursoResponse       `json:"curso_grupo"`
	Evaluador_grupo   GetOneEvaluadorResponse   `json:"evaluador_grupo"`
	Estudiantes_grupo []ListEstudiantesResponse `json:"estudiantes_grupo"`
	Nombre_grupo      string                    `json:"nombre_grupo"`
	Sigla_grupo       string                    `json:"sigla_grupo"`
}

type GetGrupoEstudianteResponse struct {
	Curso_grupo       GetOneCursoResponse       `json:"curso_grupo"`
	Evaluador_grupo   GetOneEvaluadorResponse   `json:"evaluador_grupo"`
	Estudiantes_grupo []ListEstudiantesResponse `json:"estudiantes_grupo"`
	Nombre_grupo      string                    `json:"nombre_grupo"`
	Sigla_grupo       string                    `json:"sigla_grupo"`
}

type AddNewGrupoResponse struct {
	Id_curso     int    `json:"id_curso"`
	Id_evaluador string `json:"id_evaluador"`
	Nombre_grupo string `json:"nombre_grupo"`
	Sigla_grupo  string `json:"sigla_grupo"`
}

type PutOneGrupoResponse struct {
	Id_curso     int    `json:"id_curso"`
	Id_evaluador string `json:"id_evaluador"`
	Nombre_grupo string `json:"nombre_grupo"`
	Sigla_grupo  string `json:"sigla_grupo"`
}

type DeleteGrupoResponse struct {
	Id_curso     int    `json:"id_curso"`
	Id_evaluador string `json:"id_evaluador"`
	Nombre_grupo string `json:"nombre_grupo"`
	Sigla_grupo  string `json:"sigla_grupo"`
}
