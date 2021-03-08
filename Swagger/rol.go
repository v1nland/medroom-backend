package Swagger

import "medroom-backend/Messages/Response"

type ListRolesSwagger struct {
	Status bool                       `json:"status"`
	Meta   string                     `json:"meta"`
	Data   Response.ListRolesResponse `json:"data"`
}

type GetOneRolSwagger struct {
	Status bool                       `json:"status"`
	Meta   string                     `json:"meta"`
	Data   Response.GetOneRolResponse `json:"data"`
}

type AddNewRolSwagger struct {
	Status bool                       `json:"status"`
	Meta   string                     `json:"meta"`
	Data   Response.AddNewRolResponse `json:"data"`
}

type PutOneRolSwagger struct {
	Status bool                       `json:"status"`
	Meta   string                     `json:"meta"`
	Data   Response.PutOneRolResponse `json:"data"`
}

type DeleteRolSwagger struct {
	Status bool                       `json:"status"`
	Meta   string                     `json:"meta"`
	Data   Response.DeleteRolResponse `json:"data"`
}
