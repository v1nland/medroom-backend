package SwaggerMessages

import (
	"medroom-backend/ResponseMessages"
)

type ListEvaluacionesSwagger struct {
	Status bool                                      `json:"status"`
	Meta   string                                    `json:"meta"`
	Data   ResponseMessages.ListEvaluacionesResponse `json:"data"`
}

type GetOneEvaluacionSwagger struct {
	Status bool                                      `json:"status"`
	Meta   string                                    `json:"meta"`
	Data   ResponseMessages.GetOneEvaluacionResponse `json:"data"`
}

type AddNewEvaluacionSwagger struct {
	Status bool                                      `json:"status"`
	Meta   string                                    `json:"meta"`
	Data   ResponseMessages.AddNewEvaluacionResponse `json:"data"`
}

type PutOneEvaluacionSwagger struct {
	Status bool                                      `json:"status"`
	Meta   string                                    `json:"meta"`
	Data   ResponseMessages.PutOneEvaluacionResponse `json:"data"`
}

type DeleteEvaluacionSwagger struct {
	Status bool                                      `json:"status"`
	Meta   string                                    `json:"meta"`
	Data   ResponseMessages.DeleteEvaluacionResponse `json:"data"`
}
