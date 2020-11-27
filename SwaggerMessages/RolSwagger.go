package SwaggerMessages

import (
	"medroom-backend/ResponseMessages"
)

type ListRolesSwagger struct {
	Status bool                               `json:"Status"`
	Meta   string                             `json:"Meta"`
	Data   ResponseMessages.ListRolesResponse `json:"Data"`
}

type GetOneRolSwagger struct {
	Status bool                               `json:"Status"`
	Meta   string                             `json:"Meta"`
	Data   ResponseMessages.GetOneRolResponse `json:"Data"`
}

type AddNewRolSwagger struct {
	Status bool                               `json:"Status"`
	Meta   string                             `json:"Meta"`
	Data   ResponseMessages.AddNewRolResponse `json:"Data"`
}

type PutOneRolSwagger struct {
	Status bool                               `json:"Status"`
	Meta   string                             `json:"Meta"`
	Data   ResponseMessages.PutOneRolResponse `json:"Data"`
}

type DeleteRolSwagger struct {
	Status bool                               `json:"Status"`
	Meta   string                             `json:"Meta"`
	Data   ResponseMessages.DeleteRolResponse `json:"Data"`
}
