package SwaggerMessages

import (
	"medroom-backend/ResponseMessages"
)

type ListGruposSwagger struct {
	Status bool                                `json:"status"`
	Meta   string                              `json:"meta"`
	Data   ResponseMessages.ListGruposResponse `json:"data"`
}

type GetOneGrupoSwagger struct {
	Status bool                                 `json:"status"`
	Meta   string                               `json:"meta"`
	Data   ResponseMessages.GetOneGrupoResponse `json:"data"`
}

type GetGrupoEstudianteSwagger struct {
	Status bool                                        `json:"status"`
	Meta   string                                      `json:"meta"`
	Data   ResponseMessages.GetGrupoEstudianteResponse `json:"data"`
}

type AddNewGrupoSwagger struct {
	Status bool                                 `json:"status"`
	Meta   string                               `json:"meta"`
	Data   ResponseMessages.AddNewGrupoResponse `json:"data"`
}

type PutOneGrupoSwagger struct {
	Status bool                                 `json:"status"`
	Meta   string                               `json:"meta"`
	Data   ResponseMessages.PutOneGrupoResponse `json:"data"`
}

type DeleteGrupoSwagger struct {
	Status bool                                 `json:"status"`
	Meta   string                               `json:"meta"`
	Data   ResponseMessages.DeleteGrupoResponse `json:"data"`
}
