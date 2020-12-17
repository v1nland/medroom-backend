package SwaggerMessages

import (
	"medroom-backend/ResponseMessages"
)

type ListEvaluadoresSwagger struct {
	Status bool                                     `json:"status"`
	Meta   string                                   `json:"meta"`
	Data   ResponseMessages.ListEvaluadoresResponse `json:"data"`
}

type GetOneEvaluadorSwagger struct {
	Status bool                                     `json:"status"`
	Meta   string                                   `json:"meta"`
	Data   ResponseMessages.GetOneEvaluadorResponse `json:"data"`
}

type AddNewEvaluadorSwagger struct {
	Status bool                                     `json:"status"`
	Meta   string                                   `json:"meta"`
	Data   ResponseMessages.AddNewEvaluadorResponse `json:"data"`
}

type PutOneEvaluadorSwagger struct {
	Status bool                                     `json:"status"`
	Meta   string                                   `json:"meta"`
	Data   ResponseMessages.PutOneEvaluadorResponse `json:"data"`
}

type DeleteEvaluadorSwagger struct {
	Status bool                                     `json:"status"`
	Meta   string                                   `json:"meta"`
	Data   ResponseMessages.DeleteEvaluadorResponse `json:"data"`
}

type GetMyEvaluadorSwagger struct {
	Status bool                                    `json:"status"`
	Meta   string                                  `json:"meta"`
	Data   ResponseMessages.GetMyEvaluadorResponse `json:"data"`
}

type PutMyEvaluadorSwagger struct {
	Status bool                                    `json:"status"`
	Meta   string                                  `json:"meta"`
	Data   ResponseMessages.PutMyEvaluadorResponse `json:"data"`
}
