package f_output

import (
	"medroom-backend/app/Messages/Response"
	"medroom-backend/app/models"
)

func ListPeriodos(u []models.Periodo) (output []Response.ListPeriodosResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListPeriodosResponse{
			Id:             u[i].Id,
			Nombre_periodo: u[i].Nombre_periodo,
		})
	}

	return output
}

func GetOnePeriodo(u models.Periodo) (output Response.GetOnePeriodoResponse) {
	return Response.GetOnePeriodoResponse{
		Id:             u.Id,
		Nombre_periodo: u.Nombre_periodo,
	}
}

func AddNewPeriodo(u models.Periodo) (output Response.AddNewPeriodoResponse) {
	return Response.AddNewPeriodoResponse{
		Id: u.Id,
	}
}

func PutOnePeriodo(u models.Periodo) (output Response.PutOnePeriodoResponse) {
	return Response.PutOnePeriodoResponse{
		Id: u.Id,
	}
}

func DeletePeriodo(u models.Periodo) (output Response.DeletePeriodoResponse) {
	return Response.DeletePeriodoResponse{
		Id: u.Id,
	}
}
