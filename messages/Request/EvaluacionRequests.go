package Request

type ListEvaluacionesGrupoEstudiante struct {
}
type ListEvaluacionesGrupoEvaluador struct {
}
type ListEvaluaciones struct {
}
type ListEvaluacionesEstudiante struct {
}
type GetOneEvaluacion struct {
}

type AddNewEvaluacion struct {
	Nombre_evaluacion *string `json:"nombre_evaluacion"`
}

type PutOneEvaluacion struct {
	Nombre_evaluacion *string `json:"nombre_evaluacion"`
}

type DeleteEvaluacion struct {
}
