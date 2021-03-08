package f_output

import (
	"medroom-backend/app/Messages/Response"
	"medroom-backend/app/models"
)

func ListCursos(u []models.Curso) (output []Response.ListCursosResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.ListCursosResponse{
			Periodo_curso: GetOnePeriodo(u[i].Periodo_curso),
			Grupos_curso:  ListGrupos(u[i].Grupos_curso),
			Nombre_curso:  u[i].Nombre_curso,
			Sigla_curso:   u[i].Sigla_curso,
		})
	}

	return output
}

func GetOneCurso(u models.Curso) (output Response.GetOneCursoResponse) {
	return Response.GetOneCursoResponse{
		Periodo_curso: GetOnePeriodo(u.Periodo_curso),
		Grupos_curso:  ListGrupos(u.Grupos_curso),
		Nombre_curso:  u.Nombre_curso,
		Sigla_curso:   u.Sigla_curso,
	}
}

func GetCursosEstudiante(u []models.Curso) (output []Response.GetCursosEstudianteResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.GetCursosEstudianteResponse{
			Periodo_curso: GetOnePeriodo(u[i].Periodo_curso),
			Grupos_curso:  ListGrupos(u[i].Grupos_curso),
			Nombre_curso:  u[i].Nombre_curso,
			Sigla_curso:   u[i].Sigla_curso,
		})
	}

	return output
}

func GetOneCursoEstudiante(u models.Curso) (output Response.GetOneCursoEstudianteResponse) {
	return Response.GetOneCursoEstudianteResponse{
		Periodo_curso: GetOnePeriodo(u.Periodo_curso),
		Grupos_curso:  ListGrupos(u.Grupos_curso),
		Nombre_curso:  u.Nombre_curso,
		Sigla_curso:   u.Sigla_curso,
	}
}

func GetCursosEvaluador(u []models.Curso) (output []Response.GetCursosEvaluadorResponse) {
	for i := 0; i < len(u); i++ {
		output = append(output, Response.GetCursosEvaluadorResponse{
			Periodo_curso: GetOnePeriodo(u[i].Periodo_curso),
			Grupos_curso:  ListGrupos(u[i].Grupos_curso),
			Nombre_curso:  u[i].Nombre_curso,
			Sigla_curso:   u[i].Sigla_curso,
		})
	}

	return output
}

func GetOneCursoEvaluador(u models.Curso) (output Response.GetOneCursoEvaluadorResponse) {
	return Response.GetOneCursoEvaluadorResponse{
		Periodo_curso: GetOnePeriodo(u.Periodo_curso),
		Grupos_curso:  ListGrupos(u.Grupos_curso),
		Nombre_curso:  u.Nombre_curso,
		Sigla_curso:   u.Sigla_curso,
	}
}

func AddNewCurso(u models.Curso) (output Response.AddNewCursoResponse) {
	return Response.AddNewCursoResponse{
		Id: u.Id,
	}
}

func PutOneCurso(u models.Curso) (output Response.PutOneCursoResponse) {
	return Response.PutOneCursoResponse{
		Id: u.Id,
	}
}

func DeleteCurso(u models.Curso) (output Response.DeleteCursoResponse) {
	return Response.DeleteCursoResponse{
		Id: u.Id,
	}
}
