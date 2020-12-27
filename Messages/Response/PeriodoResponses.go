package Response

type ListPeriodosResponse struct {
	Id             int    `json:"id"`
	Nombre_periodo string `json:"nombre_periodo"`
}

type GetOnePeriodoResponse struct {
	Id             int    `json:"id"`
	Nombre_periodo string `json:"nombre_periodo"`
}

type AddNewPeriodoResponse struct {
	Id int `json:"id"`
}

type PutOnePeriodoResponse struct {
	Id int `json:"id"`
}

type DeletePeriodoResponse struct {
	Id int `json:"id"`
}
