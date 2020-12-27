package Request

type ListEvaluacionesGrupoEstudiantePayload struct {
}
type ListEvaluacionesGrupoEvaluadorPayload struct {
}
type ListEvaluacionesPayload struct {
}
type ListEvaluacionesEstudiantePayload struct {
}
type GetOneEvaluacionPayload struct {
}

type AddNewEvaluacionPayload struct {
	Nombre_evaluacion *string `json:"nombre_evaluacion"`
}

type PutOneEvaluacionPayload struct {
	Nombre_evaluacion *string `json:"nombre_evaluacion"`
}

type DeleteEvaluacionPayload struct {
}
