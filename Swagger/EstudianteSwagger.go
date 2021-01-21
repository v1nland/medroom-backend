package Swagger

import "medroom-backend/Messages/Response"

type ListEstudiantesSwagger struct {
	Status bool                             `json:"status"`
	Meta   string                           `json:"meta"`
	Data   Response.ListEstudiantesResponse `json:"data"`
}

type GetOneEstudianteSwagger struct {
	Status bool                              `json:"status"`
	Meta   string                            `json:"meta"`
	Data   Response.GetOneEstudianteResponse `json:"data"`
}

type AddNewEstudianteSwagger struct {
	Status bool                              `json:"status"`
	Meta   string                            `json:"meta"`
	Data   Response.AddNewEstudianteResponse `json:"data"`
}

type PutOneEstudianteSwagger struct {
	Status bool                              `json:"status"`
	Meta   string                            `json:"meta"`
	Data   Response.PutOneEstudianteResponse `json:"data"`
}

type DeleteEstudianteSwagger struct {
	Status bool                              `json:"status"`
	Meta   string                            `json:"meta"`
	Data   Response.DeleteEstudianteResponse `json:"data"`
}

type GetMyEstudianteSwagger struct {
	Status bool                             `json:"status"`
	Meta   string                           `json:"meta"`
	Data   Response.GetMyEstudianteResponse `json:"data"`
}

type PutMyEstudianteSwagger struct {
	Status bool                             `json:"status"`
	Meta   string                           `json:"meta"`
	Data   Response.PutMyEstudianteResponse `json:"data"`
}

type AddEstudianteToGrupoSwagger struct {
	Status bool                                  `json:"status"`
	Meta   string                                `json:"meta"`
	Data   Response.AddEstudianteToGrupoResponse `json:"data"`
}
