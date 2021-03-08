package Output

import (
	"medroom-backend/Messages/Response"
	"medroom-backend/models"
)

func ListCursosOutput(u []models.Curso) (output []Response.ListCursosResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListCursosResponse{
			Periodo_curso: GetOnePeriodoOutput(u[i].Periodo_curso),
			Grupos_curso:  ListGruposOutput(u[i].Grupos_curso),
			Nombre_curso:  u[i].Nombre_curso,
			Sigla_curso:   u[i].Sigla_curso,
		})
	}

	return output
}

func GetOneCursoOutput(u models.Curso) (output Response.GetOneCursoResponse) {
	return Response.GetOneCursoResponse{
		Periodo_curso: GetOnePeriodoOutput(u.Periodo_curso),
		Grupos_curso:  ListGruposOutput(u.Grupos_curso),
		Nombre_curso:  u.Nombre_curso,
		Sigla_curso:   u.Sigla_curso,
	}
}

func GetCursosEstudianteOutput(u []models.Curso) (output []Response.GetCursosEstudianteResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.GetCursosEstudianteResponse{
			Periodo_curso: GetOnePeriodoOutput(u[i].Periodo_curso),
			Grupos_curso:  ListGruposOutput(u[i].Grupos_curso),
			Nombre_curso:  u[i].Nombre_curso,
			Sigla_curso:   u[i].Sigla_curso,
		})
	}

	return output
}

func GetOneCursoEstudianteOutput(u models.Curso) (output Response.GetOneCursoEstudianteResponse) {
	return Response.GetOneCursoEstudianteResponse{
		Periodo_curso: GetOnePeriodoOutput(u.Periodo_curso),
		Grupos_curso:  ListGruposOutput(u.Grupos_curso),
		Nombre_curso:  u.Nombre_curso,
		Sigla_curso:   u.Sigla_curso,
	}
}

func GetCursosEvaluadorOutput(u []models.Curso) (output []Response.GetCursosEvaluadorResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.GetCursosEvaluadorResponse{
			Periodo_curso: GetOnePeriodoOutput(u[i].Periodo_curso),
			Grupos_curso:  ListGruposOutput(u[i].Grupos_curso),
			Nombre_curso:  u[i].Nombre_curso,
			Sigla_curso:   u[i].Sigla_curso,
		})
	}

	return output
}

func GetOneCursoEvaluadorOutput(u models.Curso) (output Response.GetOneCursoEvaluadorResponse) {
	return Response.GetOneCursoEvaluadorResponse{
		Periodo_curso: GetOnePeriodoOutput(u.Periodo_curso),
		Grupos_curso:  ListGruposOutput(u.Grupos_curso),
		Nombre_curso:  u.Nombre_curso,
		Sigla_curso:   u.Sigla_curso,
	}
}

func AddNewCursoOutput(u models.Curso) (output Response.AddNewCursoResponse) {
	return Response.AddNewCursoResponse{
		Id: u.Id,
	}
}

func PutOneCursoOutput(u models.Curso) (output Response.PutOneCursoResponse) {
	return Response.PutOneCursoResponse{
		Id: u.Id,
	}
}

func DeleteCursoOutput(u models.Curso) (output Response.DeleteCursoResponse) {
	return Response.DeleteCursoResponse{
		Id: u.Id,
	}
}
