package Swagger

import "medroom-backend/app/messages/Response"

type ListCursosSwagger struct {
	Status bool                        `json:"status"`
	Meta   string                      `json:"meta"`
	Data   Response.ListCursosResponse `json:"data"`
}

type GetOneCursoSwagger struct {
	Status bool                         `json:"status"`
	Meta   string                       `json:"meta"`
	Data   Response.GetOneCursoResponse `json:"data"`
}

type GetCursosEstudianteSwagger struct {
	Status bool                                 `json:"status"`
	Meta   string                               `json:"meta"`
	Data   Response.GetCursosEstudianteResponse `json:"data"`
}

type GetOneCursoEstudianteSwagger struct {
	Status bool                                   `json:"status"`
	Meta   string                                 `json:"meta"`
	Data   Response.GetOneCursoEstudianteResponse `json:"data"`
}

type GetCursosEvaluadorSwagger struct {
	Status bool                                `json:"status"`
	Meta   string                              `json:"meta"`
	Data   Response.GetCursosEvaluadorResponse `json:"data"`
}

type GetOneCursoEvaluadorSwagger struct {
	Status bool                                  `json:"status"`
	Meta   string                                `json:"meta"`
	Data   Response.GetOneCursoEvaluadorResponse `json:"data"`
}

type AddNewCursoSwagger struct {
	Status bool                         `json:"status"`
	Meta   string                       `json:"meta"`
	Data   Response.AddNewCursoResponse `json:"data"`
}

type PutOneCursoSwagger struct {
	Status bool                         `json:"status"`
	Meta   string                       `json:"meta"`
	Data   Response.PutOneCursoResponse `json:"data"`
}

type DeleteCursoSwagger struct {
	Status bool                         `json:"status"`
	Meta   string                       `json:"meta"`
	Data   Response.DeleteCursoResponse `json:"data"`
}

type AddEstudianteToCursoSwagger struct {
	Status bool                                  `json:"status"`
	Meta   string                                `json:"meta"`
	Data   Response.AddEstudianteToCursoResponse `json:"data"`
}

type AddEvaluadorToCursoSwagger struct {
	Status bool                                 `json:"status"`
	Meta   string                               `json:"meta"`
	Data   Response.AddEvaluadorToCursoResponse `json:"data"`
}

type AddAdministradorAcademicoToCursoSwagger struct {
	Status bool                                              `json:"status"`
	Meta   string                                            `json:"meta"`
	Data   Response.AddAdministradorAcademicoToCursoResponse `json:"data"`
}

type GetCursosAdministradorAcademicoSwagger struct {
	Status bool                                             `json:"status"`
	Meta   string                                           `json:"meta"`
	Data   Response.GetCursosAdministradorAcademicoResponse `json:"data"`
}

type GetOneCursoAdministradorAcademicoSwagger struct {
	Status bool                                               `json:"status"`
	Meta   string                                             `json:"meta"`
	Data   Response.GetOneCursoAdministradorAcademicoResponse `json:"data"`
}
