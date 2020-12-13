package SwaggerMessages

import (
	"medroom-backend/ResponseMessages"
)

type ListRolesSwagger struct {
	Status bool                               `json:"status"`
	Meta   string                             `json:"meta"`
	Data   ResponseMessages.ListRolesResponse `json:"data"`
}

type GetOneRolSwagger struct {
	Status bool                               `json:"status"`
	Meta   string                             `json:"meta"`
	Data   ResponseMessages.GetOneRolResponse `json:"data"`
}

type AddNewRolSwagger struct {
	Status bool                               `json:"status"`
	Meta   string                             `json:"meta"`
	Data   ResponseMessages.AddNewRolResponse `json:"data"`
}

type PutOneRolSwagger struct {
	Status bool                               `json:"status"`
	Meta   string                             `json:"meta"`
	Data   ResponseMessages.PutOneRolResponse `json:"data"`
}

type DeleteRolSwagger struct {
	Status bool                               `json:"status"`
	Meta   string                             `json:"meta"`
	Data   ResponseMessages.DeleteRolResponse `json:"data"`
}
