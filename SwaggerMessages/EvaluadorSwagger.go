package SwaggerMessages

import (
	"medroom-backend/ResponseMessages"
)

type ListEvaluadoresSwagger struct {
	Status bool                                     `json:"Status"`
	Meta   string                                   `json:"Meta"`
	Data   ResponseMessages.ListEvaluadoresResponse `json:"Data"`
}

type GetOneEvaluadorSwagger struct {
	Status bool                                     `json:"Status"`
	Meta   string                                   `json:"Meta"`
	Data   ResponseMessages.GetOneEvaluadorResponse `json:"Data"`
}

type AddNewEvaluadorSwagger struct {
	Status bool                                     `json:"Status"`
	Meta   string                                   `json:"Meta"`
	Data   ResponseMessages.AddNewEvaluadorResponse `json:"Data"`
}

type PutOneEvaluadorSwagger struct {
	Status bool                                     `json:"Status"`
	Meta   string                                   `json:"Meta"`
	Data   ResponseMessages.PutOneEvaluadorResponse `json:"Data"`
}

type DeleteEvaluadorSwagger struct {
	Status bool                                     `json:"Status"`
	Meta   string                                   `json:"Meta"`
	Data   ResponseMessages.DeleteEvaluadorResponse `json:"Data"`
}
