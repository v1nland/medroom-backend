package OutputFormats

import (
	"medroom-backend/Models"
	"medroom-backend/ResponseMessages"
)

func GetPeriodosOutput(u []Models.Periodo) (output []ResponseMessages.ListPeriodosResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, ResponseMessages.ListPeriodosResponse{
			Nombre_periodo: u[i].Nombre_periodo,
		})
	}

	return output
}

func GetOnePeriodoOutput(u Models.Periodo) (output ResponseMessages.GetOnePeriodoResponse) {
	return ResponseMessages.GetOnePeriodoResponse{
		Nombre_periodo: u.Nombre_periodo,
	}
}

func AddNewPeriodoOutput(u Models.Periodo) (output ResponseMessages.AddNewPeriodoResponse) {
	return ResponseMessages.AddNewPeriodoResponse{
		Nombre_periodo: u.Nombre_periodo,
	}
}

func PutOnePeriodoOutput(u Models.Periodo) (output ResponseMessages.PutOnePeriodoResponse) {
	return ResponseMessages.PutOnePeriodoResponse{
		Nombre_periodo: u.Nombre_periodo,
	}
}

func DeletePeriodoOutput(u Models.Periodo) (output ResponseMessages.DeletePeriodoResponse) {
	return ResponseMessages.DeletePeriodoResponse{
		Nombre_periodo: u.Nombre_periodo,
	}
}
