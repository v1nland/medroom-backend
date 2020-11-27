package SwaggerMessages

import (
	"medroom-backend/ResponseMessages"
)

type ListPeriodosSwagger struct {
	Status bool                                  `json:"Status"`
	Meta   string                                `json:"Meta"`
	Data   ResponseMessages.ListPeriodosResponse `json:"Data"`
}

type GetOnePeriodoSwagger struct {
	Status bool                                   `json:"Status"`
	Meta   string                                 `json:"Meta"`
	Data   ResponseMessages.GetOnePeriodoResponse `json:"Data"`
}

type AddNewPeriodoSwagger struct {
	Status bool                                   `json:"Status"`
	Meta   string                                 `json:"Meta"`
	Data   ResponseMessages.AddNewPeriodoResponse `json:"Data"`
}

type PutOnePeriodoSwagger struct {
	Status bool                                   `json:"Status"`
	Meta   string                                 `json:"Meta"`
	Data   ResponseMessages.PutOnePeriodoResponse `json:"Data"`
}

type DeletePeriodoSwagger struct {
	Status bool                                   `json:"Status"`
	Meta   string                                 `json:"Meta"`
	Data   ResponseMessages.DeletePeriodoResponse `json:"Data"`
}
