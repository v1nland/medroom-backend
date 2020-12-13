package SwaggerMessages

import (
	"medroom-backend/ResponseMessages"
)

type ListPuntajesSwagger struct {
	Status bool                                  `json:"status"`
	Meta   string                                `json:"meta"`
	Data   ResponseMessages.ListPuntajesResponse `json:"data"`
}

type GetOnePuntajeSwagger struct {
	Status bool                                   `json:"status"`
	Meta   string                                 `json:"meta"`
	Data   ResponseMessages.GetOnePuntajeResponse `json:"data"`
}

type AddNewPuntajeSwagger struct {
	Status bool                                   `json:"status"`
	Meta   string                                 `json:"meta"`
	Data   ResponseMessages.AddNewPuntajeResponse `json:"data"`
}

type PutOnePuntajeSwagger struct {
	Status bool                                   `json:"status"`
	Meta   string                                 `json:"meta"`
	Data   ResponseMessages.PutOnePuntajeResponse `json:"data"`
}

type DeletePuntajeSwagger struct {
	Status bool                                   `json:"status"`
	Meta   string                                 `json:"meta"`
	Data   ResponseMessages.DeletePuntajeResponse `json:"data"`
}
