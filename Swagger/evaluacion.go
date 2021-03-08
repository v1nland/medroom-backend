package Swagger

import (
	"medroom-backend/Messages/Response"
)

type ListEvaluacionesGrupoEstudianteSwagger struct {
	Status bool                                             `json:"status"`
	Meta   string                                           `json:"meta"`
	Data   Response.ListEvaluacionesGrupoEstudianteResponse `json:"data"`
}

type ListEvaluacionesGrupoEvaluadorSwagger struct {
	Status bool                                            `json:"status"`
	Meta   string                                          `json:"meta"`
	Data   Response.ListEvaluacionesGrupoEvaluadorResponse `json:"data"`
}

type ListEvaluacionesSwagger struct {
	Status bool                              `json:"status"`
	Meta   string                            `json:"meta"`
	Data   Response.ListEvaluacionesResponse `json:"data"`
}

type GetOneEvaluacionSwagger struct {
	Status bool                              `json:"status"`
	Meta   string                            `json:"meta"`
	Data   Response.GetOneEvaluacionResponse `json:"data"`
}

type AddNewEvaluacionSwagger struct {
	Status bool                              `json:"status"`
	Meta   string                            `json:"meta"`
	Data   Response.AddNewEvaluacionResponse `json:"data"`
}

type PutOneEvaluacionSwagger struct {
	Status bool                              `json:"status"`
	Meta   string                            `json:"meta"`
	Data   Response.PutOneEvaluacionResponse `json:"data"`
}

type DeleteEvaluacionSwagger struct {
	Status bool                              `json:"status"`
	Meta   string                            `json:"meta"`
	Data   Response.DeleteEvaluacionResponse `json:"data"`
}
