package Swagger

import "medroom-backend/app/messages/Response"

type ListAdministradoresTiSwagger struct {
	Status bool                                   `json:"status"`
	Meta   string                                 `json:"meta"`
	Data   Response.ListAdministradoresTiResponse `json:"data"`
}

type GetOneAdministradorTiSwagger struct {
	Status bool                                   `json:"status"`
	Meta   string                                 `json:"meta"`
	Data   Response.GetOneAdministradorTiResponse `json:"data"`
}

type GetMyAdministradorTiSwagger struct {
	Status bool                                  `json:"status"`
	Meta   string                                `json:"meta"`
	Data   Response.GetMyAdministradorTiResponse `json:"data"`
}

type AddNewAdministradorTiSwagger struct {
	Status bool                                   `json:"status"`
	Meta   string                                 `json:"meta"`
	Data   Response.AddNewAdministradorTiResponse `json:"data"`
}

type PutOneAdministradorTiSwagger struct {
	Status bool                                   `json:"status"`
	Meta   string                                 `json:"meta"`
	Data   Response.PutOneAdministradorTiResponse `json:"data"`
}

type PutMyAdministradorTiSwagger struct {
	Status bool                                  `json:"status"`
	Meta   string                                `json:"meta"`
	Data   Response.PutMyAdministradorTiResponse `json:"data"`
}

type DeleteAdministradorTiSwagger struct {
	Status bool                                   `json:"status"`
	Meta   string                                 `json:"meta"`
	Data   Response.DeleteAdministradorTiResponse `json:"data"`
}
