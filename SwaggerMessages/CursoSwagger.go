package SwaggerMessages

import (
	"medroom-backend/ResponseMessages"
)

type ListCursosSwagger struct {
	Status bool                                `json:"status"`
	Meta   string                              `json:"meta"`
	Data   ResponseMessages.ListCursosResponse `json:"data"`
}

type GetOneCursoSwagger struct {
	Status bool                                 `json:"status"`
	Meta   string                               `json:"meta"`
	Data   ResponseMessages.GetOneCursoResponse `json:"data"`
}

type GetCursoEstudianteSwagger struct {
	Status bool                                        `json:"status"`
	Meta   string                                      `json:"meta"`
	Data   ResponseMessages.GetCursoEstudianteResponse `json:"data"`
}

type GetCursoEvaluadorSwagger struct {
	Status bool                                       `json:"status"`
	Meta   string                                     `json:"meta"`
	Data   ResponseMessages.GetCursoEvaluadorResponse `json:"data"`
}

type AddNewCursoSwagger struct {
	Status bool                                 `json:"status"`
	Meta   string                               `json:"meta"`
	Data   ResponseMessages.AddNewCursoResponse `json:"data"`
}

type PutOneCursoSwagger struct {
	Status bool                                 `json:"status"`
	Meta   string                               `json:"meta"`
	Data   ResponseMessages.PutOneCursoResponse `json:"data"`
}

type DeleteCursoSwagger struct {
	Status bool                                 `json:"status"`
	Meta   string                               `json:"meta"`
	Data   ResponseMessages.DeleteCursoResponse `json:"data"`
}
