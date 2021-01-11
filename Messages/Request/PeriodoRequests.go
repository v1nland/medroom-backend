package Request

type ListPeriodosPayload struct {
}

type GetOnePeriodoPayload struct {
}

type AddNewPeriodoPayload struct {
	Nombre_periodo string `json:"nombre_periodo"`
}

type PutOnePeriodoPayload struct {
	Nombre_periodo string `json:"nombre_periodo"`
}

type DeletePeriodoPayload struct {
}
