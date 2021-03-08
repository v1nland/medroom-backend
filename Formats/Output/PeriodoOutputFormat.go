package Output

import (
	"medroom-backend/Messages/Response"
	"medroom-backend/models"
)

func ListPeriodosOutput(u []models.Periodo) (output []Response.ListPeriodosResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListPeriodosResponse{
			Id:             u[i].Id,
			Nombre_periodo: u[i].Nombre_periodo,
		})
	}

	return output
}

func GetOnePeriodoOutput(u models.Periodo) (output Response.GetOnePeriodoResponse) {
	return Response.GetOnePeriodoResponse{
		Id:             u.Id,
		Nombre_periodo: u.Nombre_periodo,
	}
}

func AddNewPeriodoOutput(u models.Periodo) (output Response.AddNewPeriodoResponse) {
	return Response.AddNewPeriodoResponse{
		Id: u.Id,
	}
}

func PutOnePeriodoOutput(u models.Periodo) (output Response.PutOnePeriodoResponse) {
	return Response.PutOnePeriodoResponse{
		Id: u.Id,
	}
}

func DeletePeriodoOutput(u models.Periodo) (output Response.DeletePeriodoResponse) {
	return Response.DeletePeriodoResponse{
		Id: u.Id,
	}
}
