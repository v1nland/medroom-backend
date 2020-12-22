package SwaggerMessages

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

type GetGrupoEstudianteSwagger struct {
	Status bool                                `json:"status"`
	Meta   string                              `json:"meta"`
	Data   Response.GetGrupoEstudianteResponse `json:"data"`
}

type GetGrupoEvaluadorSwagger struct {
	Status bool                               `json:"status"`
	Meta   string                             `json:"meta"`
	Data   Response.GetGrupoEvaluadorResponse `json:"data"`
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
