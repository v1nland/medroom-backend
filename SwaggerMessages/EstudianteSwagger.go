package SwaggerMessages

import (
	"medroom-backend/ResponseMessages"
)

type ListEstudiantesSwagger struct {
	Status bool                                     `json:"Status"`
	Meta   string                                   `json:"Meta"`
	Data   ResponseMessages.ListEstudiantesResponse `json:"Data"`
}

type GetOneEstudianteSwagger struct {
	Status bool                                      `json:"Status"`
	Meta   string                                    `json:"Meta"`
	Data   ResponseMessages.GetOneEstudianteResponse `json:"Data"`
}

type AddNewEstudianteSwagger struct {
	Status bool                                      `json:"Status"`
	Meta   string                                    `json:"Meta"`
	Data   ResponseMessages.AddNewEstudianteResponse `json:"Data"`
}

type PutOneEstudianteSwagger struct {
	Status bool                                      `json:"Status"`
	Meta   string                                    `json:"Meta"`
	Data   ResponseMessages.PutOneEstudianteResponse `json:"Data"`
}

type DeleteEstudianteSwagger struct {
	Status bool                                      `json:"Status"`
	Meta   string                                    `json:"Meta"`
	Data   ResponseMessages.DeleteEstudianteResponse `json:"Data"`
}
