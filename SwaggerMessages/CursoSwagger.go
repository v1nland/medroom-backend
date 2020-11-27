package SwaggerMessages

import (
	"medroom-backend/ResponseMessages"
)

type ListCursosSwagger struct {
	Status bool                                `json:"Status"`
	Meta   string                              `json:"Meta"`
	Data   ResponseMessages.ListCursosResponse `json:"Data"`
}

type GetOneCursoSwagger struct {
	Status bool                                 `json:"Status"`
	Meta   string                               `json:"Meta"`
	Data   ResponseMessages.GetOneCursoResponse `json:"Data"`
}

type AddNewCursoSwagger struct {
	Status bool                                 `json:"Status"`
	Meta   string                               `json:"Meta"`
	Data   ResponseMessages.AddNewCursoResponse `json:"Data"`
}

type PutOneCursoSwagger struct {
	Status bool                                 `json:"Status"`
	Meta   string                               `json:"Meta"`
	Data   ResponseMessages.PutOneCursoResponse `json:"Data"`
}

type DeleteCursoSwagger struct {
	Status bool                                 `json:"Status"`
	Meta   string                               `json:"Meta"`
	Data   ResponseMessages.DeleteCursoResponse `json:"Data"`
}
