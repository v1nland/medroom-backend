package SwaggerMessages

import (
	"medroom-backend/ResponseMessages"
)

type ListEstudiantesSwagger struct {
	Status bool                                     `json:"status"`
	Meta   string                                   `json:"meta"`
	Data   ResponseMessages.ListEstudiantesResponse `json:"data"`
}

type GetOneEstudianteSwagger struct {
	Status bool                                      `json:"status"`
	Meta   string                                    `json:"meta"`
	Data   ResponseMessages.GetOneEstudianteResponse `json:"data"`
}

type AddNewEstudianteSwagger struct {
	Status bool                                      `json:"status"`
	Meta   string                                    `json:"meta"`
	Data   ResponseMessages.AddNewEstudianteResponse `json:"data"`
}

type PutOneEstudianteSwagger struct {
	Status bool                                      `json:"status"`
	Meta   string                                    `json:"meta"`
	Data   ResponseMessages.PutOneEstudianteResponse `json:"data"`
}

type DeleteEstudianteSwagger struct {
	Status bool                                      `json:"status"`
	Meta   string                                    `json:"meta"`
	Data   ResponseMessages.DeleteEstudianteResponse `json:"data"`
}

type GetMyEstudianteSwagger struct {
	Status bool                                     `json:"status"`
	Meta   string                                   `json:"meta"`
	Data   ResponseMessages.GetMyEstudianteResponse `json:"data"`
}

type PutMyEstudianteSwagger struct {
	Status bool                                     `json:"status"`
	Meta   string                                   `json:"meta"`
	Data   ResponseMessages.PutMyEstudianteResponse `json:"data"`
}
