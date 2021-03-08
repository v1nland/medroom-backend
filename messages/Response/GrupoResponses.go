package Response

type ListGruposResponse struct {
	Id                int                       `json:"id"`
	Id_curso          int                       `json:"id_curso"`
	Evaluador_grupo   GetOneEvaluadorResponse   `json:"evaluador_grupo"`
	Estudiantes_grupo []ListEstudiantesResponse `json:"estudiantes_grupo"`
	Nombre_grupo      string                    `json:"nombre_grupo"`
	Sigla_grupo       string                    `json:"sigla_grupo"`
}

type GetOneGrupoResponse struct {
	Id                int                       `json:"id"`
	Id_curso          int                       `json:"id_curso"`
	Evaluador_grupo   GetOneEvaluadorResponse   `json:"evaluador_grupo"`
	Estudiantes_grupo []ListEstudiantesResponse `json:"estudiantes_grupo"`
	Nombre_grupo      string                    `json:"nombre_grupo"`
	Sigla_grupo       string                    `json:"sigla_grupo"`
}

type GetOneGrupoEstudianteResponse struct {
	Id                int                       `json:"id"`
	Id_curso          int                       `json:"id_curso"`
	Evaluador_grupo   GetOneEvaluadorResponse   `json:"evaluador_grupo"`
	Estudiantes_grupo []ListEstudiantesResponse `json:"estudiantes_grupo"`
	Nombre_grupo      string                    `json:"nombre_grupo"`
	Sigla_grupo       string                    `json:"sigla_grupo"`
}

type GetGruposEstudianteResponse struct {
	Id                int                       `json:"id"`
	Id_curso          int                       `json:"id_curso"`
	Evaluador_grupo   GetOneEvaluadorResponse   `json:"evaluador_grupo"`
	Estudiantes_grupo []ListEstudiantesResponse `json:"estudiantes_grupo"`
	Nombre_grupo      string                    `json:"nombre_grupo"`
	Sigla_grupo       string                    `json:"sigla_grupo"`
}

type GetOneGrupoEvaluadorResponse struct {
	Id                int                       `json:"id"`
	Id_curso          int                       `json:"id_curso"`
	Evaluador_grupo   GetOneEvaluadorResponse   `json:"evaluador_grupo"`
	Estudiantes_grupo []ListEstudiantesResponse `json:"estudiantes_grupo"`
	Nombre_grupo      string                    `json:"nombre_grupo"`
	Sigla_grupo       string                    `json:"sigla_grupo"`
}

type GetGruposEvaluadorResponse struct {
	Id                int                       `json:"id"`
	Id_curso          int                       `json:"id_curso"`
	Evaluador_grupo   GetOneEvaluadorResponse   `json:"evaluador_grupo"`
	Estudiantes_grupo []ListEstudiantesResponse `json:"estudiantes_grupo"`
	Nombre_grupo      string                    `json:"nombre_grupo"`
	Sigla_grupo       string                    `json:"sigla_grupo"`
}

type AddNewGrupoResponse struct {
	Id int `json:"id"`
}

type PutOneGrupoResponse struct {
	Id int `json:"id"`
}

type DeleteGrupoResponse struct {
	Id int `json:"id"`
}

type AddEstudianteToGrupoResponse struct {
	Id int `json:"id"`
}

type AddEvaluadorToGrupoResponse struct {
	Id int `json:"id"`
}

type GetOneGrupoAdministradorAcademicoResponse struct {
	Id                int                       `json:"id"`
	Id_curso          int                       `json:"id_curso"`
	Evaluador_grupo   GetOneEvaluadorResponse   `json:"evaluador_grupo"`
	Estudiantes_grupo []ListEstudiantesResponse `json:"estudiantes_grupo"`
	Nombre_grupo      string                    `json:"nombre_grupo"`
	Sigla_grupo       string                    `json:"sigla_grupo"`
}

type GetGruposAdministradorAcademicoResponse struct {
	Id                int                       `json:"id"`
	Id_curso          int                       `json:"id_curso"`
	Evaluador_grupo   GetOneEvaluadorResponse   `json:"evaluador_grupo"`
	Estudiantes_grupo []ListEstudiantesResponse `json:"estudiantes_grupo"`
	Nombre_grupo      string                    `json:"nombre_grupo"`
	Sigla_grupo       string                    `json:"sigla_grupo"`
}
