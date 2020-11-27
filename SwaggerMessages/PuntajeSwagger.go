package SwaggerMessages

import (
	"medroom-backend/ResponseMessages"
)

type ListPuntajesSwagger struct {
	Status bool                                  `json:"Status"`
	Meta   string                                `json:"Meta"`
	Data   ResponseMessages.ListPuntajesResponse `json:"Data"`
}

type GetOnePuntajeSwagger struct {
	Status bool                                   `json:"Status"`
	Meta   string                                 `json:"Meta"`
	Data   ResponseMessages.GetOnePuntajeResponse `json:"Data"`
}

type AddNewPuntajeSwagger struct {
	Status bool                                   `json:"Status"`
	Meta   string                                 `json:"Meta"`
	Data   ResponseMessages.AddNewPuntajeResponse `json:"Data"`
}

type PutOnePuntajeSwagger struct {
	Status bool                                   `json:"Status"`
	Meta   string                                 `json:"Meta"`
	Data   ResponseMessages.PutOnePuntajeResponse `json:"Data"`
}

type DeletePuntajeSwagger struct {
	Status bool                                   `json:"Status"`
	Meta   string                                 `json:"Meta"`
	Data   ResponseMessages.DeletePuntajeResponse `json:"Data"`
}
