package Swagger

import "medroom-backend/app/Messages/Response"

type ListAdministradoresAcademicosSwagger struct {
	Status bool                                           `json:"status"`
	Meta   string                                         `json:"meta"`
	Data   Response.ListAdministradoresAcademicosResponse `json:"data"`
}

type GetOneAdministradorAcademicoSwagger struct {
	Status bool                                          `json:"status"`
	Meta   string                                        `json:"meta"`
	Data   Response.GetOneAdministradorAcademicoResponse `json:"data"`
}

type GetMyAdministradorAcademicoSwagger struct {
	Status bool                                         `json:"status"`
	Meta   string                                       `json:"meta"`
	Data   Response.GetMyAdministradorAcademicoResponse `json:"data"`
}

type AddNewAdministradorAcademicoSwagger struct {
	Status bool                                          `json:"status"`
	Meta   string                                        `json:"meta"`
	Data   Response.AddNewAdministradorAcademicoResponse `json:"data"`
}

type PutOneAdministradorAcademicoSwagger struct {
	Status bool                                          `json:"status"`
	Meta   string                                        `json:"meta"`
	Data   Response.PutOneAdministradorAcademicoResponse `json:"data"`
}

type PutMyAdministradorAcademicoSwagger struct {
	Status bool                                         `json:"status"`
	Meta   string                                       `json:"meta"`
	Data   Response.PutMyAdministradorAcademicoResponse `json:"data"`
}

type DeleteAdministradorAcademicoSwagger struct {
	Status bool                                          `json:"status"`
	Meta   string                                        `json:"meta"`
	Data   Response.DeleteAdministradorAcademicoResponse `json:"data"`
}
