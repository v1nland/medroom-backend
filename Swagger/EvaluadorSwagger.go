package Swagger

import "medroom-backend/Messages/Response"

type ListEvaluadoresSwagger struct {
	Status bool                             `json:"status"`
	Meta   string                           `json:"meta"`
	Data   Response.ListEvaluadoresResponse `json:"data"`
}

type GetOneEvaluadorSwagger struct {
	Status bool                             `json:"status"`
	Meta   string                           `json:"meta"`
	Data   Response.GetOneEvaluadorResponse `json:"data"`
}

type AddNewEvaluadorSwagger struct {
	Status bool                             `json:"status"`
	Meta   string                           `json:"meta"`
	Data   Response.AddNewEvaluadorResponse `json:"data"`
}

type PutOneEvaluadorSwagger struct {
	Status bool                             `json:"status"`
	Meta   string                           `json:"meta"`
	Data   Response.PutOneEvaluadorResponse `json:"data"`
}

type DeleteEvaluadorSwagger struct {
	Status bool                             `json:"status"`
	Meta   string                           `json:"meta"`
	Data   Response.DeleteEvaluadorResponse `json:"data"`
}

type GetMyEvaluadorSwagger struct {
	Status bool                            `json:"status"`
	Meta   string                          `json:"meta"`
	Data   Response.GetMyEvaluadorResponse `json:"data"`
}

type PutMyEvaluadorSwagger struct {
	Status bool                            `json:"status"`
	Meta   string                          `json:"meta"`
	Data   Response.PutMyEvaluadorResponse `json:"data"`
}

type AddEvaluadorToGrupoSwagger struct {
	Status bool                                 `json:"status"`
	Meta   string                               `json:"meta"`
	Data   Response.AddEvaluadorToGrupoResponse `json:"data"`
}
