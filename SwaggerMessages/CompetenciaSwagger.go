package SwaggerMessages

import (
	"medroom-backend/ResponseMessages"
)

type ListCompetenciasSwagger struct {
	Status bool                                      `json:"status"`
	Meta   string                                    `json:"meta"`
	Data   ResponseMessages.ListCompetenciasResponse `json:"data"`
}

type GetOneCompetenciaSwagger struct {
	Status bool                                       `json:"status"`
	Meta   string                                     `json:"meta"`
	Data   ResponseMessages.GetOneCompetenciaResponse `json:"data"`
}

type AddNewCompetenciaSwagger struct {
	Status bool                                       `json:"status"`
	Meta   string                                     `json:"meta"`
	Data   ResponseMessages.AddNewCompetenciaResponse `json:"data"`
}

type PutOneCompetenciaSwagger struct {
	Status bool                                       `json:"status"`
	Meta   string                                     `json:"meta"`
	Data   ResponseMessages.PutOneCompetenciaResponse `json:"data"`
}

type DeleteCompetenciaSwagger struct {
	Status bool                                       `json:"status"`
	Meta   string                                     `json:"meta"`
	Data   ResponseMessages.DeleteCompetenciaResponse `json:"data"`
}
