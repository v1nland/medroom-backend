package Swagger

import "medroom-backend/messages/Response"

type ListPeriodosSwagger struct {
	Status bool                          `json:"status"`
	Meta   string                        `json:"meta"`
	Data   Response.ListPeriodosResponse `json:"data"`
}

type GetOnePeriodoSwagger struct {
	Status bool                           `json:"status"`
	Meta   string                         `json:"meta"`
	Data   Response.GetOnePeriodoResponse `json:"data"`
}

type AddNewPeriodoSwagger struct {
	Status bool                           `json:"status"`
	Meta   string                         `json:"meta"`
	Data   Response.AddNewPeriodoResponse `json:"data"`
}

type PutOnePeriodoSwagger struct {
	Status bool                           `json:"status"`
	Meta   string                         `json:"meta"`
	Data   Response.PutOnePeriodoResponse `json:"data"`
}

type DeletePeriodoSwagger struct {
	Status bool                           `json:"status"`
	Meta   string                         `json:"meta"`
	Data   Response.DeletePeriodoResponse `json:"data"`
}
