package SwaggerMessages

import (
	"medroom-backend/ResponseMessages"
)

type ListAdministradoresAcademicosSwagger struct {
	Status bool                                                   `json:"status"`
	Meta   string                                                 `json:"meta"`
	Data   ResponseMessages.ListAdministradoresAcademicosResponse `json:"data"`
}

type GetOneAdministradorAcademicoSwagger struct {
	Status bool                                                  `json:"status"`
	Meta   string                                                `json:"meta"`
	Data   ResponseMessages.GetOneAdministradorAcademicoResponse `json:"data"`
}

type AddNewAdministradorAcademicoSwagger struct {
	Status bool                                                  `json:"status"`
	Meta   string                                                `json:"meta"`
	Data   ResponseMessages.AddNewAdministradorAcademicoResponse `json:"data"`
}

type PutOneAdministradorAcademicoSwagger struct {
	Status bool                                                  `json:"status"`
	Meta   string                                                `json:"meta"`
	Data   ResponseMessages.PutOneAdministradorAcademicoResponse `json:"data"`
}

type DeleteAdministradorAcademicoSwagger struct {
	Status bool                                                  `json:"status"`
	Meta   string                                                `json:"meta"`
	Data   ResponseMessages.DeleteAdministradorAcademicoResponse `json:"data"`
}
