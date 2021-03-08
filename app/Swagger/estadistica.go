package Swagger

import (
	"medroom-backend/app/messages/Response"
)

type EvolucionEstudiantePorCompetenciaSwagger struct {
	Status bool                                               `json:"status"`
	Meta   string                                             `json:"meta"`
	Data   Response.EvolucionEstudiantePorCompetenciaResponse `json:"data"`
}

type EvolucionEstudiantePorEvaluacionSwagger struct {
	Status bool                                              `json:"status"`
	Meta   string                                            `json:"meta"`
	Data   Response.EvolucionEstudiantePorEvaluacionResponse `json:"data"`
}
