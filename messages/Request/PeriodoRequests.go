package Request

type ListPeriodos struct {
}

type GetOnePeriodo struct {
}

type AddNewPeriodo struct {
	Nombre_periodo *string `json:"nombre_periodo"`
}

type PutOnePeriodo struct {
	Nombre_periodo *string `json:"nombre_periodo"`
}

type DeletePeriodo struct {
}
