package SwaggerMessages

import (
	"medroom-backend/ResponseMessages"
)

type ListAdministradoresTiSwagger struct {
	Status bool                                           `json:"status"`
	Meta   string                                         `json:"meta"`
	Data   ResponseMessages.ListAdministradoresTiResponse `json:"data"`
}

type GetOneAdministradorTiSwagger struct {
	Status bool                                           `json:"status"`
	Meta   string                                         `json:"meta"`
	Data   ResponseMessages.GetOneAdministradorTiResponse `json:"data"`
}

type AddNewAdministradorTiSwagger struct {
	Status bool                                           `json:"status"`
	Meta   string                                         `json:"meta"`
	Data   ResponseMessages.AddNewAdministradorTiResponse `json:"data"`
}

type PutOneAdministradorTiSwagger struct {
	Status bool                                           `json:"status"`
	Meta   string                                         `json:"meta"`
	Data   ResponseMessages.PutOneAdministradorTiResponse `json:"data"`
}

type DeleteAdministradorTiSwagger struct {
	Status bool                                           `json:"status"`
	Meta   string                                         `json:"meta"`
	Data   ResponseMessages.DeleteAdministradorTiResponse `json:"data"`
}
