package SwaggerMessages

import (
	"medroom-backend/ResponseMessages"
)

type EvolucionEstudiantePorEvaluacionSwagger struct {
	Status bool                                                      `json:"status"`
	Meta   string                                                    `json:"meta"`
	Data   ResponseMessages.EvolucionEstudiantePorEvaluacionResponse `json:"data"`
}

type EvolucionEstudiantePorCompetenciaSwagger struct {
	Status bool                                                       `json:"status"`
	Meta   string                                                     `json:"meta"`
	Data   ResponseMessages.EvolucionEstudiantePorCompetenciaResponse `json:"data"`
}
