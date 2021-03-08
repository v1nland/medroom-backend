package Swagger

import (
	"medroom-backend/app/messages/Response"
)

type GetOneCalificacionEstudianteSwagger struct {
	Status bool                                          `json:"status"`
	Meta   string                                        `json:"meta"`
	Data   Response.GetOneCalificacionEstudianteResponse `json:"data"`
}

type AddNewCalificacionEstudianteSwagger struct {
	Status bool                                          `json:"status"`
	Meta   string                                        `json:"meta"`
	Data   Response.AddNewCalificacionEstudianteResponse `json:"data"`
}
