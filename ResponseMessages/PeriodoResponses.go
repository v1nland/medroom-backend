package ResponseMessages

type ListPeriodosResponse struct {
	ID             int    `json:"id"`
	Nombre_periodo string `json:"nombre_periodo"`
}

type GetOnePeriodoResponse struct {
	ID             int    `json:"id"`
	Nombre_periodo string `json:"nombre_periodo"`
}

type AddNewPeriodoResponse struct {
	ID             int    `json:"id"`
	Nombre_periodo string `json:"nombre_periodo"`
}

type PutOnePeriodoResponse struct {
	ID             int    `json:"id"`
	Nombre_periodo string `json:"nombre_periodo"`
}

type DeletePeriodoResponse struct {
	ID             int    `json:"id"`
	Nombre_periodo string `json:"nombre_periodo"`
}
