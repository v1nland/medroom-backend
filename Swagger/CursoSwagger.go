package Swagger

import "medroom-backend/Messages/Response"

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

type GetCursoEstudianteSwagger struct {
	Status bool                                `json:"status"`
	Meta   string                              `json:"meta"`
	Data   Response.GetCursoEstudianteResponse `json:"data"`
}

type GetCursoEvaluadorSwagger struct {
	Status bool                               `json:"status"`
	Meta   string                             `json:"meta"`
	Data   Response.GetCursoEvaluadorResponse `json:"data"`
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
