package Swagger

import "medroom-backend/Messages/Response"

type EvolucionEstudiantePorEvaluacionSwagger struct {
	Status bool                                              `json:"status"`
	Meta   string                                            `json:"meta"`
	Data   Response.EvolucionEstudiantePorEvaluacionResponse `json:"data"`
}

type EvolucionEstudiantePorCompetenciaSwagger struct {
	Status bool                                               `json:"status"`
	Meta   string                                             `json:"meta"`
	Data   Response.EvolucionEstudiantePorCompetenciaResponse `json:"data"`
}
