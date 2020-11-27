package SwaggerMessages

import (
	"medroom-backend/ResponseMessages"
)

type ListCompetenciasSwagger struct {
	Status bool                                      `json:"Status"`
	Meta   string                                    `json:"Meta"`
	Data   ResponseMessages.ListCompetenciasResponse `json:"Data"`
}

type GetOneCompetenciaSwagger struct {
	Status bool                                       `json:"Status"`
	Meta   string                                     `json:"Meta"`
	Data   ResponseMessages.GetOneCompetenciaResponse `json:"Data"`
}

type AddNewCompetenciaSwagger struct {
	Status bool                                       `json:"Status"`
	Meta   string                                     `json:"Meta"`
	Data   ResponseMessages.AddNewCompetenciaResponse `json:"Data"`
}

type PutOneCompetenciaSwagger struct {
	Status bool                                       `json:"Status"`
	Meta   string                                     `json:"Meta"`
	Data   ResponseMessages.PutOneCompetenciaResponse `json:"Data"`
}

type DeleteCompetenciaSwagger struct {
	Status bool                                       `json:"Status"`
	Meta   string                                     `json:"Meta"`
	Data   ResponseMessages.DeleteCompetenciaResponse `json:"Data"`
}
