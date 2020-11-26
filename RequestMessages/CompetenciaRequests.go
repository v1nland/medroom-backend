package RequestMessages

type ListCompetenciasPayload struct {
}

type GetOneCompetenciaPayload struct {
}

type AddNewCompetenciaPayload struct {
	Nombre_competencia string `json:"nombre_competencia"`
}

type PutOneCompetenciaPayload struct {
	Nombre_competencia string `json:"nombre_competencia"`
}

type DeleteCompetenciaPayload struct {
}
