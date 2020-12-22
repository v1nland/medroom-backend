package SwaggerMessages

import "medroom-backend/Messages/Response"

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

type GenerarEvaluacionSwagger struct {
	Status bool                               `json:"status"`
	Meta   string                             `json:"meta"`
	Data   Response.GenerarEvaluacionResponse `json:"data"`
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
