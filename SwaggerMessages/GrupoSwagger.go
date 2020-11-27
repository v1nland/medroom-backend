package SwaggerMessages

import (
	"medroom-backend/ResponseMessages"
)

type ListGruposSwagger struct {
	Status bool                                `json:"Status"`
	Meta   string                              `json:"Meta"`
	Data   ResponseMessages.ListGruposResponse `json:"Data"`
}

type GetOneGrupoSwagger struct {
	Status bool                                 `json:"Status"`
	Meta   string                               `json:"Meta"`
	Data   ResponseMessages.GetOneGrupoResponse `json:"Data"`
}

type AddNewGrupoSwagger struct {
	Status bool                                 `json:"Status"`
	Meta   string                               `json:"Meta"`
	Data   ResponseMessages.AddNewGrupoResponse `json:"Data"`
}

type PutOneGrupoSwagger struct {
	Status bool                                 `json:"Status"`
	Meta   string                               `json:"Meta"`
	Data   ResponseMessages.PutOneGrupoResponse `json:"Data"`
}

type DeleteGrupoSwagger struct {
	Status bool                                 `json:"Status"`
	Meta   string                               `json:"Meta"`
	Data   ResponseMessages.DeleteGrupoResponse `json:"Data"`
}
