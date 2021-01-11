package Swagger

import "medroom-backend/Messages/Response"

type ListGruposSwagger struct {
	Status bool                        `json:"status"`
	Meta   string                      `json:"meta"`
	Data   Response.ListGruposResponse `json:"data"`
}

type GetOneGrupoSwagger struct {
	Status bool                         `json:"status"`
	Meta   string                       `json:"meta"`
	Data   Response.GetOneGrupoResponse `json:"data"`
}

type GetOneGrupoEstudianteSwagger struct {
	Status bool                                   `json:"status"`
	Meta   string                                 `json:"meta"`
	Data   Response.GetOneGrupoEstudianteResponse `json:"data"`
}

type GetGruposEstudianteSwagger struct {
	Status bool                                 `json:"status"`
	Meta   string                               `json:"meta"`
	Data   Response.GetGruposEstudianteResponse `json:"data"`
}

type GetOneGrupoEvaluadorSwagger struct {
	Status bool                                  `json:"status"`
	Meta   string                                `json:"meta"`
	Data   Response.GetOneGrupoEvaluadorResponse `json:"data"`
}

type GetGruposEvaluadorSwagger struct {
	Status bool                                `json:"status"`
	Meta   string                              `json:"meta"`
	Data   Response.GetGruposEvaluadorResponse `json:"data"`
}

type AddNewGrupoSwagger struct {
	Status bool                         `json:"status"`
	Meta   string                       `json:"meta"`
	Data   Response.AddNewGrupoResponse `json:"data"`
}

type PutOneGrupoSwagger struct {
	Status bool                         `json:"status"`
	Meta   string                       `json:"meta"`
	Data   Response.PutOneGrupoResponse `json:"data"`
}

type DeleteGrupoSwagger struct {
	Status bool                         `json:"status"`
	Meta   string                       `json:"meta"`
	Data   Response.DeleteGrupoResponse `json:"data"`
}
