package OutputFormats

import (
	"medroom-backend/Models"
	"medroom-backend/ResponseMessages"
)

func GetCursosOutput(u []Models.Curso) (output []ResponseMessages.ListCursosResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, ResponseMessages.ListCursosResponse{
			Periodo_curso: GetOnePeriodoOutput(u[i].Periodo_curso),
			Nombre_curso:  u[i].Nombre_curso,
			Sigla_curso:   u[i].Sigla_curso,
		})
	}

	return output
}

func GetOneCursoOutput(u Models.Curso) (output ResponseMessages.GetOneCursoResponse) {
	return ResponseMessages.GetOneCursoResponse{
		Periodo_curso: GetOnePeriodoOutput(u.Periodo_curso),
		Nombre_curso:  u.Nombre_curso,
		Sigla_curso:   u.Sigla_curso,
	}
}

func AddNewCursoOutput(u Models.Curso) (output ResponseMessages.AddNewCursoResponse) {
	return ResponseMessages.AddNewCursoResponse{
		Id_periodo:   u.Id_periodo,
		Nombre_curso: u.Nombre_curso,
		Sigla_curso:  u.Sigla_curso,
	}
}

func PutOneCursoOutput(u Models.Curso) (output ResponseMessages.PutOneCursoResponse) {
	return ResponseMessages.PutOneCursoResponse{
		Id_periodo:   u.Id_periodo,
		Nombre_curso: u.Nombre_curso,
		Sigla_curso:  u.Sigla_curso,
	}
}

func DeleteCursoOutput(u Models.Curso) (output ResponseMessages.DeleteCursoResponse) {
	return ResponseMessages.DeleteCursoResponse{
		Id_periodo:   u.Id_periodo,
		Nombre_curso: u.Nombre_curso,
		Sigla_curso:  u.Sigla_curso,
	}
}
