package Output

import (
	"medroom-backend/Messages/Response"
	"medroom-backend/Models"
)

func GetCursosOutput(u []Models.Curso) (output []Response.ListCursosResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListCursosResponse{
			Periodo_curso: GetOnePeriodoOutput(u[i].Periodo_curso),
			Grupos_curso:  GetGruposOutput(u[i].Grupos_curso),
			Nombre_curso:  u[i].Nombre_curso,
			Sigla_curso:   u[i].Sigla_curso,
		})
	}

	return output
}

func GetOneCursoOutput(u Models.Curso) (output Response.GetOneCursoResponse) {
	return Response.GetOneCursoResponse{
		Periodo_curso: GetOnePeriodoOutput(u.Periodo_curso),
		Grupos_curso:  GetGruposOutput(u.Grupos_curso),
		Nombre_curso:  u.Nombre_curso,
		Sigla_curso:   u.Sigla_curso,
	}
}

func GetCursoEstudianteOutput(u Models.Curso) (output Response.GetCursoEstudianteResponse) {
	return Response.GetCursoEstudianteResponse{
		Periodo_curso: GetOnePeriodoOutput(u.Periodo_curso),
		Grupos_curso:  GetGruposOutput(u.Grupos_curso),
		Nombre_curso:  u.Nombre_curso,
		Sigla_curso:   u.Sigla_curso,
	}
}

func GetCursoEvaluadorOutput(u Models.Curso) (output Response.GetCursoEvaluadorResponse) {
	return Response.GetCursoEvaluadorResponse{
		Periodo_curso: GetOnePeriodoOutput(u.Periodo_curso),
		Grupos_curso:  GetGruposOutput(u.Grupos_curso),
		Nombre_curso:  u.Nombre_curso,
		Sigla_curso:   u.Sigla_curso,
	}
}

func AddNewCursoOutput(u Models.Curso) (output Response.AddNewCursoResponse) {
	return Response.AddNewCursoResponse{
		Id_periodo:   u.Id_periodo,
		Nombre_curso: u.Nombre_curso,
		Sigla_curso:  u.Sigla_curso,
	}
}

func PutOneCursoOutput(u Models.Curso) (output Response.PutOneCursoResponse) {
	return Response.PutOneCursoResponse{
		Id_periodo:   u.Id_periodo,
		Nombre_curso: u.Nombre_curso,
		Sigla_curso:  u.Sigla_curso,
	}
}

func DeleteCursoOutput(u Models.Curso) (output Response.DeleteCursoResponse) {
	return Response.DeleteCursoResponse{
		Id_periodo:   u.Id_periodo,
		Nombre_curso: u.Nombre_curso,
		Sigla_curso:  u.Sigla_curso,
	}
}
