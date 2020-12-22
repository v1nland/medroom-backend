package Output

import (
	"medroom-backend/Messages/Response"
	"medroom-backend/Models"
)

func ListPeriodosOutput(u []Models.Periodo) (output []Response.ListPeriodosResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListPeriodosResponse{
			Id:             u[i].Id,
			Nombre_periodo: u[i].Nombre_periodo,
		})
	}

	return output
}

func GetOnePeriodoOutput(u Models.Periodo) (output Response.GetOnePeriodoResponse) {
	return Response.GetOnePeriodoResponse{
		Id:             u.Id,
		Nombre_periodo: u.Nombre_periodo,
	}
}

func AddNewPeriodoOutput(u Models.Periodo) (output Response.AddNewPeriodoResponse) {
	return Response.AddNewPeriodoResponse{
		Id:             u.Id,
		Nombre_periodo: u.Nombre_periodo,
	}
}

func PutOnePeriodoOutput(u Models.Periodo) (output Response.PutOnePeriodoResponse) {
	return Response.PutOnePeriodoResponse{
		Id:             u.Id,
		Nombre_periodo: u.Nombre_periodo,
	}
}

func DeletePeriodoOutput(u Models.Periodo) (output Response.DeletePeriodoResponse) {
	return Response.DeletePeriodoResponse{
		Id:             u.Id,
		Nombre_periodo: u.Nombre_periodo,
	}
}
