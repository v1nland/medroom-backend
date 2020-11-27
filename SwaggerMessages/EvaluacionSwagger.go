package SwaggerMessages

import (
	"medroom-backend/ResponseMessages"
)

type ListEvaluacionesSwagger struct {
	Status bool                                      `json:"Status"`
	Meta   string                                    `json:"Meta"`
	Data   ResponseMessages.ListEvaluacionesResponse `json:"Data"`
}

type GetOneEvaluacionSwagger struct {
	Status bool                                      `json:"Status"`
	Meta   string                                    `json:"Meta"`
	Data   ResponseMessages.GetOneEvaluacionResponse `json:"Data"`
}

type AddNewEvaluacionSwagger struct {
	Status bool                                      `json:"Status"`
	Meta   string                                    `json:"Meta"`
	Data   ResponseMessages.AddNewEvaluacionResponse `json:"Data"`
}

type PutOneEvaluacionSwagger struct {
	Status bool                                      `json:"Status"`
	Meta   string                                    `json:"Meta"`
	Data   ResponseMessages.PutOneEvaluacionResponse `json:"Data"`
}

type DeleteEvaluacionSwagger struct {
	Status bool                                      `json:"Status"`
	Meta   string                                    `json:"Meta"`
	Data   ResponseMessages.DeleteEvaluacionResponse `json:"Data"`
}
