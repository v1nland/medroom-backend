package SwaggerMessages

import (
	"medroom-backend/ResponseMessages"
)

type ListPeriodosSwagger struct {
	Status bool                                  `json:"status"`
	Meta   string                                `json:"meta"`
	Data   ResponseMessages.ListPeriodosResponse `json:"data"`
}

type GetOnePeriodoSwagger struct {
	Status bool                                   `json:"status"`
	Meta   string                                 `json:"meta"`
	Data   ResponseMessages.GetOnePeriodoResponse `json:"data"`
}

type AddNewPeriodoSwagger struct {
	Status bool                                   `json:"status"`
	Meta   string                                 `json:"meta"`
	Data   ResponseMessages.AddNewPeriodoResponse `json:"data"`
}

type PutOnePeriodoSwagger struct {
	Status bool                                   `json:"status"`
	Meta   string                                 `json:"meta"`
	Data   ResponseMessages.PutOnePeriodoResponse `json:"data"`
}

type DeletePeriodoSwagger struct {
	Status bool                                   `json:"status"`
	Meta   string                                 `json:"meta"`
	Data   ResponseMessages.DeletePeriodoResponse `json:"data"`
}
