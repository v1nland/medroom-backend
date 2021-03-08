package Swagger

import "medroom-backend/Messages/Response"

type ListPuntajesSwagger struct {
	Status bool                          `json:"status"`
	Meta   string                        `json:"meta"`
	Data   Response.ListPuntajesResponse `json:"data"`
}

type GetOnePuntajeSwagger struct {
	Status bool                           `json:"status"`
	Meta   string                         `json:"meta"`
	Data   Response.GetOnePuntajeResponse `json:"data"`
}

type AddNewPuntajeSwagger struct {
	Status bool                           `json:"status"`
	Meta   string                         `json:"meta"`
	Data   Response.AddNewPuntajeResponse `json:"data"`
}

type PutOnePuntajeSwagger struct {
	Status bool                           `json:"status"`
	Meta   string                         `json:"meta"`
	Data   Response.PutOnePuntajeResponse `json:"data"`
}

type DeletePuntajeSwagger struct {
	Status bool                           `json:"status"`
	Meta   string                         `json:"meta"`
	Data   Response.DeletePuntajeResponse `json:"data"`
}
